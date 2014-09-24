package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Addr struct {
  ClassName string
  Type core.IType
  Entity core.IEntity
}

func NewAddr(t core.IType, e core.IEntity) *Addr {
  return &Addr { "ir.Addr", t, e }
}

func (self Addr) AsExpr() core.IExpr {
  return self
}

func (self Addr) GetType() core.IType {
  return self.Type
}

func (self Addr) IsAddr() bool {
  return true
}

func (self Addr) IsConstant() bool {
  return false
}

func (self Addr) IsVar() bool {
  return false
}

func (self Addr) GetAddressNode(t core.IType) core.IExpr {
  panic("unexpected node for LHS")
}
