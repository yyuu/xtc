package sysdep

import (
  "fmt"
  bs_asm "bitbucket.org/yyuu/bs/asm"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_ir "bitbucket.org/yyuu/bs/ir"
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

type LinuxX86CodeGenerator struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
  naturalType int
}

func NewLinuxX86CodeGenerator(errorHandler *bs_core.ErrorHandler, options *bs_core.Options) *LinuxX86CodeGenerator {
  return &LinuxX86CodeGenerator { errorHandler, options, bs_asm.TYPE_INT32 }
}

func (self *LinuxX86CodeGenerator) Generate(ir *bs_ir.IR) AssemblyCode {
  self.errorHandler.Debug("starting code generator.")
  self.locateSymbols(ir)
  asm := self.generateAssemblyCode(ir)
  self.errorHandler.Debug("finished code generator.")
  return asm
}

func (self *LinuxX86CodeGenerator) locateSymbols(ir *bs_ir.IR) {
  constSymbols := bs_asm.NewSymbolTable(CONST_SYMBOL_BASE)
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

func (self *LinuxX86CodeGenerator) locateStringLiteral(ent *bs_entity.ConstantEntry, syms *bs_asm.SymbolTable) {
  ent.SetSymbol(syms.NewSymbol())
  if self.options.IsPositionIndependent() {
    offset := self.localGOTSymbol(ent.GetSymbol())
    ent.SetMemref(self.mem4(offset, self.gotBaseReg()))
  } else {
    ent.SetMemref(self.mem1(ent.GetSymbol()))
    ent.SetAddress(self.imm2(ent.GetSymbol()))
  }
}

func (self *LinuxX86CodeGenerator) locateGlobalVariable(ent bs_core.IEntity) {
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

func (self *LinuxX86CodeGenerator) locateFunction(fun bs_core.IFunction) {
  fun.SetCallingSymbol(self.callingSymbol(fun))
  self.locateGlobalVariable(fun)
}

func (self *LinuxX86CodeGenerator) symbol(sym string, isPrivate bool) bs_core.ISymbol {
  if isPrivate {
    return self.privateSymbol(sym)
  } else {
    return self.globalSymbol(sym)
  }
}

func (self *LinuxX86CodeGenerator) globalSymbol(sym string) bs_core.ISymbol {
  return bs_asm.NewNamedSymbol(sym)
}

func (self *LinuxX86CodeGenerator) privateSymbol(sym string) bs_core.ISymbol {
  return bs_asm.NewNamedSymbol(sym)
}

func (self *LinuxX86CodeGenerator) callingSymbol(fun bs_core.IFunction) bs_core.ISymbol {
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

func (self *LinuxX86CodeGenerator) shouldUsePLT(ent bs_core.IEntity) bool {
  return self.options.IsPositionIndependent() && !self.optimizeGvarAccess(ent)
}

func (self *LinuxX86CodeGenerator) optimizeGvarAccess(ent bs_core.IEntity) bool {
  return self.options.IsPIERequired() && ent.IsDefined()
}

func (self *LinuxX86CodeGenerator) generateAssemblyCode(ir *bs_ir.IR) *LinuxX86AssemblyCode {
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

func (self *LinuxX86CodeGenerator) newAssemblyCode() *LinuxX86AssemblyCode {
  return NewLinuxX86AssemblyCode(self.naturalType, bs_asm.NewSymbolTable(LABEL_SYMBOL_BASE))
}

func (self *LinuxX86CodeGenerator) generateDataSection(file *LinuxX86AssemblyCode, gvars []*bs_entity.DefinedVariable) {
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

func (self *LinuxX86CodeGenerator) generateImmediate(file *LinuxX86AssemblyCode, size int64, node bs_core.IExpr) {
  switch expr := node.(type) {
    case *bs_ir.Int: {
      switch size {
        case 1: file._byte(bs_asm.NewIntegerLiteral(expr.GetValue()))
        case 2: file._value(bs_asm.NewIntegerLiteral(expr.GetValue()))
        case 4: file._long(bs_asm.NewIntegerLiteral(expr.GetValue()))
        case 8: file._quad(bs_asm.NewIntegerLiteral(expr.GetValue()))
        default: {
          panic("entry size must be 1,2,4,8")
        }
      }
    }
    case *bs_ir.Str: {
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

func (self *LinuxX86CodeGenerator) generateReadOnlyDataSection(file *LinuxX86AssemblyCode, constants *bs_entity.ConstantTable) {
  file._section(".rodata")
  entries := constants.GetEntries()
  for i := range entries {
    ent := entries[i]
    file.label2(ent.GetSymbol())
    file._string(ent.GetValue())
  }
}

func (self *LinuxX86CodeGenerator) generateTextSection(file *LinuxX86AssemblyCode, functions []*bs_entity.DefinedFunction) {
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

func (self *LinuxX86CodeGenerator) generateCommonSymbols(file *LinuxX86AssemblyCode, variables []*bs_entity.DefinedVariable) {
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
var got = bs_asm.NewNamedSymbol("_GLOBAL_OFFSET_TABLE_")
func (self *LinuxX86CodeGenerator) loadGOTBaseAddress(file *LinuxX86AssemblyCode, reg *x86Register) {
  file.call(self.picThunkSymbol(reg))
  file.add(self.imm2(got), reg)
}

func (self *LinuxX86CodeGenerator) gotBaseReg() *x86Register {
  return self.bx()
}

func (self *LinuxX86CodeGenerator) globalGOTSymbol(base bs_core.ISymbol) bs_core.ISymbol {
  return bs_asm.NewSuffixedSymbol(base, "@GOT")
}

func (self *LinuxX86CodeGenerator) localGOTSymbol(base bs_core.ISymbol) bs_core.ISymbol {
  return bs_asm.NewSuffixedSymbol(base, "@GOTOFF")
}

func (self *LinuxX86CodeGenerator) pltSymbol(base bs_core.ISymbol) bs_core.ISymbol {
  return bs_asm.NewSuffixedSymbol(base, "@PLT")
}

func (self *LinuxX86CodeGenerator) picThunkSymbol(reg *x86Register) bs_core.ISymbol {
  return bs_asm.NewNamedSymbol("__i686.get_pc_thunk." + reg.GetBaseName())
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
func (self *LinuxX86CodeGenerator) picThunk(file *LinuxX86AssemblyCode, reg *x86Register) {
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

func (self *LinuxX86CodeGenerator) alignStack(size int64) int64 {
  return (size + STACK_WORD_SIZE - 1) / STACK_WORD_SIZE * STACK_WORD_SIZE
}

func (self *LinuxX86CodeGenerator) stackSizeFromWordNum(numWords int64) int64 {
  return numWords * STACK_WORD_SIZE
}

type stackFrameInfo struct {
  saveRegs []*x86Register
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

func (self *LinuxX86CodeGenerator) compileFunctionBody(file *LinuxX86AssemblyCode, fun *bs_entity.DefinedFunction) {
  frame := &stackFrameInfo { []*x86Register { }, int64(0), int64(0) }
  self.locateParameters(fun.GetParameters())
  frame.lvarSize = self.locateLocalVariables(fun.LocalVariableScope(), int64(0))

  body := self.optimize(self.compileStmts(fun))
  frame.saveRegs = self.usedCalleeSaveRegisters(body)
  frame.tempSize = body.virtualStack.maxSize()

  self.fixLocalVariableOffsets(fun.LocalVariableScope(), frame.lvarOffset())
  self.fixTempVariableOffsets(body, frame.tempOffset())

  if self.options.IsVerboseAsm() {
    self.printStackFrameLayout(file, frame, fun.GetLocalVariables())
  }
  self.generateFunctionBody(file, body, frame)
}

func (self *LinuxX86CodeGenerator) optimize(body *LinuxX86AssemblyCode) *LinuxX86AssemblyCode {
  self.errorHandler.Warn("FIXME: CodeGenerator#optimize: not implemented")
  return body
}

func (self *LinuxX86CodeGenerator) printStackFrameLayout(file *LinuxX86AssemblyCode, frame *stackFrameInfo, lvars []*bs_entity.DefinedVariable) {
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
  mem bs_core.IMemoryReference
  name string
}

var as *LinuxX86AssemblyCode
var epilogue *bs_asm.Label

func (self *LinuxX86CodeGenerator) compileStmts(fun *bs_entity.DefinedFunction) *LinuxX86AssemblyCode {
  as = self.newAssemblyCode()
  epilogue = bs_asm.NewUnnamedLabel()
  stmts := fun.GetIR()
  for i := range stmts {
    self.compileStmt(stmts[i])
  }
  as.label1(epilogue)
  return as
}

func (self *LinuxX86CodeGenerator) usedCalleeSaveRegisters(body *LinuxX86AssemblyCode) []*x86Register {
  result := []*x86Register { }
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

var CALLEE_SAVE_REGISTERS = []int { x86_bx, x86_bp, x86_si, x86_di }

func (self *LinuxX86CodeGenerator) calleeSaveRegisters() []*x86Register {
  regs := make([]*x86Register, len(CALLEE_SAVE_REGISTERS))
  for i := range CALLEE_SAVE_REGISTERS {
    regs[i] = newX86Register(CALLEE_SAVE_REGISTERS[i], self.naturalType)
  }
  return regs
}

func (self *LinuxX86CodeGenerator) generateFunctionBody(file *LinuxX86AssemblyCode, body *LinuxX86AssemblyCode, frame *stackFrameInfo) {
  file.virtualStack.reset()
  self.prologue(file, frame.saveRegs, frame.frameSize())
  if self.options.IsPositionIndependent() && body.doesUses(self.gotBaseReg()) {
    self.loadGOTBaseAddress(file, self.gotBaseReg())
  }
  file.addAll(body.GetAssemblies())
  self.epilogue(file, frame.saveRegs)
  file.virtualStack.fixOffset(0)
}

func (self *LinuxX86CodeGenerator) prologue(file *LinuxX86AssemblyCode, saveRegs []*x86Register, frameSize int64) {
  file.push(self.bp())
  file.mov1(self.sp(), self.bp())
  for i := range saveRegs {
    reg := saveRegs[i]
    file.virtualPush(reg)
  }
  self.extendStack(file, frameSize)
}

func (self *LinuxX86CodeGenerator) epilogue(file *LinuxX86AssemblyCode, savedRegs []*x86Register) {
  for i := range savedRegs {
    reg := savedRegs[len(savedRegs)-1-i]
    file.virtualPop(reg)
  }
  file.mov1(self.bp(), self.sp())
  file.pop(self.bp())
  file.ret()
}

const PARAM_START_WORD = int64(2) // return addr and saved up

func (self *LinuxX86CodeGenerator) locateParameters(params []*bs_entity.Parameter) {
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
func (self *LinuxX86CodeGenerator) locateLocalVariables(scope *bs_entity.LocalScope, parentStackLen int64) int64 {
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

func (self *LinuxX86CodeGenerator) relocatableMem(offset int64, base *x86Register) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(bs_asm.NewIntegerLiteral(offset), base, false)
}

func (self *LinuxX86CodeGenerator) fixLocalVariableOffsets(scope *bs_entity.LocalScope, n int64) {
  vs := scope.AllLocalVariables()
  for i := range vs {
    vs[i].GetMemref().FixOffset(-n)
  }
}

func (self *LinuxX86CodeGenerator) fixTempVariableOffsets(asm *LinuxX86AssemblyCode, n int64) {
  asm.virtualStack.fixOffset(-n)
}

func (self *LinuxX86CodeGenerator) extendStack(file *LinuxX86AssemblyCode, n int64) {
  if 0 < n {
    file.sub(self.imm1(n), self.sp())
  }
}

func (self *LinuxX86CodeGenerator) rewindStack(file *LinuxX86AssemblyCode, n int64) {
  if 0 < n {
    file.add(self.imm1(n), self.sp())
  }

}

//
// Statements
//

func (self *LinuxX86CodeGenerator) compileStmt(stmt bs_core.IStmt) {
  if self.options.IsVerboseAsm() {
    as.comment(fmt.Sprint(stmt.GetLocation()))
  }
//stmt.accept(self) // FIXME:
}

//
// Expressions
//

func (self *LinuxX86CodeGenerator) compile(n bs_core.IExpr) {
  if self.options.IsVerboseAsm() {
    as.comment(fmt.Sprintf("%s {", n))
    as.indentComment()
  }
//n.accept(self) // FIXME:
  if self.options.IsVerboseAsm() {
    as.unindentComment()
    as.comment("}")
  }
}

func (self *LinuxX86CodeGenerator) doesRequireRegisterOperand(op int) bool {
  switch op {
    case bs_ir.OP_S_DIV:      fallthrough
    case bs_ir.OP_U_DIV:      fallthrough
    case bs_ir.OP_S_MOD:      fallthrough
    case bs_ir.OP_U_MOD:      fallthrough
    case bs_ir.OP_BIT_LSHIFT: fallthrough
    case bs_ir.OP_BIT_RSHIFT: fallthrough
    case bs_ir.OP_ARITH_RSHIFT: {
      return true
    }
    default: {
      return false
    }
  }
}

func (self *LinuxX86CodeGenerator) compileBinaryOp(op int, left *x86Register, right bs_core.IOperand) {
  switch op {
    case bs_ir.OP_ADD: as.add(right, left)
    case bs_ir.OP_SUB: as.sub(right, left)
    case bs_ir.OP_MUL: as.imul(right, left)
    case bs_ir.OP_S_DIV: fallthrough
    case bs_ir.OP_S_MOD: {
      as.cltd()
      as.idiv(self.cxT(left.GetTypeId()))
      if op == bs_ir.OP_S_MOD {
        as.mov1(self.dx(), left)
      }
    }
    case bs_ir.OP_U_DIV: fallthrough
    case bs_ir.OP_U_MOD: {
      as.mov2(self.imm1(int64(0)), self.dx())
      as.div(self.cxT(left.GetTypeId()))
      if op == bs_ir.OP_U_MOD {
        as.mov3(self.dx(), left)
      }
    }
    case bs_ir.OP_BIT_AND: as.and(right, left)
    case bs_ir.OP_BIT_OR:  as.or(right, left)
    case bs_ir.OP_BIT_XOR: as.xor(right, left)
    case bs_ir.OP_BIT_LSHIFT:   as.sal(self.cl(), left)
    case bs_ir.OP_BIT_RSHIFT:   as.shr(self.cl(), left)
    case bs_ir.OP_ARITH_RSHIFT: as.sar(self.cl(), left)
    default: {
      as.cmp(right, self.axT(left.GetTypeId()))
      switch op {
        case bs_ir.OP_EQ:     as.sete(self.al())
        case bs_ir.OP_NEQ:    as.setne(self.al())
        case bs_ir.OP_S_GT:   as.setg(self.al())
        case bs_ir.OP_S_GTEQ: as.setge(self.al())
        case bs_ir.OP_S_LT:   as.setl(self.al())
        case bs_ir.OP_S_LTEQ: as.setle(self.al())
        case bs_ir.OP_U_GT:   as.seta(self.al())
        case bs_ir.OP_U_GTEQ: as.setae(self.al())
        case bs_ir.OP_U_LT:   as.setb(self.al())
        case bs_ir.OP_U_LTEQ: as.setbe(self.al())
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
func (self *LinuxX86CodeGenerator) loadConstant(node bs_core.IExpr, reg *x86Register) {
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

func (self *LinuxX86CodeGenerator) loadVariable(v *bs_ir.Var, dest *x86Register) {
  if v.GetMemref() == nil {
    a := dest.ForType(self.naturalType)
    as.mov2(v.GetAddress(), a)
    self.load(self.mem2(a), dest.ForType(v.GetTypeId()))
  } else {
    self.load(v.GetMemref(), dest.ForType(v.GetTypeId()))
  }
}

func (self *LinuxX86CodeGenerator) loadAddress(v bs_core.IEntity, dest *x86Register) {
  if v.GetAddress() != nil {
    as.mov2(v.GetAddress(), dest)
  } else {
    as.lea(v.GetMemref(), dest)
  }
}

func (self *LinuxX86CodeGenerator) ax() *x86Register {
  return newX86Register(x86_ax, self.naturalType)
}

func (self *LinuxX86CodeGenerator) al() *x86Register {
  return self.axT(bs_asm.TYPE_INT8)
}

func (self *LinuxX86CodeGenerator) bx() *x86Register {
  return newX86Register(x86_bx, self.naturalType)
}

func (self *LinuxX86CodeGenerator) cx() *x86Register {
  return newX86Register(x86_cx, self.naturalType)
}

func (self *LinuxX86CodeGenerator) cl() *x86Register {
  return self.cxT(bs_asm.TYPE_INT8)
}

func (self *LinuxX86CodeGenerator) dx() *x86Register {
  return newX86Register(x86_dx, self.naturalType)
}

func (self *LinuxX86CodeGenerator) axT(t int) *x86Register {
  return newX86Register(x86_ax, t)
}

func (self *LinuxX86CodeGenerator) bxT(t int) *x86Register {
  return newX86Register(x86_bx, t)
}

func (self *LinuxX86CodeGenerator) cxT(t int) *x86Register {
  return newX86Register(x86_cx, t)
}

func (self *LinuxX86CodeGenerator) dxT(t int) *x86Register {
  return newX86Register(x86_dx, t)
}

func (self *LinuxX86CodeGenerator) si() *x86Register {
  return newX86Register(x86_si, self.naturalType)
}

func (self *LinuxX86CodeGenerator) di() *x86Register {
  return newX86Register(x86_di, self.naturalType)
}

func (self *LinuxX86CodeGenerator) bp() *x86Register {
  return newX86Register(x86_bp, self.naturalType)
}

func (self *LinuxX86CodeGenerator) sp() *x86Register {
  return newX86Register(x86_sp, self.naturalType)
}

func (self *LinuxX86CodeGenerator) mem1(sym bs_core.ISymbol) *bs_asm.DirectMemoryReference {
  return bs_asm.NewDirectMemoryReference(sym)
}

func (self *LinuxX86CodeGenerator) mem2(reg *x86Register) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(bs_asm.NewIntegerLiteral(0), reg, true)
}

func (self *LinuxX86CodeGenerator) mem3(offset int64, reg *x86Register) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(bs_asm.NewIntegerLiteral(offset), reg, true)
}

func (self *LinuxX86CodeGenerator) mem4(offset bs_core.ISymbol, reg *x86Register) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(offset, reg, true)
}

func (self *LinuxX86CodeGenerator) imm1(n int64) *bs_asm.ImmediateValue {
  return bs_asm.NewImmediateValue(bs_asm.NewIntegerLiteral(n))
}

func (self *LinuxX86CodeGenerator) imm2(lit bs_core.ILiteral) *bs_asm.ImmediateValue {
  return bs_asm.NewImmediateValue(lit)
}

func (self *LinuxX86CodeGenerator) load(mem bs_core.IMemoryReference, reg *x86Register) {
  as.mov2(mem, reg)
}

func (self *LinuxX86CodeGenerator) store(reg *x86Register, mem bs_core.IMemoryReference) {
  as.mov3(reg, mem)
}
