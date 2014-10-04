package compiler

import (
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_ir "bitbucket.org/yyuu/xtc/ir"
  xtc_x86_linux "bitbucket.org/yyuu/xtc/x86/linux"
  xtc_parser "bitbucket.org/yyuu/xtc/parser"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
)

type Compiler struct {
  errorHandler *xtc_core.ErrorHandler
  options *xtc_core.Options
}

func NewCompiler(name string, args []string) *Compiler {
  options := xtc_core.ParseOptions(name, args)
  var logLevel = xtc_core.LOG_INFO
  if options.IsVerboseMode() {
    logLevel = xtc_core.LOG_DEBUG
  }
  errorHandler := xtc_core.NewErrorHandler(logLevel)
  return &Compiler { errorHandler, options }
}

func (self *Compiler) SourceFiles() []*xtc_core.SourceFile {
  files := self.options.SourceFiles()
  sources := make([]*xtc_core.SourceFile, len(files))
  for i := range files {
    sources[i] = xtc_core.NewSourceFile(files[i], files[i], "")
  }
  return sources
}

func (self *Compiler) Compile() (string, error) {
  dst, err := self.phase1(self.SourceFiles())
  if err != nil {
    self.errorHandler.Fatal(err)
    return "", err
  }
  path := dst.GetPath()
  if self.options.OutputFilename() != "" {
    err := os.Rename(path, self.options.OutputFilename())
    if err != nil {
      self.errorHandler.Fatal(err)
      return "", err
    }
    path = self.options.OutputFilename()
  }
  return path, nil
}

func (self *Compiler) CompileString(s string) (string, error) {
  src, err := xtc_core.NewTemporarySourceFile("", xtc_core.EXT_PROGRAM_SOURCE, []byte(s))
  if err != nil {
    self.errorHandler.Fatal(err)
    return "", err
  }
  dst, err := self.phase1([]*xtc_core.SourceFile { src })
  if err != nil {
    self.errorHandler.Fatal(err)
    return "", err
  }
  return dst.GetPath(), nil
}

func (self *Compiler) phase1(sources []*xtc_core.SourceFile) (*xtc_core.SourceFile, error) {
  if len(sources) < 1 {
    return nil, fmt.Errorf("no program sources given")
  }
  if ! self.options.IsCompileRequired() {
    return sources[0], nil
  }
  if sources[0].IsProgramSource() {
    defer func() {
      for i := range sources {
        if sources[i].IsGenerated() {
          self.errorHandler.Debugf("Remove temporary file: %s", sources[i])
          sources[i].Remove()
        }
      }
    }()
    dst, err := self.compile(sources)
    if err != nil {
      return nil, err
    }
    return self.phase2(dst)
  } else {
    return self.phase2(sources[0])
  }
}

func (self *Compiler) phase2(src *xtc_core.SourceFile) (*xtc_core.SourceFile, error) {
  if ! self.options.IsAssembleRequired() {
    return src, nil
  }
  if src.IsAssemblySource() {
    defer func() {
      if src.IsGenerated() {
        self.errorHandler.Debugf("Remove temporary file: %s", src)
        src.Remove()
      }
    }()
    dst, err := self.assemble(src)
    if err != nil {
      return nil, err
    }
    return self.phase3(dst)
  } else {
    return self.phase3(src)
  }
}

func (self *Compiler) phase3(src *xtc_core.SourceFile) (*xtc_core.SourceFile, error) {
  if ! self.options.IsLinkRequired() {
    return src, nil
  }
  if src.IsObjectFile() {
    defer func() {
      if src.IsGenerated() {
        self.errorHandler.Debugf("Remove temporary file: %s", src)
        src.Remove()
      }
    }()
    return self.link(src)
  }
  return src, nil
}

