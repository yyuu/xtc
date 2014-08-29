package main

import (
  "bufio"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/parser"
)

var flagSet = flag.NewFlagSet(os.Args[0], 1)
var verbose = flagSet.Bool("v", false, "verbose mode")

func main() {
  flagSet.Parse(os.Args[1:])

  files := flagSet.Args()
  if 0 < len(files) {
    for i := range files {
      cs, err := ioutil.ReadFile(files[i])
      if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }
      p(parser.ParseExpr(string(cs)))
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
      p(parser.ParseExpr(s))
    }
  }
}

func p(ast *ast.AST, err error) {
  if err != nil {
    panic(err)
  }
  fmt.Print(*ast)
}
