package sysdep

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
)

type CodeGenerator struct {
  errorHandler *core.ErrorHandler
}

func NewCodeGenerator(errorHandler *core.ErrorHandler) *CodeGenerator {
  return &CodeGenerator { errorHandler }
}

func NewCodeGeneratorFor(errorHandler *core.ErrorHandler, platform string) *CodeGenerator {
  switch platform {
    case "x86-linux": return NewCodeGenerator(errorHandler)
    default: panic(fmt.Errorf("unknown platform: %s", platform))
  }
}

func (self *CodeGenerator) Generate(ir *bs_ir.IR) *AssemblyCode {
  self.errorHandler.Debug("starting code generator.")
  self.locateSymbols(ir)
  x := self.generateAssemblyCode(ir)
  self.errorHandler.Debug("finished code generator.")
  return x
}

func (self *CodeGenerator) locateSymbols(ir *bs_ir.IR) {
  self.errorHandler.Warn("FIXME: CodeGenerate#localSymbols not implemented")
}

func (self *CodeGenerator) generateAssemblyCode(ir *bs_ir.IR) *AssemblyCode {
  file := NewAssemblyCode()
  return file
}
