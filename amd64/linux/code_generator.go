package amd64_linux

import (
  "fmt"
  xtc_asm "bitbucket.org/yyuu/xtc/asm"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_ir "bitbucket.org/yyuu/xtc/ir"
  "bitbucket.org/yyuu/xtc/amd64"
)

const (
  LABEL_SYMBOL_BASE = ".L"
  CONST_SYMBOL_BASE = ".LC"

// Flags
  SectionFlag_allocatable = "a"
  SectionFlag_writable = "w"
  SectionFlag_executable = "x"
  SectionFlag_sectiongroup = "G"
  SectionFlag_strings = "S"
  SectionFlag_threadlocalstorage = "T"

// argument of "G" flag
  Linkage_linkonce = "comdat"

// Types
  SectionType_bits = "@progbits"
  SectionType_nobits = "@nobits"
  SectionType_note = "@note"

  SymbolType_function = "@function"

  PICThunkSectionFlags = SectionFlag_allocatable + SectionFlag_executable + SectionFlag_sectiongroup

  STACK_WORD_SIZE = int64(4)
)

type CodeGenerator struct {
  errorHandler *xtc_core.ErrorHandler
  options *xtc_core.Options
  naturalType int
}

func NewCodeGenerator(errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options) *CodeGenerator {
  return &CodeGenerator { errorHandler, options, xtc_asm.TYPE_INT32 }
}

func (self *CodeGenerator) Generate(ir *xtc_ir.IR) (*AssemblyCode, error) {
  self.locateSymbols(ir)
  asm := self.generateAssemblyCode(ir)
  if self.errorHandler.ErrorOccured() {
    return nil, fmt.Errorf("found %d error(s).", self.errorHandler.GetErrors())
  }
  return asm, nil
}

func (self *CodeGenerator) locateSymbols(ir *xtc_ir.IR) {
  constSymbols := xtc_asm.NewSymbolTable(CONST_SYMBOL_BASE)
  es := ir.GetConstantTable().GetEntries()
  for i := range es {
    self.locateStringLiteral(es[i], constSymbols)
  }
  vs := ir.AllGlobalVariables()
  for i := range vs {
    self.locateGlobalVariable(vs[i])
  }
  fs := ir.AllFunctions()
  for i := range fs {
    self.locateFunction(fs[i])
  }
}

func (self *CodeGenerator) locateStringLiteral(ent *xtc_entity.ConstantEntry, syms *xtc_asm.SymbolTable) {
  ent.SetSymbol(syms.NewSymbol())
  if self.options.IsPositionIndependent() {
    offset := self.localGOTSymbol(ent.GetSymbol())
    ent.SetMemref(self.mem4(offset, self.gotBaseReg()))
  } else {
    ent.SetMemref(self.mem1(ent.GetSymbol()))
    ent.SetAddress(self.imm2(ent.GetSymbol()))
  }
}

func (self *CodeGenerator) locateGlobalVariable(ent xtc_core.IEntity) {
  sym := self.symbol(ent.SymbolString(), ent.IsPrivate())
  if self.options.IsPositionIndependent() {
    if ent.IsPrivate() || self.optimizeGvarAccess(ent) {
      ent.SetMemref(self.mem4(self.globalGOTSymbol(sym), self.gotBaseReg()))
    } else {
      ent.SetAddress(self.mem4(self.globalGOTSymbol(sym), self.gotBaseReg()))
    }
  } else {
    ent.SetMemref(self.mem1(sym))
    ent.SetAddress(self.imm2(sym))
  }
}

func (self *CodeGenerator) locateFunction(fun xtc_core.IFunction) {
  fun.SetCallingSymbol(self.callingSymbol(fun))
  self.locateGlobalVariable(fun)
}

func (self *CodeGenerator) symbol(sym string, isPrivate bool) xtc_core.ISymbol {
  if isPrivate {
    return self.privateSymbol(sym)
  } else {
    return self.globalSymbol(sym)
  }
}

func (self *CodeGenerator) globalSymbol(sym string) xtc_core.ISymbol {
  return xtc_asm.NewNamedSymbol(sym)
}

func (self *CodeGenerator) privateSymbol(sym string) xtc_core.ISymbol {
  return xtc_asm.NewNamedSymbol(sym)
}

