package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

type DereferenceChecker struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
}

func NewDereferenceChecker(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *DereferenceChecker {
  return &DereferenceChecker { errorHandler, table }
}

func (self *DereferenceChecker) Check(a *ast.AST) {
  self.errorHandler.Warnln("DereferenceChecker#Check is not implemented yet")
}
