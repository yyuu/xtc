package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Int struct {
  ClassName string
  Type core.IType
  Value int64
}

func NewInt(t core.IType, value int64) *Int {
  return &Int { "ir.Int", t, value }
}

func (self Int) AsExpr() core.IExpr {
  return self
}

func (self Int) GetType() core.IType {
  return self.Type
}

func (self Int) IsAddr() bool {
  return false
}

func (self Int) IsConstant() bool {
  return true
}

func (self Int) IsVar() bool {
  return false
}

func (self Int) GetAddressNode(t core.IType) core.IExpr {
  panic("unexpected node for LHS")
}
