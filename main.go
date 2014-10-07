package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "strings"
  xtc_compiler "bitbucket.org/yyuu/xtc/compiler"
)

func main() {
  compiler := xtc_compiler.NewCompiler(os.Args[0], os.Args[1:])
  files := compiler.SourceFiles()
  if 0 < len(files) {
    _, err := compiler.Compile()
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
  } else {
    repl(compiler)
  }
}

func repl(compiler *xtc_compiler.Compiler) {
  defer func() {
    if s := recover(); s != nil {
      fmt.Fprintf(os.Stderr, "recovered: %s\n", s)
      repl(compiler)
    }
  }()
  stdin  := bufio.NewReader(os.Stdin)
  stdout := bufio.NewWriter(os.Stdout)

  var sources []string
  for {
    stdout.WriteString("> ")
    stdout.Flush()
    source, err := stdin.ReadString('\n')
    if err != nil {
      os.Exit(1)
    }
    // FIXME: should not use such a stupid command-line
    switch strings.TrimSpace(source) {
      case "": {
        continue
      }
      case "c": fallthrough
      case "cl": fallthrough
      case "cle": fallthrough
      case "clea": fallthrough
      case "clear": fallthrough
      case "cls": fallthrough
      case "re": fallthrough
      case "res": fallthrough
      case "rese": fallthrough
      case "reset": {
        sources = []string { }
      }
      case "e": fallthrough
      case "ev": fallthrough
      case "eva": fallthrough
      case "eval": fallthrough
      case "r": fallthrough
      case "ru": fallthrough
      case "run": {
        dst, err := compiler.CompileString(strings.Join(sources, ""))
        if err != nil {
          fmt.Fprintln(os.Stderr, err)
          continue
        }
        _, err = os.Stat(dst)
        if os.IsNotExist(err) {
          fmt.Fprintln(os.Stderr, err)
          continue
        }
        abspath, err := filepath.Abs(dst)
        if err != nil {
          fmt.Fprintln(os.Stderr, err)
          continue
        }
        cmd := exec.Command(abspath)
        out, err := cmd.CombinedOutput()
        if err != nil {
          fmt.Fprint(os.Stderr, string(out))
          fmt.Fprint(os.Stderr, err)
          continue
        }
        fmt.Fprint(os.Stdout, string(out))
      }
      case "ex": fallthrough
      case "exi": fallthrough
      case "exit": {
        os.Exit(0)
      }
      case "l": fallthrough
      case "li": fallthrough
      case "lis": fallthrough
      case "list": fallthrough
      case "ls": fallthrough
      case "s": fallthrough
      case "sh": fallthrough
      case "sho": fallthrough
      case "show": {
        fmt.Print(strings.Join(sources, ""))
      }
      default: {
        sources = append(sources, source)
      }
    }
  }
}
