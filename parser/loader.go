package parser

import (
  "fmt"
  "os"
  "strings"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
)

type libraryLoader struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
}

func newLibraryLoader(errorHandler *bs_core.ErrorHandler, options *bs_core.Options) *libraryLoader {
  return &libraryLoader { errorHandler, options }
}

func (self *libraryLoader) loadLibrary(name string) *bs_ast.Declaration {
  path, ok := self.searchLibrary(name)
  if ! ok {
    panic(fmt.Errorf("No such file or directory: %s.bs", name))
  }
  src := bs_core.NewSourceFile(path, path, bs_core.EXT_PROGRAM_SOURCE)
  ast, err := Parse(src, self.errorHandler, self.options)
  if err != nil {
    panic(err)
  }
  return ast.GetDeclaration()
}

func (self *libraryLoader) searchLibrary(name string) (string, bool) {
  file := strings.Replace(name, ".", "/", -1)
  libraryPath := self.options.GetLibraryPath()
  for i := range libraryPath {
    path := fmt.Sprintf("%s/%s%s", libraryPath[i], file, bs_core.EXT_PROGRAM_SOURCE)
    _, err := os.Stat(path)
    if ! os.IsNotExist(err) {
      return path, true
    }
  }
  return "", false
}
