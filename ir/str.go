package ir

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type Str struct {
  ClassName string
  Type core.IType
  Entry *entity.ConstantEntry
}

func NewStr(t core.IType, entry *entity.ConstantEntry) *Str {
  return &Str { "ir.Str", t, entry }
}

func (self Str) AsExpr() core.IExpr {
  return self
}

func (self Str) GetType() core.IType {
  return self.Type
}

func (self Str) IsAddr() bool {
  return false
}

func (self Str) IsConstant() bool {
  return true
}

func (self Str) IsVar() bool {
  return false
}

func (self Str) GetAddressNode(t core.IType) core.IExpr {
  panic("unexpected node for LHS")
}
