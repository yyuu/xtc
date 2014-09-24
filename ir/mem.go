package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Mem struct {
  ClassName string
  Type core.IType
  Expr core.IExpr
}

func NewMem(t core.IType, expr core.IExpr) *Mem {
  return &Mem { "ir.Mem", t, expr }
}

func (self Mem) AsExpr() core.IExpr {
  return self
}

func (self Mem) GetType() core.IType {
  return self.Type
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

func (self Mem) GetAddressNode(t core.IType) core.IExpr {
  return self.Expr
}
