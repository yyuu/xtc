package sysdep

import (
  "fmt"
  bs_asm "bitbucket.org/yyuu/bs/asm"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_ir "bitbucket.org/yyuu/bs/ir"
)

const LABEL_SYMBOL_BASE = ".L"
const CONST_SYMBOL_BASE = ".LC"

type LinuxX86CodeGenerator struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
  naturalType int
  stackWordSize int64
}

func NewLinuxX86CodeGenerator(errorHandler *bs_core.ErrorHandler, options *bs_core.Options) *LinuxX86CodeGenerator {
  stackWordSize := int64(4)
  return &LinuxX86CodeGenerator { errorHandler, options, bs_asm.TYPE_INT32, stackWordSize }
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
  self.errorHandler.Warn("locateGlobalVariable not implemented")
}

func (self *LinuxX86CodeGenerator) locateFunction(fun bs_core.IFunction) {
  self.errorHandler.Warn("locateFunction not implemented")
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
  return NewLinuxX86AssemblyCode(self.naturalType, self.stackWordSize, bs_asm.NewSymbolTable(LABEL_SYMBOL_BASE))
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
    file.label(bs_asm.NewLabel(sym))
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
    file.label(bs_asm.NewLabel(ent.GetSymbol()))
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
    file.label(bs_asm.NewLabel(sym))
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
func (self *LinuxX86CodeGenerator) loadGOTBaseAddress(file *LinuxX86AssemblyCode, reg bs_core.IRegister) {
  // FIXME:
}

func (self *LinuxX86CodeGenerator) gotBaseReg() bs_core.IRegister {
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

func (self *LinuxX86CodeGenerator) picThunkSymbol(reg bs_core.IRegister) bs_core.ISymbol {
  return bs_asm.NewNamedSymbol("__i686.get_pc_thunk." + reg.(*X86Register).GetBaseName())
}

func (self *LinuxX86CodeGenerator) picThunk(file *LinuxX86AssemblyCode, reg bs_core.IRegister) {
//sym := self.picThunkSymbol(reg)
  panic("not implemented")
}

func (self *LinuxX86CodeGenerator) compileFunctionBody(file *LinuxX86AssemblyCode, fun *bs_entity.DefinedFunction) {
  self.errorHandler.Warnf("FIXME: CodeGenerator#compileFunctionBody not implemented: %s", fun.GetName())
}

func (self *LinuxX86CodeGenerator) loadConstant(node bs_core.IExpr, reg bs_core.IRegister) {
  panic("not implemented")
}

func (self *LinuxX86CodeGenerator) loadVariable(v *bs_ir.Var, dest bs_core.IRegister) {
  panic("not implemented")
}

func (self *LinuxX86CodeGenerator) loadAddress(v bs_core.IEntity, dest bs_core.IRegister) {
  panic("not implemented")
}

func (self *LinuxX86CodeGenerator) ax() bs_core.IRegister {
  return NewX86Register(X86_AX, self.naturalType)
}

func (self *LinuxX86CodeGenerator) axT(t int) bs_core.IRegister {
  return NewX86Register(X86_AX, t)
}

func (self *LinuxX86CodeGenerator) bx() bs_core.IRegister {
  return NewX86Register(X86_BX, self.naturalType)
}

func (self *LinuxX86CodeGenerator) bxT(t int) bs_core.IRegister {
  return NewX86Register(X86_BX, t)
}

func (self *LinuxX86CodeGenerator) cx() bs_core.IRegister {
  return NewX86Register(X86_CX, self.naturalType)
}

func (self *LinuxX86CodeGenerator) cxT(t int) bs_core.IRegister {
  return NewX86Register(X86_CX, t)
}

func (self *LinuxX86CodeGenerator) dx() bs_core.IRegister {
  return NewX86Register(X86_DX, self.naturalType)
}

func (self *LinuxX86CodeGenerator) dxT(t int) bs_core.IRegister {
  return NewX86Register(X86_DX, t)
}

func (self *LinuxX86CodeGenerator) si() bs_core.IRegister {
  return NewX86Register(X86_SI, self.naturalType)
}

func (self *LinuxX86CodeGenerator) di() bs_core.IRegister {
  return NewX86Register(X86_DI, self.naturalType)
}

func (self *LinuxX86CodeGenerator) bp() bs_core.IRegister {
  return NewX86Register(X86_BP, self.naturalType)
}

func (self *LinuxX86CodeGenerator) sp() bs_core.IRegister {
  return NewX86Register(X86_SP, self.naturalType)
}

func (self *LinuxX86CodeGenerator) mem1(sym bs_core.ISymbol) *bs_asm.DirectMemoryReference {
  return bs_asm.NewDirectMemoryReference(sym)
}

func (self *LinuxX86CodeGenerator) mem2(reg bs_core.IRegister) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(bs_asm.NewIntegerLiteral(0), reg)
}

func (self *LinuxX86CodeGenerator) mem3(offset int64, reg bs_core.IRegister) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(bs_asm.NewIntegerLiteral(offset), reg)
}

func (self *LinuxX86CodeGenerator) mem4(offset bs_core.ISymbol, reg bs_core.IRegister) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(offset, reg)
}

func (self *LinuxX86CodeGenerator) imm1(n int64) *bs_asm.ImmediateValue {
  return bs_asm.NewImmediateValue(bs_asm.NewIntegerLiteral(n))
}

func (self *LinuxX86CodeGenerator) imm2(lit bs_core.ILiteral) *bs_asm.ImmediateValue {
  return bs_asm.NewImmediateValue(lit)
}
