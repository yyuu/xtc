package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Assign struct {
  ClassName string
  Location core.Location
  LHS core.IExpr
  RHS core.IExpr
}

func NewAssign(loc core.Location, lhs core.IExpr, rhs core.IExpr) *Assign {
  return &Assign { "ir.Assign", loc, lhs, rhs }
}

func (self Assign) AsStmt() core.IStmt {
  return self
}

func (self Assign) GetLocation() core.Location {
  return self.Location
}