func (self *CodeGenerator) callingSymbol(fun xtc_core.IFunction) xtc_core.ISymbol {
  if fun.IsPrivate() {
    return self.privateSymbol(fun.SymbolString())
  } else {
    sym := self.globalSymbol(fun.SymbolString())
    if self.shouldUsePLT(fun) {
      return self.pltSymbol(sym)
    } else {
      return sym
    }
  }
}

func (self *CodeGenerator) shouldUsePLT(ent xtc_core.IEntity) bool {
  return self.options.IsPositionIndependent() && !self.optimizeGvarAccess(ent)
}

func (self *CodeGenerator) optimizeGvarAccess(ent xtc_core.IEntity) bool {
  return self.options.IsPIERequired() && ent.IsDefined()
}

func (self *CodeGenerator) generateAssemblyCode(ir *xtc_ir.IR) *AssemblyCode {
  file := self.newAssemblyCode()
  file._file(ir.GetFileName())
  if ir.IsGlobalVariableDefined() {
    self.generateDataSection(file, ir.GetDefinedGlobalVariables())
  }
  if ir.IsStringLiteralDefined() {
    self.generateReadOnlyDataSection(file, ir.GetConstantTable())
  }
  if ir.IsFunctionDefined() {
    self.generateTextSection(file, ir.GetDefinedFunctions())
  }
  if ir.IsCommonSymbolDefined() {
    self.generateCommonSymbols(file, ir.GetDefinedCommonSymbols())
  }
  if self.options.IsPositionIndependent() {
    self.picThunk(file, self.gotBaseReg())
  }
  return file
}

func (self *CodeGenerator) newAssemblyCode() *AssemblyCode {
  return NewAssemblyCode(self.naturalType, xtc_asm.NewSymbolTable(LABEL_SYMBOL_BASE))
}

func (self *CodeGenerator) generateDataSection(file *AssemblyCode, gvars []*xtc_entity.DefinedVariable) {
  file._data()
  for i := range gvars {
    gvar := gvars[i]
    sym := self.globalSymbol(gvar.SymbolString())
    if !gvar.IsPrivate() {
      file._globl(sym)
    }
    file._align(int64(gvar.GetType().Alignment()))
    file._type(sym, "@object")
    file._size(sym, fmt.Sprintf("%d", gvar.GetType().AllocSize()))
    file.label2(sym)
    self.generateImmediate(file, int64(gvar.GetType().AllocSize()), gvar.GetIR())
  }
}

func (self *CodeGenerator) generateImmediate(file *AssemblyCode, size int64, node xtc_core.IExpr) {
  switch expr := node.(type) {
    case *xtc_ir.Int: {
      switch size {
        case 1: file._byte(xtc_asm.NewIntegerLiteral(expr.GetValue()))
        case 2: file._value(xtc_asm.NewIntegerLiteral(expr.GetValue()))
        case 4: file._long(xtc_asm.NewIntegerLiteral(expr.GetValue()))
        case 8: file._quad(xtc_asm.NewIntegerLiteral(expr.GetValue()))
        default: {
          panic("entry size must be 1,2,4,8")
        }
      }
    }
    case *xtc_ir.Str: {
      switch size {
        case 4: file._long(expr.GetSymbol())
        case 8: file._quad(expr.GetSymbol())
        default: {
          panic("pointer size must be 4,8")
        }
      }
    }
    default: {
      panic(fmt.Errorf("unknown literal node type: %s", node))
    }
  }
}

func (self *CodeGenerator) generateReadOnlyDataSection(file *AssemblyCode, constants *xtc_entity.ConstantTable) {
  file._section(".rodata")
  entries := constants.GetEntries()
  for i := range entries {
    ent := entries[i]
    file.label2(ent.GetSymbol())
    file._string(ent.GetValue())
  }
}

func (self *CodeGenerator) generateTextSection(file *AssemblyCode, functions []*xtc_entity.DefinedFunction) {
  file._text()
  for i := range functions {
    fun := functions[i]
    sym := self.globalSymbol(fun.GetName())
    if ! fun.IsPrivate() {
      file._globl(sym)
    }
    file._type(sym, "@function")
    file.label2(sym)
    self.compileFunctionBody(file, fun)
    file._size(sym, fmt.Sprintf(".-%s", sym))
  }
}

