package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/ir"
  "bitbucket.org/yyuu/bs/typesys"
)

type IRGenerator struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
}

func NewIRGenerator(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *IRGenerator {
  return &IRGenerator { errorHandler, table }
}

func (self *IRGenerator) Generate(a *ast.AST) *ir.IR {
  return a.GenerateIR()
}
