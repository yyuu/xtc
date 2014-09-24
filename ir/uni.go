package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Uni struct {
  ClassName string
  Type core.IType
  Op int
  Expr core.IExpr
}

func NewUni(t core.IType, op int, expr core.IExpr) *Uni {
  return &Uni { "ir.Uni", t, op, expr }
}

func (self Uni) AsExpr() core.IExpr {
  return self
}

func (self Uni) GetType() core.IType {
  return self.Type
}

func (self Uni) IsAddr() bool {
  return false
}

func (self Uni) IsConstant() bool {
  return false
}

func (self Uni) IsVar() bool {
  return false
}

func (self Uni) GetAddressNode(t core.IType) core.IExpr {
  panic("unexpected node for LHS")
}
