package sysdep

import (
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
)

type X86CodeGenerator struct {
  errorHandler *bs_core.ErrorHandler
}

func NewX86CodeGenerator(errorHandler *bs_core.ErrorHandler) *X86CodeGenerator {
  return &X86CodeGenerator { errorHandler }
}

func (self *X86CodeGenerator) Generate(ir *bs_ir.IR) IAssemblyCode {
  self.errorHandler.Debug("starting code generator.")
  self.locateSymbols(ir)
  x := self.generateAssemblyCode(ir)
  self.errorHandler.Debug("finished code generator.")
  return x
}

func (self *X86CodeGenerator) locateSymbols(ir *bs_ir.IR) {
  self.errorHandler.Warn("FIXME* X86CodeGenerater#localSymbols not implemented")
}

func (self *X86CodeGenerator) generateAssemblyCode(ir *bs_ir.IR) IAssemblyCode {
  file := NewX86AssemblyCode()
  return file
}
