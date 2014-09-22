package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

type TypeChecker struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
}

func NewTypeChecker(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *TypeChecker {
  return &TypeChecker { errorHandler, table }
}

func (self *TypeChecker) Check(a *ast.AST) {
  self.errorHandler.Warnln("TypeChecker#Check is not implemented yet")
}
