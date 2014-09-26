package main

import (
  "bufio"
  "encoding/json"
  "flag"
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

var flagSet = flag.NewFlagSet(os.Args[0], 1)
var dump = flagSet.Int("d", 0, "dump mode")
var verbose = flagSet.Int("v", 0, "verbose mode")
var errorHandler = bs_core.NewErrorHandler(bs_core.LOG_DEBUG)

const (
  DUMP_AST = 1<<iota
  DUMP_SEMANT
  DUMP_IR
  DUMP_ASM
)

func main() {
  flagSet.Parse(os.Args[1:])
  bs_parser.Verbose = *verbose
  bs_compiler.Verbose = *verbose

  files := flagSet.Args()
  if 0 < len(files) {
    for i := range files {
      ep(bs_parser.ParseFile(files[i]))
    }
  } else {
    repl()
  }
}

func repl() {
  defer func() {
    if s := recover(); s != nil {
      fmt.Fprintf(os.Stderr, "recovered: %s\n", s)
      repl()
    }
  }()
  in  := bufio.NewReader(os.Stdin)
  out := bufio.NewWriter(os.Stdout)
  for {
    s := r(in, out)
    if s != "" {
      ep(bs_parser.ParseExpr(s))
    }
  }
}

func r(in *bufio.Reader, out *bufio.Writer) string {
  out.WriteString("> ")
  out.Flush()
  s, err := in.ReadString('\n')
  if err != nil {
    os.Exit(1)
  }
  return strings.TrimSpace(s)
}

func ep(ast *bs_ast.AST, err error) {
  if err != nil {
    panic(err)
  }
  if (*dump & DUMP_AST) != 0 {
    dumpAST(ast)
  }
  types := bs_typesys.NewTypeTableFor("x86-linux")
  sem := semanticAnalyze(ast, types)
  if (*dump & DUMP_SEMANT) != 0 {
    dumpSemant(sem)
  }
  ir := bs_compiler.NewIRGenerator(errorHandler, types).Generate(sem)
  if (*dump & DUMP_IR) != 0 {
    dumpIR(ir)
  }
  asm := generateAssembly(ir)
  if (*dump & DUMP_ASM) != 0 {
    dumpAsm(asm)
  }
}

func semanticAnalyze(ast *bs_ast.AST, types *bs_typesys.TypeTable) *bs_ast.AST {
  bs_compiler.NewLocalResolver(errorHandler).Resolve(ast)
  bs_compiler.NewTypeResolver(errorHandler, types).Resolve(ast)
  types.SemanticCheck(errorHandler)
  bs_compiler.NewDereferenceChecker(errorHandler, types).Check(ast)
  bs_compiler.NewTypeChecker(errorHandler, types).Check(ast)
  return ast
}

func generateAssembly(ir *bs_ir.IR) *bs_sysdep.AssemblyCode {
  return bs_sysdep.NewCodeGeneratorFor(errorHandler, "x86-linux").Generate(ir)
}

func dumpAST(ast *bs_ast.AST) {
  cs, err := json.MarshalIndent(ast, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// AST")
  fmt.Println(string(cs))
}

func dumpSemant(ast *bs_ast.AST) {
  cs, err := json.MarshalIndent(ast, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// Semantics")
  fmt.Println(string(cs))
}

func dumpIR(ir *bs_ir.IR) {
  cs, err := json.MarshalIndent(ir, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// IR")
  fmt.Println(string(cs))
}

func dumpAsm(asm *bs_sysdep.AssemblyCode) {
  cs, err := json.MarshalIndent(asm, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println("// Asm")
  fmt.Println(string(cs))
}
