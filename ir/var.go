package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Var struct {
  ClassName string
  Type core.IType
  Entity core.IEntity
}

func NewVar(t core.IType, e core.IEntity) *Var {
  return &Var { "ir.Var", t, e }
}

func (self Var) AsExpr() core.IExpr {
  return self
}

func (self Var) GetType() core.IType {
  return self.Type
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

func (self Var) GetAddressNode(t core.IType) core.IExpr {
  return NewAddr(t, self.Entity)
}
