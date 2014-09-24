package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Bin struct {
  ClassName string
  Type core.IType
  Op int
  Left core.IExpr
  Right core.IExpr
}

func NewBin(t core.IType, op int, left core.IExpr, right core.IExpr) *Bin {
  return &Bin { "ir.Bin", t, op, left, right }
}

func (self Bin) AsExpr() core.IExpr {
  return self
}

func (self Bin) GetType() core.IType {
  return self.Type
}

func (self Bin) IsAddr() bool {
  return false
}

func (self Bin) IsConstant() bool {
  return false
}

func (self Bin) IsVar() bool {
  return false
}

func (self Bin) GetAddressNode(t core.IType) core.IExpr {
  panic("unexpected node for LHS")
}