func (self *CodeGenerator) generateCommonSymbols(file *AssemblyCode, variables []*xtc_entity.DefinedVariable) {
  for i := range variables {
    v := variables[i]
    sym := self.globalSymbol(v.SymbolString())
    if v.IsPrivate() {
      file._local(sym)
    }
    t := v.GetType()
    file._comm(sym, int64(t.AllocSize()), int64(t.Alignment()))
  }
}

//
// PIC/PIE related constants and codes
//
var got = xtc_asm.NewNamedSymbol("_GLOBAL_OFFSET_TABLE_")
func (self *CodeGenerator) loadGOTBaseAddress(file *AssemblyCode, reg *amd64.Register) {
  file.call(self.picThunkSymbol(reg))
  file.add(self.imm2(got), reg)
}

func (self *CodeGenerator) gotBaseReg() *amd64.Register {
  return self.bx()
}

func (self *CodeGenerator) globalGOTSymbol(base xtc_core.ISymbol) xtc_core.ISymbol {
  return xtc_asm.NewSuffixedSymbol(base, "@GOT")
}

func (self *CodeGenerator) localGOTSymbol(base xtc_core.ISymbol) xtc_core.ISymbol {
  return xtc_asm.NewSuffixedSymbol(base, "@GOTOFF")
}

func (self *CodeGenerator) pltSymbol(base xtc_core.ISymbol) xtc_core.ISymbol {
  return xtc_asm.NewSuffixedSymbol(base, "@PLT")
}

func (self *CodeGenerator) picThunkSymbol(reg *amd64.Register) xtc_core.ISymbol {
  return xtc_asm.NewNamedSymbol("__i686.get_pc_thunk." + reg.GetBaseName())
}

/**
 * Output PIC thunk.
 * ELF section declaration format is:
 *
 *     .section NAME, FLAGS, TYPE, flag_arguments
 *
 * FLAGS, TYPE, flag_arguments and optional.
 * For "M" flag (a member of a section group),
 * following format is used:
 *
 *     .section NAME, "...M", TYPE, section_group_name, linkage
 */
func (self *CodeGenerator) picThunk(file *AssemblyCode, reg *amd64.Register) {
  sym := self.picThunkSymbol(reg)
  file._section2(fmt.Sprintf(".text.%s", sym),
                 fmt.Sprintf("\"%s\"", PICThunkSectionFlags),
                 SectionType_bits, fmt.Sprint(sym), Linkage_linkonce)
  file._globl(sym)
  file._hidden(sym)
  file._type(sym, "@function")
  file.label2(sym)
  file.mov2(self.mem2(self.sp()), reg)
  file.ret()
}

//
// Compile Function
//

/* Standard IA-32 stack frame layout
 *
 * ======================= esp #3 (stack top just before function call)
 * next arg 1
 * ---------------------
 * next arg 2
 * ---------------------
 * next arg 3
 * ---------------------   esp #2 (stack top after alloca call)
 * alloca area
 * ---------------------   esp #1 (stack top just after prelude)
 * temporary
 * variables...
 * ---------------------   -16(%ebp)
 * lvar 3
 * ---------------------   -12(%ebp)
 * lvar 2
 * ---------------------   -8(%ebp)
 * lvar 1
 * ---------------------   -4(%ebp)
 * callee-saved register
 * ======================= 0(%ebp)
 * saved ebp
 * ---------------------   4(%ebp)
 * return address
 * ---------------------   8(%ebp)
 * arg 1
 * ---------------------   12(%ebp)
 * arg 2
 * ---------------------   16(%ebp)
 * arg 3
 * ...
 * ...
 * ======================= stack bottom
 */

func (self *CodeGenerator) alignStack(size int64) int64 {
  return (size + STACK_WORD_SIZE - 1) / STACK_WORD_SIZE * STACK_WORD_SIZE
}

func (self *CodeGenerator) stackSizeFromWordNum(numWords int64) int64 {
  return numWords * STACK_WORD_SIZE
}

type stackFrameInfo struct {
  saveRegs []*amd64.Register
  lvarSize int64
  tempSize int64
}

func (self *stackFrameInfo) saveRegsSize() int64 {
  return int64(len(self.saveRegs)) * STACK_WORD_SIZE
}

func (self *stackFrameInfo) lvarOffset() int64 {
  return self.saveRegsSize()
}

func (self *stackFrameInfo) tempOffset() int64 {
  return self.saveRegsSize() + self.lvarSize
}

