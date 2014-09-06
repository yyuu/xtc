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
  "bitbucket.org/yyuu/bs/parser"
)

var flagSet = flag.NewFlagSet(os.Args[0], 1)
var dump = flagSet.Bool("d", true, "dump mode")
var verbose = flagSet.Bool("v", false, "verbose mode")

func main() {
  flagSet.Parse(os.Args[1:])
  parser.VERBOSE = *verbose

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
    out.WriteString("> ")
    out.Flush()
    s, err := in.ReadString('\n')
    if err != nil {
      break
    }
    if strings.TrimSpace(s) != "" {
      ep(parser.ParseExpr(s))
    }
  }
}

func ep(ast *ast.AST, err error) *ast.AST {
  if err != nil {
    panic(err)
  }
  if *dump == true {
    d(ast)
  }
  // TODO: evaluate AST
  fmt.Println(ast)
  return ast
}

func d(ast *ast.AST) *ast.AST {
  cs, err := json.MarshalIndent(ast, "", "  ")
  if err != nil {
    panic(err)
  }
  fmt.Println(string(cs))
  return ast
}
