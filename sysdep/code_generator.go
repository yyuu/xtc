package sysdep

import (
  "fmt"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
)

type CodeGenerator interface {
  Generate(ir *bs_ir.IR) AssemblyCode
}

func NewCodeGeneratorFor(errorHandler *bs_core.ErrorHandler, options *bs_core.Options, platformId int) CodeGenerator {
  switch platformId {
    case bs_core.PLATFORM_LINUX_X86: {
      return NewLinuxX86CodeGenerator(errorHandler, options)
    }
    default: {
      panic(fmt.Errorf("unknown platformId %d", platformId))
    }
  }
}
