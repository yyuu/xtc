package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  bs_compiler "bitbucket.org/yyuu/bs/compiler"
)

func main() {
  compiler := bs_compiler.NewCompiler(os.Args[0], os.Args[1:])
  files := compiler.SourceFiles()
  if 0 < len(files) {
    compiler.Compile()
  } else {
    repl(compiler)
  }
}

func repl(compiler *bs_compiler.Compiler) {
  defer func() {
    if s := recover(); s != nil {
      fmt.Fprintf(os.Stderr, "recovered: %s\n", s)
      repl(compiler)
    }
  }()
  stdin  := bufio.NewReader(os.Stdin)
  stdout := bufio.NewWriter(os.Stdout)
  for {
    stdout.WriteString("> ")
    stdout.Flush()
    source, err := stdin.ReadString('\n')
    if err != nil {
      os.Exit(1)
    }
    if strings.TrimSpace(source) != "" {
      compiler.CompileString(source)
    }
  }
}