func (self *stackFrameInfo) frameSize() int64 {
  return self.saveRegsSize() + self.lvarSize + self.tempSize
}

func (self *CodeGenerator) compileFunctionBody(file *AssemblyCode, fun *xtc_entity.DefinedFunction) {
  frame := &stackFrameInfo { []*amd64.Register { }, int64(0), int64(0) }
  self.locateParameters(fun.GetParameters())
  frame.lvarSize = self.locateLocalVariables(fun.LocalVariableScope(), int64(0))

  body := self.optimize(self.compileStmts(fun))
  frame.saveRegs = self.usedCalleeSaveRegisters(body)
  frame.tempSize = body.virtualStack.MaxSize()

  self.fixLocalVariableOffsets(fun.LocalVariableScope(), frame.lvarOffset())
  self.fixTempVariableOffsets(body, frame.tempOffset())

  if self.options.IsVerboseAsm() {
    self.printStackFrameLayout(file, frame, fun.GetLocalVariables())
  }
  self.generateFunctionBody(file, body, frame)
}

func (self *CodeGenerator) optimize(body *AssemblyCode) *AssemblyCode {
  self.errorHandler.Warn("FIXME: CodeGenerator#optimize: not implemented")
  return body
}

func (self *CodeGenerator) printStackFrameLayout(file *AssemblyCode, frame *stackFrameInfo, lvars []*xtc_entity.DefinedVariable) {
  vars := []*memInfo { }
  for i := range lvars {
    vars = append(vars, &memInfo { lvars[i].GetMemref(), lvars[i].GetName() })
  }
  vars = append(vars, &memInfo { self.mem3(int64(0), self.bp()), "return address" })
  vars = append(vars, &memInfo { self.mem3(int64(4), self.bp()), "saved %ebp" })
  if 0 < frame.saveRegsSize() {
    vars = append(vars, &memInfo {
      self.mem3(-frame.saveRegsSize(), self.bp()),
      fmt.Sprintf("saved callee-saved registers (%d bytes)", frame.saveRegsSize()),
    })
  }
  if 0 < frame.tempSize {
    vars = append(vars, &memInfo {
      self.mem3(-frame.frameSize(), self.bp()),
      fmt.Sprintf("tmp variables (%d bytes)", frame.tempSize),
    })
  }
  // TODO: sort vars
  file.comment("---- Stack Frame Layout -----------")
  for i := range vars {
    file.comment(fmt.Sprintf("%s: %s", vars[i].mem, vars[i].name))
  }
  file.comment("-----------------------------------")
}

type memInfo struct {
  mem xtc_core.IMemoryReference
  name string
}

var as *AssemblyCode
var epilogue *xtc_asm.Label

func (self *CodeGenerator) compileStmts(fun *xtc_entity.DefinedFunction) *AssemblyCode {
  as = self.newAssemblyCode()
  epilogue = xtc_asm.NewUnnamedLabel()
  stmts := fun.GetIR()
  for i := range stmts {
    self.compileStmt(stmts[i])
  }
  as.label1(epilogue)
  return as
}

func (self *CodeGenerator) usedCalleeSaveRegisters(body *AssemblyCode) []*amd64.Register {
  result := []*amd64.Register { }
  regs := self.calleeSaveRegisters()
  for i := range regs {
    reg := regs[i]
    if body.doesUses(reg) {
      if reg != self.bp() {
        result = append(result, reg)
      }
    }
  }
  return result
}

var CALLEE_SAVE_REGISTERS = []int { amd64.BX, amd64.BP, amd64.SI, amd64.DI }

func (self *CodeGenerator) calleeSaveRegisters() []*amd64.Register {
  regs := make([]*amd64.Register, len(CALLEE_SAVE_REGISTERS))
  for i := range CALLEE_SAVE_REGISTERS {
    regs[i] = amd64.NewRegister(CALLEE_SAVE_REGISTERS[i], self.naturalType)
  }
  return regs
}

func (self *CodeGenerator) generateFunctionBody(file *AssemblyCode, body *AssemblyCode, frame *stackFrameInfo) {
  file.virtualStack.Reset()
  self.prologue(file, frame.saveRegs, frame.frameSize())
  if self.options.IsPositionIndependent() && body.doesUses(self.gotBaseReg()) {
    self.loadGOTBaseAddress(file, self.gotBaseReg())
  }
  file.addAll(body.GetAssemblies())
  self.epilogue(file, frame.saveRegs)
  file.virtualStack.FixOffset(0)
}

