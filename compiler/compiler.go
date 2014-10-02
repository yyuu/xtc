package compiler

import (
  "encoding/json"
  "fmt"
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
  return &Compiler {
    errorHandler: errorHandler,
    options: options,
  }
}

func (self *Compiler) SourceFiles() []string {
  return self.options.SourceFiles()
}

func (self *Compiler) Compile() {
  files := self.SourceFiles()
  for i := range files {
    ast, err := bs_parser.ParseFile(files[i], self.errorHandler, self.options)
    if err != nil { self.errorHandler.Fatal(err) }
    self.build(ast)
  }
}

func (self *Compiler) CompileString(source string) {
  ast, err := bs_parser.ParseExpr(source, self.errorHandler, self.options)
  if err != nil { self.errorHandler.Fatal(err) }
  self.build(ast)
}

func (self *Compiler) build(ast *bs_ast.AST) {
  self.dumpAST(ast)
  types := bs_typesys.NewTypeTableFor(self.options.TargetPlatform())
  sem, err := self.semanticAnalyze(ast, types)
  if err != nil { self.errorHandler.Fatal(err) }
  self.dumpSemant(sem)
  ir := NewIRGenerator(self.errorHandler, self.options, types).Generate(sem)
  self.dumpIR(ir)
  asm, err := self.generateAssembly(ir)
  if err != nil { self.errorHandler.Fatal(err) }
  self.dumpAsm(asm)
  self.printAsm(asm)
}

func (self *Compiler) semanticAnalyze(ast *bs_ast.AST, types *bs_typesys.TypeTable) (*bs_ast.AST, error) {
  NewLocalResolver(self.errorHandler, self.options).Resolve(ast)
  NewTypeResolver(self.errorHandler, self.options, types).Resolve(ast)
  types.SemanticCheck(self.errorHandler)
  NewDereferenceChecker(self.errorHandler, self.options, types).Check(ast)
  NewTypeChecker(self.errorHandler, self.options, types).Check(ast)
  return ast, nil
}

func (self *Compiler) generateAssembly(ir *bs_ir.IR) (bs_sysdep.AssemblyCode, error) {
  code_generator := bs_sysdep.NewCodeGeneratorFor(self.errorHandler, self.options, self.options.TargetPlatform())
  return code_generator.Generate(ir), nil
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
