package ir

import (
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type Str struct {
  ClassName string
  TypeId int
  Entry *entity.ConstantEntry
}

func NewStr(t int, entry *entity.ConstantEntry) *Str {
  return &Str { "ir.Str", t, entry }
}

func (self Str) AsExpr() core.IExpr {
  return self
}

func (self Str) GetTypeId() int {
  return self.TypeId
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

func (self Str) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}