func (self *CodeGenerator) prologue(file *AssemblyCode, saveRegs []*amd64.Register, frameSize int64) {
  file.push(self.bp())
  file.mov1(self.sp(), self.bp())
  for i := range saveRegs {
    reg := saveRegs[i]
    file.virtualPush(reg)
  }
  self.extendStack(file, frameSize)
}

func (self *CodeGenerator) epilogue(file *AssemblyCode, savedRegs []*amd64.Register) {
  for i := range savedRegs {
    reg := savedRegs[len(savedRegs)-1-i]
    file.virtualPop(reg)
  }
  file.mov1(self.bp(), self.sp())
  file.pop(self.bp())
  file.ret()
}

const PARAM_START_WORD = int64(2) // return addr and saved up

func (self *CodeGenerator) locateParameters(params []*xtc_entity.Parameter) {
  numWords := PARAM_START_WORD
  for i := range params {
    params[i].SetMemref(self.mem3(self.stackSizeFromWordNum(numWords), self.bp()))
    numWords++
  }
}

/**
 * Allocate addresses of local variables, but offset is still
 * not determined, assign unfixed IndirectMemoryReference.
 */
func (self *CodeGenerator) locateLocalVariables(scope *xtc_entity.LocalScope, parentStackLen int64) int64 {
  n := parentStackLen
  vars := scope.GetLocalVariables()
  for i := range vars {
    n = self.alignStack(n + int64(vars[i].GetType().AllocSize()))
    vars[i].SetMemref(self.relocatableMem(-n, self.bp()))
  }
  maxLen := n
  scopes := scope.GetChildren()
  for i := range scopes {
    children := self.locateLocalVariables(scopes[i], n)
    if maxLen < children {
      maxLen = children
    }
  }
  return maxLen
}

func (self *CodeGenerator) relocatableMem(offset int64, base *amd64.Register) *xtc_asm.IndirectMemoryReference {
  return xtc_asm.NewIndirectMemoryReference(xtc_asm.NewIntegerLiteral(offset), base, false)
}

func (self *CodeGenerator) fixLocalVariableOffsets(scope *xtc_entity.LocalScope, n int64) {
  vs := scope.AllLocalVariables()
  for i := range vs {
    vs[i].GetMemref().FixOffset(-n)
  }
}

func (self *CodeGenerator) fixTempVariableOffsets(asm *AssemblyCode, n int64) {
  asm.virtualStack.FixOffset(-n)
}

func (self *CodeGenerator) extendStack(file *AssemblyCode, n int64) {
  if 0 < n {
    file.sub(self.imm1(n), self.sp())
  }
}

func (self *CodeGenerator) rewindStack(file *AssemblyCode, n int64) {
  if 0 < n {
    file.add(self.imm1(n), self.sp())
  }

}

//
// Statements
//

func (self *CodeGenerator) compileStmt(stmt xtc_core.IStmt) {
  if self.options.IsVerboseAsm() {
    as.comment(fmt.Sprint(stmt.GetLocation()))
  }
  xtc_ir.VisitStmt(self, stmt)
}

//
// Expressions
//

func (self *CodeGenerator) compile(n xtc_core.IExpr) {
  if self.options.IsVerboseAsm() {
    as.comment(fmt.Sprintf("%s {", n))
    as.indentComment()
  }
  xtc_ir.VisitExpr(self, n)
  if self.options.IsVerboseAsm() {
    as.unindentComment()
    as.comment("}")
  }
}

func (self *CodeGenerator) doesRequireRegisterOperand(op int) bool {
  switch op {
    case xtc_ir.OP_S_DIV:      fallthrough
    case xtc_ir.OP_U_DIV:      fallthrough
    case xtc_ir.OP_S_MOD:      fallthrough
    case xtc_ir.OP_U_MOD:      fallthrough
    case xtc_ir.OP_BIT_LSHIFT: fallthrough
    case xtc_ir.OP_BIT_RSHIFT: fallthrough
    case xtc_ir.OP_ARITH_RSHIFT: {
      return true
    }
    default: {
      return false
    }
  }
}