func (self *Compiler) compile(sources []*xtc_core.SourceFile) (*xtc_core.SourceFile, error) {
  if len(sources) < 1 {
    return nil, fmt.Errorf("no program sources given")
  }
  dst := sources[0].ToAssemblySource()
  var ast *xtc_ast.AST
  for i := range sources {
    parsed, err := xtc_parser.Parse(sources[i], self.errorHandler, self.options)
    if err != nil {
      return nil, err
    }
    if ast == nil {
      ast = parsed
    } else {
      decl := ast.GetDeclaration()
      decl.AddDeclaration(parsed.GetDeclaration())
    }
  }
  if self.options.DumpAST() {
    self.dumpAST(ast)
    return sources[0], nil
  }
  types := xtc_typesys.NewTypeTableFor(self.options.TargetPlatform())
  sem, err := self.semanticAnalyze(ast, types)
  if err != nil {
    return nil, err
  }
  if self.options.DumpSemantic() {
    self.dumpSemant(sem)
    return sources[0], nil
  }
  ir, err := NewIRGenerator(self.errorHandler, self.options, types).Generate(sem)
  if err != nil {
    return nil, err
  }
  if self.options.DumpIR() {
    self.dumpIR(ir)
    return sources[0], nil
  }
  asm, err := self.generateAssembly(ir)
  if err != nil {
    return nil, err
  }
  if self.options.DumpAsm() {
    self.dumpAsm(asm)
    return sources[0], nil
  }
  if self.options.PrintAsm() {
    self.printAsm(asm)
    return sources[0], nil
  }
  dst.WriteAll([]byte(asm.ToSource()))
  return dst, nil
}

func (self *Compiler) semanticAnalyze(ast *xtc_ast.AST, types *xtc_typesys.TypeTable) (*xtc_ast.AST, error) {
  sem1, err := NewLocalResolver(self.errorHandler, self.options).Resolve(ast)
  if err != nil {
    return nil, err
  }
  sem2, err := NewTypeResolver(self.errorHandler, self.options, types).Resolve(sem1)
  if err != nil {
    return nil, err
  }
  types.SemanticCheck(self.errorHandler)
  sem3, err := NewDereferenceChecker(self.errorHandler, self.options, types).Check(sem2)
  if err != nil {
    return nil, err
  }
  return NewTypeChecker(self.errorHandler, self.options, types).Check(sem3)
}

func (self *Compiler) generateAssembly(ir *xtc_ir.IR) (xtc_core.IAssemblyCode, error) {
  switch self.options.TargetPlatform() {
    case xtc_core.PLATFORM_X86_LINUX: {
      return xtc_x86_linux.NewCodeGenerator(self.errorHandler, self.options).Generate(ir)
    }
    default: {
      return nil, fmt.Errorf("unknown platform: %d", self.options.TargetPlatform())
    }
  }
}

func (self *Compiler) assemble(src *xtc_core.SourceFile) (*xtc_core.SourceFile, error) {
  dst := src.ToObjectFile()
  err := exec.Command("/usr/bin/as", "--32", "-o", fmt.Sprint(dst), fmt.Sprint(src)).Run()
  if err != nil {
    return nil, err
  }
  return dst, nil
}

func (self *Compiler) link(src *xtc_core.SourceFile) (*xtc_core.SourceFile, error) {
  dst := src.ToExecutableFile()
  self.errorHandler.Warn("FIXME: Compiler#link: not implemented")
  dst.WriteAll([]byte { })
  os.Chmod(dst.GetPath(), 0755) // FIXME: support umask
  return dst, nil
}

func (self *Compiler) dump(name string, x interface{}) {
  bytes, err := json.MarshalIndent(x, "", "  ")
  if err != nil {
    self.errorHandler.Fatal(err)
  }
  fmt.Println("// " + name)
  fmt.Println(string(bytes))
}

func (self *Compiler) dumpAST(ast *xtc_ast.AST) {
  self.dump("AST", ast)
}

func (self *Compiler) dumpSemant(ast *xtc_ast.AST) {
  self.dump("Semantics", ast)
}

func (self *Compiler) dumpIR(ir *xtc_ir.IR) {
  self.dump("IR", ir)
}

func (self *Compiler) dumpAsm(asm xtc_core.IAssemblyCode) {
  self.dump("Asm", asm)
}

func (self *Compiler) printAsm(asm xtc_core.IAssemblyCode) {
  fmt.Println(asm.ToSource())
}
