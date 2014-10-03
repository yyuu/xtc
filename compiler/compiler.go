package compiler

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
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

func (self *Compiler) SourceFiles() []string {
  return self.options.SourceFiles()
}

func (self *Compiler) Compile() {
  files := self.SourceFiles()
  for i := range files {
    err := self.phase1(NewSourceFile(files[i]))
    if err != nil {
      self.errorHandler.Fatal(err)
    }
  }
}

func (self *Compiler) CompileString(s string) {
  tmpdir, err := ioutil.TempDir("/tmp", "xtc.")
  if err != nil {
    self.errorHandler.Fatal(err)
  }
  defer func() {
    err := os.RemoveAll(tmpdir)
    if err != nil {
      self.errorHandler.Fatal(err)
    }
  }()
  tempfile := tmpdir + "/a.xtc"
  err = ioutil.WriteFile(tempfile, []byte(s), 0644)
  if err != nil {
    self.errorHandler.Fatal(err)
  }
  err = self.phase1(NewSourceFile(tempfile))
  if err != nil {
    self.errorHandler.Fatal(err)
  }
}

func (self *Compiler) phase1(src *SourceFile) error {
  if src.IsProgramSource() {
    dst := src.ToAssemblySource()
    err := self.compile(src, dst)
    if err != nil {
      return err
    }
    return self.phase2(dst)
  } else {
    return self.phase2(src)
  }
}

func (self *Compiler) phase2(src *SourceFile) error {
  if src.IsAssemblySource() {
    dst := src.ToObjectFile()
    err := self.assemble(src, dst)
    if err != nil {
      return err
    }
    return self.phase3(dst)
  } else {
    return self.phase3(src)
  }
}

func (self *Compiler) phase3(src *SourceFile) error {
  if src.IsObjectFile() {
    return self.link(src)
  } else {
    return fmt.Errorf("not an object file: %s", src)
  }
}

func (self *Compiler) compile(src *SourceFile, dst *SourceFile) error {
  ast, err := bs_parser.Parse(src, self.errorHandler, self.options)
  if err != nil {
    return err
  }
  self.dumpAST(ast)
  types := bs_typesys.NewTypeTableFor(self.options.TargetPlatform())
  sem, err := self.semanticAnalyze(ast, types)
  if err != nil {
    return err
  }
  self.dumpSemant(sem)
  ir, err := NewIRGenerator(self.errorHandler, self.options, types).Generate(sem)
  if err != nil {
    return err
  }
  self.dumpIR(ir)
  asm, err := self.generateAssembly(ir)
  if err != nil {
    return err
  }
  self.dumpAsm(asm)
  self.printAsm(asm)
  dst.Write([]byte(asm.ToSource()))
  return nil
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

func (self *Compiler) assemble(src *SourceFile, dst *SourceFile) error {
  return fmt.Errorf("not implemented")
}

func (self *Compiler) link(src *SourceFile) error {
  return fmt.Errorf("not implemented")
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