func (self *CodeGenerator) compileBinaryOp(op int, left *amd64.Register, right xtc_core.IOperand) {
  switch op {
    case xtc_ir.OP_ADD: as.add(right, left)
    case xtc_ir.OP_SUB: as.sub(right, left)
    case xtc_ir.OP_MUL: as.imul(right, left)
    case xtc_ir.OP_S_DIV: fallthrough
    case xtc_ir.OP_S_MOD: {
      as.cltd()
      as.idiv(self.cxT(left.GetTypeId()))
      if op == xtc_ir.OP_S_MOD {
        as.mov1(self.dx(), left)
      }
    }
    case xtc_ir.OP_U_DIV: fallthrough
    case xtc_ir.OP_U_MOD: {
      as.mov2(self.imm1(int64(0)), self.dx())
      as.div(self.cxT(left.GetTypeId()))
      if op == xtc_ir.OP_U_MOD {
        as.mov3(self.dx(), left)
      }
    }
    case xtc_ir.OP_BIT_AND: as.and(right, left)
    case xtc_ir.OP_BIT_OR:  as.or(right, left)
    case xtc_ir.OP_BIT_XOR: as.xor(right, left)
    case xtc_ir.OP_BIT_LSHIFT:   as.sal(self.cl(), left)
    case xtc_ir.OP_BIT_RSHIFT:   as.shr(self.cl(), left)
    case xtc_ir.OP_ARITH_RSHIFT: as.sar(self.cl(), left)
    default: {
      as.cmp(right, self.axT(left.GetTypeId()))
      switch op {
        case xtc_ir.OP_EQ:     as.sete(self.al())
        case xtc_ir.OP_NEQ:    as.setne(self.al())
        case xtc_ir.OP_S_GT:   as.setg(self.al())
        case xtc_ir.OP_S_GTEQ: as.setge(self.al())
        case xtc_ir.OP_S_LT:   as.setl(self.al())
        case xtc_ir.OP_S_LTEQ: as.setle(self.al())
        case xtc_ir.OP_U_GT:   as.seta(self.al())
        case xtc_ir.OP_U_GTEQ: as.setae(self.al())
        case xtc_ir.OP_U_LT:   as.setb(self.al())
        case xtc_ir.OP_U_LTEQ: as.setbe(self.al())
        default: {
          panic(fmt.Errorf("unknown binary operator: %d", op))
        }
      }
      as.movzx(self.al(), left)
    }
  }
}

//
// Utilities
//

/**
 * Load constant value.  You must check node by #isConstant
 * before calling this method.
 */
func (self *CodeGenerator) loadConstant(node xtc_core.IExpr, reg *amd64.Register) {
  if node.GetAsmValue() != nil {
    as.mov2(node.GetAsmValue(), reg)
  } else {
    if node.GetMemref() != nil {
      as.lea(node.GetMemref(), reg)
    } else {
      panic("must not happen: constant has no asm value")
    }
  }
}

func (self *CodeGenerator) loadVariable(v *xtc_ir.Var, dest *amd64.Register) {
  if v.GetMemref() == nil {
    a := dest.ForType(self.naturalType)
    as.mov2(v.GetAddress(), a)
    self.load(self.mem2(a), dest.ForType(v.GetTypeId()))
  } else {
    self.load(v.GetMemref(), dest.ForType(v.GetTypeId()))
  }
}

func (self *CodeGenerator) loadAddress(v xtc_core.IEntity, dest *amd64.Register) {
  if v.GetAddress() != nil {
    as.mov2(v.GetAddress(), dest)
  } else {
    as.lea(v.GetMemref(), dest)
  }
}

func (self *CodeGenerator) ax() *amd64.Register {
  return amd64.NewRegister(amd64.AX, self.naturalType)
}

func (self *CodeGenerator) al() *amd64.Register {
  return self.axT(xtc_asm.TYPE_INT8)
}

func (self *CodeGenerator) bx() *amd64.Register {
  return amd64.NewRegister(amd64.BX, self.naturalType)
}

func (self *CodeGenerator) cx() *amd64.Register {
  return amd64.NewRegister(amd64.CX, self.naturalType)
}

func (self *CodeGenerator) cl() *amd64.Register {
  return self.cxT(xtc_asm.TYPE_INT8)
}

func (self *CodeGenerator) dx() *amd64.Register {
  return amd64.NewRegister(amd64.DX, self.naturalType)
}

func (self *CodeGenerator) axT(t int) *amd64.Register {
  return amd64.NewRegister(amd64.AX, t)
}

