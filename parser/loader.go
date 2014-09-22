package parser

import (
  "os"
  "strings"
  "bitbucket.org/yyuu/bs/ast"
)

func loadLibrary(name string) *ast.Declaration {
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  LIBRARY_PATH := pwd + "/bslib" // FIXME: should be configurable
  path := LIBRARY_PATH + "/" + strings.Replace(name, ".", "/", -1) + ".txt"
  aAST, err := ParseFile(path)
  if err != nil {
    panic(err)
  }
  return aAST.GetDeclaration()
}
