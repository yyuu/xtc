package ir

import (
  "bitbucket.org/yyuu/xtc/core"
)

type ExprStmt struct {
  ClassName string
  Location core.Location
  Expr core.IExpr
}

func NewExprStmt(loc core.Location, expr core.IExpr) *ExprStmt {
  return &ExprStmt { "ir.ExprStmt", loc, expr }
}

func (self *ExprStmt) AsStmt() core.IStmt {
  return self
}

func (self ExprStmt) GetLocation() core.Location {
  return self.Location
}

func (self ExprStmt) GetExpr() core.IExpr {
  return self.Expr
}
