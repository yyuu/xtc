package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Call struct {
  ClassName string
  Type core.IType
  Expr core.IExpr
  Args []core.IExpr
}

func NewCall(t core.IType, expr core.IExpr, args []core.IExpr) *Call {
  return &Call { "ir.Call", t, expr, args }
}

func (self Call) AsExpr() core.IExpr {
  return self
}

func (self Call) GetType() core.IType {
  return self.Type
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

func (self Call) GetAddressNode(t core.IType) core.IExpr {
  panic("unexpected node for LHS")
}
