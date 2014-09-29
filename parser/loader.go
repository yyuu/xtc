package parser

import (
  "os"
  "strings"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
)

func loadLibrary(name string, errorHandler *bs_core.ErrorHandler, options *bs_core.Options) *bs_ast.Declaration {
  pwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  LIBRARY_PATH := pwd + "/bslib" // FIXME: should be configurable
  path := LIBRARY_PATH + "/" + strings.Replace(name, ".", "/", -1) + ".txt"
  ast, err := ParseFile(path, errorHandler, options)
  if err != nil {
    panic(err)
  }
  return ast.GetDeclaration()
}