func (self *CodeGenerator) bxT(t int) *amd64.Register {
  return amd64.NewRegister(amd64.BX, t)
}

func (self *CodeGenerator) cxT(t int) *amd64.Register {
  return amd64.NewRegister(amd64.CX, t)
}

func (self *CodeGenerator) dxT(t int) *amd64.Register {
  return amd64.NewRegister(amd64.DX, t)
}

func (self *CodeGenerator) si() *amd64.Register {
  return amd64.NewRegister(amd64.SI, self.naturalType)
}

func (self *CodeGenerator) di() *amd64.Register {
  return amd64.NewRegister(amd64.DI, self.naturalType)
}

func (self *CodeGenerator) bp() *amd64.Register {
  return amd64.NewRegister(amd64.BP, self.naturalType)
}

func (self *CodeGenerator) sp() *amd64.Register {
  return amd64.NewRegister(amd64.SP, self.naturalType)
}

func (self *CodeGenerator) mem1(sym xtc_core.ISymbol) *xtc_asm.DirectMemoryReference {
  return xtc_asm.NewDirectMemoryReference(sym)
}

func (self *CodeGenerator) mem2(reg *amd64.Register) *xtc_asm.IndirectMemoryReference {
  return xtc_asm.NewIndirectMemoryReference(xtc_asm.NewIntegerLiteral(0), reg, true)
}

func (self *CodeGenerator) mem3(offset int64, reg *amd64.Register) *xtc_asm.IndirectMemoryReference {
  return xtc_asm.NewIndirectMemoryReference(xtc_asm.NewIntegerLiteral(offset), reg, true)
}

func (self *CodeGenerator) mem4(offset xtc_core.ISymbol, reg *amd64.Register) *xtc_asm.IndirectMemoryReference {
  return xtc_asm.NewIndirectMemoryReference(offset, reg, true)
}

func (self *CodeGenerator) imm1(n int64) *xtc_asm.ImmediateValue {
  return xtc_asm.NewImmediateValue(xtc_asm.NewIntegerLiteral(n))
}

func (self *CodeGenerator) imm2(lit xtc_core.ILiteral) *xtc_asm.ImmediateValue {
  return xtc_asm.NewImmediateValue(lit)
}

func (self *CodeGenerator) load(mem xtc_core.IMemoryReference, reg *amd64.Register) {
  as.mov2(mem, reg)
}

func (self *CodeGenerator) store(reg *amd64.Register, mem xtc_core.IMemoryReference) {
  as.mov3(reg, mem)
}

func (self *CodeGenerator) VisitStmt(unknown xtc_core.IStmt) interface{} {
  switch node := unknown.(type) {
    case *xtc_ir.Assign: {
      switch {
        case node.GetLHS().IsAddr() && node.GetLHS().GetMemref() != nil: {
          self.compile(node.GetRHS())
          self.store(self.axT(node.GetLHS().GetTypeId()), node.GetLHS().GetMemref())
        }
        case node.GetRHS().IsConstant(): {
          self.compile(node.GetLHS())
          as.mov1(self.ax(), self.cx())
          self.loadConstant(node.GetRHS(), self.ax())
          self.store(self.axT(node.GetLHS().GetTypeId()), self.mem2(self.cx()))
        }
        default: {
          self.compile(node.GetRHS())
          as.virtualPush(self.ax())
          self.compile(node.GetLHS())
          as.mov1(self.ax(), self.cx())
          as.virtualPop(self.ax())
          self.store(self.axT(node.GetLHS().GetTypeId()), self.mem2(self.cx()))
        }
      }
    }
    case *xtc_ir.CJump: {
      self.compile(node.GetCond())
      t := node.GetCond().GetTypeId()
      as.test(self.axT(t), self.axT(t))
      as.jnz(node.GetThenLabel())
      as.jmp(node.GetElseLabel())
    }
    case *xtc_ir.ExprStmt: {
      self.compile(node.GetExpr())
    }
    case *xtc_ir.Jump: {
      as.jmp(node.GetLabel())
    }
    case *xtc_ir.LabelStmt: {
      as.label1(node.GetLabel())
    }
    case *xtc_ir.Return: {
      if node.HasExpr() {
        self.compile(node.GetExpr())
      }
      as.jmp(epilogue)
    }
    case *xtc_ir.Switch: {
      self.compile(node.GetCond())
      t := node.GetCond().GetTypeId()
      cases := node.GetCases()
      for i := range cases {
        c := cases[i]
        as.mov2(self.imm1(c.GetValue()), self.cx())
        as.cmp(self.cxT(t), self.axT(t))
        as.je(c.GetLabel())
      }
      as.jmp(node.GetDefaultLabel())
    }
    default: {
      panic(fmt.Errorf("must not happen: unknown IR stmt: %s", unknown))
    }
  }
  return nil
}

