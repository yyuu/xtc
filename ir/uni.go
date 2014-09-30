package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Uni struct {
  ClassName string
  TypeId int
  Op int
  Expr core.IExpr
}

func NewUni(t int, op int, expr core.IExpr) *Uni {
  return &Uni { "ir.Uni", t, op, expr }
}

func (self *Uni) AsExpr() core.IExpr {
  return self
}

func (self Uni) GetTypeId() int {
  return self.TypeId
}

func (self Uni) IsAddr() bool {
  return false
}

func (self Uni) IsConstant() bool {
  return false
}

func (self Uni) IsVar() bool {
  return false
}

func (self *Uni) GetAddress() core.IOperand {
  panic("#GetAddress called")
}

func (self *Uni) GetAsmValue() core.IImmediateValue {
  panic("#GetAsmValue called")
}

func (self *Uni) GetMemref() core.IMemoryReference {
  panic("#GetMemref called")
}

func (self Uni) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}
