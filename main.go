package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "strings"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_compiler "bitbucket.org/yyuu/bs/compiler"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_ir "bitbucket.org/yyuu/bs/ir"
  bs_parser "bitbucket.org/yyuu/bs/parser"
  bs_sysdep "bitbucket.org/yyuu/bs/sysdep"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
)

var errorHandler = bs_core.NewErrorHandler(bs_core.LOG_DEBUG)

func main() {
  options := bs_core.ParseOptions(os.Args[0], os.Args[1:])
  files := options.SourceFiles()
  if 0 < len(files) {
    for i := range files {
      ast, err := bs_parser.ParseFile(files[i], errorHandler, options)
      ep(ast, err, options)
    }
  } else {
    repl(options)
  }
}

func repl(options *bs_core.Options) {
  defer func() {
    if s := recover(); s != nil {
      fmt.Fprintf(os.Stderr, "recovered: %s\n", s)
      repl(options)
    }
  }()
  in  := bufio.NewReader(os.Stdin)
  out := bufio.NewWriter(os.Stdout)
  for {
    s := r(in, out, options)
    if s != "" {
      ast, err := bs_parser.ParseExpr(s, errorHandler, options)
      ep(ast, err, options)
    }
  }
}

func r(in *bufio.Reader, out *bufio.Writer, options *bs_core.Options) string {
  out.WriteString("> ")
  out.Flush()
  s, err := in.ReadString('\n')
  if err != nil {
    os.Exit(1)
  }
  return strings.TrimSpace(s)
}

func ep(ast *bs_ast.AST, err error, options *bs_core.Options) {
  if err != nil {
    panic(err)
  }
  dumpAST(ast, options)
  types := bs_typesys.NewTypeTableFor(options.TargetPlatform())
  sem := semanticAnalyze(ast, types, options)
  dumpSemant(sem, options)
  ir := bs_compiler.NewIRGenerator(errorHandler, options, types).Generate(sem)
  dumpIR(ir, options)
  asm := generateAssembly(ir, options)
  dumpAsm(asm, options)
  printAsm(asm, options)
}

func semanticAnalyze(ast *bs_ast.AST, types *bs_typesys.TypeTable, options *bs_core.Options) *bs_ast.AST {
  bs_compiler.NewLocalResolver(errorHandler, options).Resolve(ast)
  bs_compiler.NewTypeResolver(errorHandler, options, types).Resolve(ast)
  types.SemanticCheck(errorHandler)
  bs_compiler.NewDereferenceChecker(errorHandler, options, types).Check(ast)
  bs_compiler.NewTypeChecker(errorHandler, options, types).Check(ast)
  return ast
}

func generateAssembly(ir *bs_ir.IR, options *bs_core.Options) bs_sysdep.AssemblyCode {
  code_generator := bs_sysdep.NewCodeGeneratorFor(errorHandler, options, options.TargetPlatform())
  return code_generator.Generate(ir)
}

func dumpAST(ast *bs_ast.AST, options *bs_core.Options) {
  if ! options.DumpAST() {
    return
  }
  cs, err := json.MarshalIndent(ast, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// AST")
  fmt.Println(string(cs))
}

func dumpSemant(ast *bs_ast.AST, options *bs_core.Options) {
  if ! options.DumpSemantic() {
    return
  }
  cs, err := json.MarshalIndent(ast, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// Semantics")
  fmt.Println(string(cs))
}

func dumpIR(ir *bs_ir.IR, options *bs_core.Options) {
  if ! options.DumpIR() {
    return
  }
  cs, err := json.MarshalIndent(ir, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// IR")
  fmt.Println(string(cs))
}

func dumpAsm(asm bs_sysdep.AssemblyCode, options *bs_core.Options) {
  if ! options.DumpAsm() {
    return
  }
  cs, err := json.MarshalIndent(asm, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// Asm")
  fmt.Println(string(cs))
}

func printAsm(asm bs_sysdep.AssemblyCode, options *bs_core.Options) {
  if ! options.PrintAsm() {
    return
  }
  fmt.Println(asm.ToSource())
}
