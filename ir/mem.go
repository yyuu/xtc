package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Mem struct {
  ClassName string
  TypeId int
  Expr core.IExpr
}

func NewMem(t int, expr core.IExpr) *Mem {
  return &Mem { "ir.Mem", t, expr }
}

func (self Mem) AsExpr() core.IExpr {
  return self
}

func (self Mem) GetTypeId() int {
  return self.TypeId
}

func (self Mem) IsAddr() bool {
  return false
}

func (self Mem) IsConstant() bool {
  return false
}

func (self Mem) IsVar() bool {
  return false
}

func (self Mem) GetAddressNode(t int) core.IExpr {
  return self.Expr
}