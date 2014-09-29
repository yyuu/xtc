package sysdep

import (
  "fmt"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
)

type ICodeGenerator interface {
  Generate(ir *bs_ir.IR) IAssemblyCode
}

func NewCodeGeneratorFor(errorHandler *bs_core.ErrorHandler, platformId int) ICodeGenerator {
  switch platformId {
    case bs_core.PLATFORM_X86_LINUX: {
      return NewX86CodeGenerator(errorHandler)
    }
    default: {
      panic(fmt.Errorf("unknown platformId %d", platformId))
    }
  }
}
