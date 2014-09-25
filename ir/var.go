package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Var struct {
  ClassName string
  TypeId int
  Entity core.IEntity
}

func NewVar(t int, e core.IEntity) *Var {
  return &Var { "ir.Var", t, e }
}

func (self Var) AsExpr() core.IExpr {
  return self
}

func (self Var) GetTypeId() int {
  return self.TypeId
}

func (self Var) IsAddr() bool {
  return false
}

func (self Var) IsConstant() bool {
  return false
}

func (self Var) IsVar() bool {
  return true
}

func (self Var) GetAddressNode(t int) core.IExpr {
  return NewAddr(t, self.Entity)
}
