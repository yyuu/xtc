package core

import (
  "flag"
)

type assemblerOptions struct {
}

func newAssemblerOptions(flagSet *flag.FlagSet) *assemblerOptions {
  return &assemblerOptions {
  }
}

type codeGeneratorOptions struct {
  optimizeLevel *int
  generatePIC *bool
  generatePIE *bool
  verboseAsm *bool
}

func newCodeGeneratorOptions(flagSet *flag.FlagSet) *codeGeneratorOptions {
  return &codeGeneratorOptions {
    flagSet.Int("O", 0, "O"),
    flagSet.Bool("fpic", false, "fpic"),
    flagSet.Bool("fpie", false, "fpie"),
    flagSet.Bool("fverbose-asm", false, "fverbose-asm"),
  }
}

type linkerOptions struct {
  generateSharedLibrary *bool
  generatePIE *bool
  noStartFiles *bool
  noDefaultLibs *bool
}

func newLinkerOptions(flagSet *flag.FlagSet) *linkerOptions {
  return &linkerOptions {
    flagSet.Bool("shared", false, "shared"),
    flagSet.Bool("pie", false, "pie"),
    flagSet.Bool("nostartfiles", false, "nostartfiles"),
    flagSet.Bool("nodefaultlibs", false, "nodefaultlibs"),
  }
}

type Options struct {
  flagSet *flag.FlagSet
  *assemblerOptions
  *codeGeneratorOptions
  *linkerOptions
  checkSyntax *bool
  dumpTokens *bool
  dumpAST *bool
  dumpStmt *bool
  dumpExpr *bool
  dumpSemantic *bool
  dumpReference *bool
  dumpIR *bool
  dumpAsm *bool
  printAsm *bool
  verbose *bool
}

func NewOptions(name string) *Options {
  flagSet := flag.NewFlagSet(name, flag.ExitOnError)
  return &Options {
    flagSet,
    newAssemblerOptions(flagSet),
    newCodeGeneratorOptions(flagSet),
    newLinkerOptions(flagSet),
    flagSet.Bool("check-syntax", false, "check syntax"),
    flagSet.Bool("dump-tokens", false, "dump tokens"),
    flagSet.Bool("dump-ast", false, "dump ast"),
    flagSet.Bool("dump-stmt", false, "dump stmt"),
    flagSet.Bool("dump-expr", false, "dump expr"),
    flagSet.Bool("dump-semantic", false, "dump semantic"),
    flagSet.Bool("dump-reference", false, "dump reference"),
    flagSet.Bool("dump-ir", false, "dump ir"),
    flagSet.Bool("dump-asm", false, "dump asm"),
    flagSet.Bool("print-asm", false, "print asm"),
    flagSet.Bool("verbose", false, "verbose"),
  }
}

func ParseOptions(name string, args []string) *Options {
  return NewOptions(name).Parse(args)
}

func (self *Options) Parse(args []string) *Options {
  self.flagSet.Parse(args)
  return self
}

func (self *Options) IsAssembleRequired() bool {
  panic("not implemented")
}

func (self *Options) IsLinkRequired() bool {
  panic("not implemented")
}

func (self *Options) IsVerboseMode() bool {
  return *self.verbose
}

func (self *Options) IsVerboseAsm() bool {
  return *self.codeGeneratorOptions.verboseAsm
}

func (self *Options) TargetPlatform() int {
  return PLATFORM_LINUX_X86
}

func (self *Options) IsGenratingSharedLibrary() bool {
  return *self.linkerOptions.generateSharedLibrary
}

func (self *Options) SourceFiles() []string {
  return self.flagSet.Args()
}

func (self *Options) CheckSyntax() bool {
  return *self.checkSyntax
}

func (self *Options) DumpTokens() bool {
  return *self.dumpTokens
}

func (self *Options) DumpAST() bool {
  return *self.dumpAST
}

func (self *Options) DumpStmt() bool {
  return *self.dumpStmt
}

func (self *Options) DumpExpr() bool {
  return *self.dumpExpr
}

func (self *Options) DumpSemantic() bool {
  return *self.dumpSemantic
}

func (self *Options) DumpReference() bool {
  return *self.dumpReference
}

func (self *Options) DumpIR() bool {
  return *self.dumpIR
}

func (self *Options) DumpAsm() bool {
  return *self.dumpAsm
}

func (self *Options) PrintAsm() bool {
  return *self.printAsm
}

func (self *Options) IsPositionIndependent() bool {
  return *self.codeGeneratorOptions.generatePIC || *self.codeGeneratorOptions.generatePIE
}

func (self *Options) IsPICRequired() bool {
  return *self.codeGeneratorOptions.generatePIC
}

func (self *Options) IsPIERequired() bool {
  return *self.codeGeneratorOptions.generatePIE
}