func (self *CodeGenerator) VisitExpr(unknown xtc_core.IExpr) interface{} {
  switch node := unknown.(type) {
    case *xtc_ir.Addr: {
      self.loadAddress(node.GetEntity(), self.ax())
    }
    case *xtc_ir.Bin: {
      op := node.GetOp()
      t := node.GetTypeId()
      switch {
        case node.GetRight().IsConstant() && ! self.doesRequireRegisterOperand(op): {
          self.compile(node.GetLeft())
          self.compileBinaryOp(op, self.axT(t), node.GetRight().GetAsmValue())
        }
        case node.GetRight().IsConstant(): {
          self.compile(node.GetLeft())
          self.loadConstant(node.GetRight(), self.cx())
          self.compileBinaryOp(op, self.axT(t), self.cxT(t))
        }
        case node.GetRight().IsVar(): {
          self.compile(node.GetLeft())
          self.loadVariable(node.GetRight().(*xtc_ir.Var), self.cxT(t))
          self.compileBinaryOp(op, self.axT(t), self.cxT(t))
        }
        case node.GetRight().IsAddr(): {
          self.compile(node.GetLeft())
          self.loadAddress(node.GetRight().GetEntityForce(), self.cxT(t))
          self.compileBinaryOp(op, self.axT(t), self.cxT(t))
        }
        case node.GetRight().IsConstant() || node.GetLeft().IsVar() || node.GetLeft().IsAddr(): {
          self.compile(node.GetRight())
          as.mov1(self.ax(), self.cx())
          self.compile(node.GetLeft())
          self.compileBinaryOp(op, self.axT(t), self.cxT(t))
        }
        default: {
          self.compile(node.GetRight())
          as.virtualPush(self.ax())
          self.compile(node.GetLeft())
          as.virtualPop(self.cx())
          self.compileBinaryOp(op, self.axT(t), self.cxT(t))
        }
      }
    }
    case *xtc_ir.Call: {
      args := node.GetArgs()
      for i := range args {
        arg := args[len(args)-1-i]
        self.compile(arg)
        as.push(self.ax())
      }
      if node.IsStaticCall() {
        as.call(node.GetFunction().GetCallingSymbol())
      } else {
        self.compile(node.GetExpr())
        as.callAbsolute(self.ax())
      }
      self.rewindStack(as, self.stackSizeFromWordNum(int64(node.NumArgs())))
    }
    case *xtc_ir.Int: {
      as.mov2(self.imm1(node.GetValue()), self.ax())
    }
    case *xtc_ir.Mem: {
      self.compile(node.GetExpr())
      self.load(self.mem2(self.ax()), self.axT(node.GetTypeId()))
    }
    case *xtc_ir.Str: {
      self.loadConstant(node, self.ax())
    }
    case *xtc_ir.Uni: {
      src := node.GetExpr().GetTypeId()
      dest := node.GetTypeId()
      self.compile(node.GetExpr())
      switch node.GetOp() {
        case xtc_ir.OP_UMINUS: {
          as.neg(self.axT(src))
        }
        case xtc_ir.OP_BIT_NOT: {
          as.not(self.axT(src))
        }
        case xtc_ir.OP_NOT: {
          as.test(self.axT(src), self.axT(src))
          as.sete(self.al())
          as.movzx(self.al(), self.axT(dest))
        }
        case xtc_ir.OP_S_CAST: {
          as.movsx(self.axT(src), self.axT(dest))
        }
        case xtc_ir.OP_U_CAST: {
          as.movzx(self.axT(src), self.axT(dest))
        }
        default: {
          panic(fmt.Errorf("unknown unary operator: %d", node.GetOp()))
        }
      }
    }
    case *xtc_ir.Var: {
      self.loadVariable(node, self.ax())
    }
    default: {
      panic(fmt.Errorf("must not happen: unknown IR expr: %s", unknown))
    }
  }
  return nil
}
