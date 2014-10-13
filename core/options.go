package core

import (
  "flag"
  "os"
  "os/exec"
  "strings"
)

type assemblerOptions struct {
}

func newAssemblerOptions(flagSet *flag.FlagSet) *assemblerOptions {
  return &assemblerOptions { }
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
  targetPlatform int
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
  compile *bool
  assemble *bool
  link *bool
  output *string
  verbose *bool
}

func NewOptions(name string) *Options {
  flagSet := flag.NewFlagSet(name, flag.ExitOnError)
  return &Options {
    detectTargetPlatform(),
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
    flagSet.Bool("S", false, "S"), // compile
    flagSet.Bool("c", false, "c"), // assemble
    flagSet.Bool("link", false, "link"), // link
    flagSet.String("o", "", "o"),
    flagSet.Bool("verbose", false, "verbose"),
  }
}

func ParseOptions(name string, args []string) *Options {
  return NewOptions(name).Parse(args)
}

func detectTargetPlatform() int {
  env_arch := strings.ToLower(strings.TrimSpace(os.Getenv("XTCARCH")))
  env_os := strings.ToLower(strings.TrimSpace(os.Getenv("XTCOS")))
  switch env_arch {
    case "amd64": switch env_os {
      case "linux": return PLATFORM_AMD64_LINUX
      default: {
        // FIXME: detect current operating system
        return PLATFORM_AMD64_LINUX
      }
    }
    case "x86": switch env_os {
      case "linux": return PLATFORM_X86_LINUX
      default: {
        // FIXME: detect current operating system
        return PLATFORM_X86_LINUX
      }
    }
    default: {
//    return detectSourcePlatform()
      return PLATFORM_X86_LINUX // FIXME: only x86-linux has been implemented
    }
  }
}

func detectSourcePlatform() int {
  // FIXME: should support platforms other than POSIX
  cmd_m := exec.Command("uname", "-m")
  out_m, err := cmd_m.Output()
  if err != nil {
    panic(err)
  }

  cmd_s := exec.Command("uname", "-s")
  out_s, err := cmd_s.Output()
  if err != nil {
    panic(err)
  }

  uname_m := strings.ToLower(strings.TrimSpace(string(out_m)))
  uname_s := strings.ToLower(strings.TrimSpace(string(out_s)))
  switch uname_m {
    case "amd64": fallthrough
    case "x86_64": switch uname_s {
      case "linux": return PLATFORM_AMD64_LINUX
      default: {
        panic("unknown operating system: " + uname_s)
      }
    }
    case "i386": fallthrough
    case "i486": fallthrough
    case "i586": fallthrough
    case "i686": fallthrough
    case "i786": fallthrough
    case "x86": switch uname_s {
      case "linux": return PLATFORM_X86_LINUX
      default: {
        panic("unknown operating system: " + uname_s)
      }
    }
    default: {
      panic("unknown architecture: " + uname_m)
    }
  }
}

func (self *Options) Parse(args []string) *Options {
  self.flagSet.Parse(args)
  switch {
    case *self.link: {
      *self.compile = true
      *self.assemble = true
    }
    case *self.assemble: {
      *self.compile = true
      *self.link = false
    }
    case *self.compile: {
      *self.assemble = false
      *self.link = false
    }
    default: {
      *self.compile = true
      *self.assemble = true
      *self.link = true
    }
  }
  return self
}

func (self *Options) IsCompileRequired() bool {
  return *self.compile
}

func (self *Options) IsAssembleRequired() bool {
  return *self.assemble
}

func (self *Options) IsLinkRequired() bool {
  return *self.link
}

func (self *Options) IsVerboseMode() bool {
  return *self.verbose
}

func (self *Options) IsVerboseAsm() bool {
  return *self.codeGeneratorOptions.verboseAsm
}

func (self *Options) TargetPlatform() int {
  return self.targetPlatform
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

func (self *Options) GetLibraryPath() []string {
  var libraryPath []string
  environ := os.Getenv("XTCPATH")
  if 0 < len(environ) {
    libraryPath = append(libraryPath, environ)
  }
  libraryPath = append(libraryPath, "xtcpath")
  return libraryPath
}

func (self *Options) OutputFilename() string {
  return *self.output
}
