package sysdep

import (
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
)

type LinuxX86CodeGenerator struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
}

func NewLinuxX86CodeGenerator(errorHandler *bs_core.ErrorHandler, options *bs_core.Options) *LinuxX86CodeGenerator {
  return &LinuxX86CodeGenerator { errorHandler, options }
}

func (self *LinuxX86CodeGenerator) Generate(ir *bs_ir.IR) AssemblyCode {
  self.errorHandler.Debug("starting code generator.")
  self.locateSymbols(ir)
  x := self.generateAssemblyCode(ir)
  self.errorHandler.Debug("finished code generator.")
  return x
}

func (self *LinuxX86CodeGenerator) locateSymbols(ir *bs_ir.IR) {
  self.errorHandler.Warn("FIXME* X86CodeGenerater#localSymbols not implemented")
}

func (self *LinuxX86CodeGenerator) generateAssemblyCode(ir *bs_ir.IR) AssemblyCode {
  file := NewLinuxX86AssemblyCode()
  return file
}
