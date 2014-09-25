package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Call struct {
  ClassName string
  TypeId int
  Expr core.IExpr
  Args []core.IExpr
}

func NewCall(t int, expr core.IExpr, args []core.IExpr) *Call {
  return &Call { "ir.Call", t, expr, args }
}

func (self Call) AsExpr() core.IExpr {
  return self
}

func (self Call) GetTypeId() int {
  return self.TypeId
}

func (self Call) IsAddr() bool {
  return false
}

func (self Call) IsConstant() bool {
  return false
}

func (self Call) IsVar() bool {
  return false
}

func (self Call) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}
