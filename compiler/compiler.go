package compiler

import (
  "encoding/json"
  "fmt"
  "os"
  "os/exec"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
  bs_x86_linux "bitbucket.org/yyuu/bs/x86/linux"
  bs_parser "bitbucket.org/yyuu/bs/parser"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
)

type Compiler struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
}

func NewCompiler(name string, args []string) *Compiler {
  options := bs_core.ParseOptions(name, args)
  var logLevel = bs_core.LOG_INFO
  if options.IsVerboseMode() {
    logLevel = bs_core.LOG_DEBUG
  }
  errorHandler := bs_core.NewErrorHandler(logLevel)
  return &Compiler { errorHandler, options }
}

func (self *Compiler) SourceFiles() []*bs_core.SourceFile {
  files := self.options.SourceFiles()
  sources := make([]*bs_core.SourceFile, len(files))
  for i := range files {
    sources[i] = bs_core.NewSourceFile(files[i], files[i], "")
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
  src, err := bs_core.NewTemporarySourceFile("", bs_core.EXT_PROGRAM_SOURCE, []byte(s))
  if err != nil {
    self.errorHandler.Fatal(err)
    return "", err
  }
  dst, err := self.phase1([]*bs_core.SourceFile { src })
  if err != nil {
    self.errorHandler.Fatal(err)
    return "", err
  }
  return dst.GetPath(), nil
}

func (self *Compiler) phase1(sources []*bs_core.SourceFile) (*bs_core.SourceFile, error) {
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

func (self *Compiler) phase2(src *bs_core.SourceFile) (*bs_core.SourceFile, error) {
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

func (self *Compiler) phase3(src *bs_core.SourceFile) (*bs_core.SourceFile, error) {
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
  } else {
    return nil, fmt.Errorf("not an object file: %s", src)
  }
}

func (self *Compiler) compile(sources []*bs_core.SourceFile) (*bs_core.SourceFile, error) {
  if len(sources) < 1 {
    return nil, fmt.Errorf("no program sources given")
  }
  dst := sources[0].ToAssemblySource()
  var ast *bs_ast.AST
  for i := range sources {
    parsed, err := bs_parser.Parse(sources[i], self.errorHandler, self.options)
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
    return dst, nil
  }
  types := bs_typesys.NewTypeTableFor(self.options.TargetPlatform())
  sem, err := self.semanticAnalyze(ast, types)
  if err != nil {
    return nil, err
  }
  if self.options.DumpSemantic() {
    self.dumpSemant(sem)
    return dst, nil
  }
  ir, err := NewIRGenerator(self.errorHandler, self.options, types).Generate(sem)
  if err != nil {
    return nil, err
  }
  if self.options.DumpIR() {
    self.dumpIR(ir)
    return dst, nil
  }
  asm, err := self.generateAssembly(ir)
  if err != nil {
    return nil, err
  }
  if self.options.DumpAsm() {
    self.dumpAsm(asm)
    return dst, nil
  }
  if self.options.PrintAsm() {
    self.printAsm(asm)
    return dst, nil
  }
  dst.WriteAll([]byte(asm.ToSource()))
  return dst, nil
}

func (self *Compiler) semanticAnalyze(ast *bs_ast.AST, types *bs_typesys.TypeTable) (*bs_ast.AST, error) {
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

func (self *Compiler) generateAssembly(ir *bs_ir.IR) (bs_core.IAssemblyCode, error) {
  switch self.options.TargetPlatform() {
    case bs_core.PLATFORM_X86_LINUX: {
      return bs_x86_linux.NewCodeGenerator(self.errorHandler, self.options).Generate(ir)
    }
    default: {
      return nil, fmt.Errorf("unknown platform: %d", self.options.TargetPlatform())
    }
  }
}

func (self *Compiler) assemble(src *bs_core.SourceFile) (*bs_core.SourceFile, error) {
  dst := src.ToObjectFile()
  err := exec.Command("/usr/bin/as", "--32", "-o", fmt.Sprint(dst), fmt.Sprint(src)).Run()
  if err != nil {
    return nil, err
  }
  return dst, nil
}

func (self *Compiler) link(src *bs_core.SourceFile) (*bs_core.SourceFile, error) {
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

func (self *Compiler) dumpAST(ast *bs_ast.AST) {
  self.dump("AST", ast)
}

func (self *Compiler) dumpSemant(ast *bs_ast.AST) {
  self.dump("Semantics", ast)
}

func (self *Compiler) dumpIR(ir *bs_ir.IR) {
  self.dump("IR", ir)
}

func (self *Compiler) dumpAsm(asm bs_core.IAssemblyCode) {
  self.dump("Asm", asm)
}

func (self *Compiler) printAsm(asm bs_core.IAssemblyCode) {
  fmt.Println(asm.ToSource())
}
