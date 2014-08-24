package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "bitbucket.org/yyuu/bs/lexer"
)

var flagSet = flag.NewFlagSet(os.Args[0], 1)
var verbose = flagSet.Bool("v", false, "verbose mode")

func main() {
  flagSet.Parse(os.Args[1:])

  files := flagSet.Args()
  for i := range files {
    cs, err := ioutil.ReadFile(files[i])
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
    }
    lexer := lexer.NewLexer(files[i], string(cs))
    for {
      t := lexer.GetToken()
      if t == nil {
        break
      }
      fmt.Println("Token:", t.ToString())
    }
  }
}
