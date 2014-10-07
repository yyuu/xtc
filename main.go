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
    line, err := stdin.ReadString('\n')
    if err != nil {
      os.Exit(1)
    }
    trimedLine := strings.TrimSpace(line)
    // FIXME: should not use such a stupid command-line
    if 0 < len(trimedLine) && strings.Index(trimedLine, ":") == 0 {
      switch trimedLine[1:] {
        case "clear": fallthrough
        case "cls": fallthrough
        case "reset": {
          sources = []string { }
        }
        case "eval": fallthrough
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
            fmt.Fprintln(os.Stderr, err)
            continue
          }
          fmt.Fprint(os.Stdout, string(out))
        }
        case "exit": fallthrough
        case "quit": {
          os.Exit(0)
        }
        case "list": fallthrough
        case "ls": fallthrough
        case "show": {
          fmt.Print(strings.Join(sources, ""))
        }
        case "del": fallthrough
        case "delete": {
          var lineno int
          fmt.Print("line number: ")
          fmt.Scanf("%d", &lineno)
          if lineno <= len(sources) {
            if lineno == 1 {
              sources = sources[1:]
            } else {
              sources = append(sources[0:lineno-1], sources[lineno:]...)
            }
          } else {
            fmt.Fprintln(os.Stderr, "out of range")
          }
        }
        default: {
          fmt.Fprintln(os.Stderr, "available commands:")
          fmt.Fprintln(os.Stderr, "clean, clear, cls, reset, eval, run, exit, quit, list, ls, show")
        }
      }
    } else {
      sources = append(sources, line)
    }
  }
}
