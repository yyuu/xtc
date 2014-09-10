package main

import (
  "bufio"
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/compiler"
  "bitbucket.org/yyuu/bs/parser"
)

var flagSet = flag.NewFlagSet(os.Args[0], 1)
var dump = flagSet.Bool("D", false, "dump mode")
var verbose = flagSet.Int("v", 0, "verbose mode")

func main() {
  flagSet.Parse(os.Args[1:])
  parser.Verbose = *verbose
  compiler.Verbose = *verbose

  files := flagSet.Args()
  if 0 < len(files) {
    for i := range files {
      cs, err := ioutil.ReadFile(files[i])
      if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }
      ep(parser.ParseExpr(string(cs)))
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
      ep(parser.ParseExpr(s))
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

func ep(a *ast.AST, err error) *ast.AST {
  if err != nil {
    panic(err)
  }

  if *dump == true {
    d(a)
  }

  compiler.NewLocalResolver().Resolve(a)
//compiler.NewTypeResolver(types).Resolve(a)
//types.SemanticCheck()
//compiler.NewDereferenceChecker(types).Check(a)
//compiler.NewTypeChecker(types).Check(a)

  // TODO: evaluate AST
  fmt.Fprintln(os.Stdout, a)
  return a
}

func d(a *ast.AST) *ast.AST {
  cs, err := json.MarshalIndent(a, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Fprintln(os.Stderr, string(cs))
  return a
}
