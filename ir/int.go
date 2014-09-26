package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Int struct {
  ClassName string
  TypeId int
  Value int64
}

func NewInt(t int, value int64) *Int {
  return &Int { "ir.Int", t, value }
}

func (self Int) AsExpr() core.IExpr {
  return self
}

func (self Int) GetTypeId() int {
  return self.TypeId
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

func (self Int) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}

func (self Int) GetValue() int64 {
  return self.Value
}
