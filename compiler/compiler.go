package compiler

import (
  "encoding/json"
  "fmt"
  "os/exec"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
  bs_parser "bitbucket.org/yyuu/bs/parser"
  bs_sysdep "bitbucket.org/yyuu/bs/sysdep"
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

func (self *Compiler) Compile() {
  _, err := self.phase1(self.SourceFiles())
  if err != nil {
    self.errorHandler.Fatal(err)
    return
  }
}

func (self *Compiler) CompileString(s string) {
  src, err := bs_core.NewTemporarySourceFile("", bs_core.EXT_PROGRAM_SOURCE, []byte(s))
  if err != nil {
    self.errorHandler.Fatal(err)
    return
  }
  _, err = self.phase1([]*bs_core.SourceFile { src })
  if err != nil {
    self.errorHandler.Fatal(err)
    return
  }
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
  self.dumpAST(ast)
  types := bs_typesys.NewTypeTableFor(self.options.TargetPlatform())
  sem, err := self.semanticAnalyze(ast, types)
  if err != nil {
    return nil, err
  }
  self.dumpSemant(sem)
  ir, err := NewIRGenerator(self.errorHandler, self.options, types).Generate(sem)
  if err != nil {
    return nil, err
  }
  self.dumpIR(ir)
  asm, err := self.generateAssembly(ir)
  if err != nil {
    return nil, err
  }
  self.dumpAsm(asm)
  self.printAsm(asm)
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

func (self *Compiler) generateAssembly(ir *bs_ir.IR) (bs_sysdep.AssemblyCode, error) {
  code_generator := bs_sysdep.NewCodeGeneratorFor(self.errorHandler, self.options, self.options.TargetPlatform())
  return code_generator.Generate(ir)
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
  return dst, nil
}

func (self *Compiler) dumpAST(ast *bs_ast.AST) {
  if ! self.options.DumpAST() {
    return
  }
  cs, err := json.MarshalIndent(ast, "", "  ")
  if err != nil {
    self.errorHandler.Fatal(err)
  }
  fmt.Println("// AST")
  fmt.Println(string(cs))
}

func (self *Compiler) dumpSemant(ast *bs_ast.AST) {
  if ! self.options.DumpSemantic() {
    return
  }
  cs, err := json.MarshalIndent(ast, "", "  ")
  if err != nil {
    self.errorHandler.Fatal(err)
  }
  fmt.Println("// Semantics")
  fmt.Println(string(cs))
}

func (self *Compiler) dumpIR(ir *bs_ir.IR) {
  if ! self.options.DumpIR() {
    return
  }
  cs, err := json.MarshalIndent(ir, "", "  ")
  if err != nil {
    self.errorHandler.Fatal(err)
  }
  fmt.Println("// IR")
  fmt.Println(string(cs))
}

func (self *Compiler) dumpAsm(asm bs_sysdep.AssemblyCode) {
  if ! self.options.DumpAsm() {
    return
  }
  cs, err := json.MarshalIndent(asm, "", "  ")
  if err != nil {
    self.errorHandler.Fatal(err)
  }
  fmt.Println("// Asm")
  fmt.Println(string(cs))
}

func (self *Compiler) printAsm(asm bs_sysdep.AssemblyCode) {
  if ! self.options.PrintAsm() {
    return
  }
  fmt.Println(asm.ToSource())
}
