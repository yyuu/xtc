package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Addr struct {
  ClassName string
  TypeId int
  Entity core.IEntity
}

func NewAddr(t int, e core.IEntity) *Addr {
  return &Addr { "ir.Addr", t, e }
}

func (self *Addr) AsExpr() core.IExpr {
  return self
}

func (self Addr) GetTypeId() int {
  return self.TypeId
}

func (self Addr) GetEntity() core.IEntity {
  return self.Entity
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

func (self *Addr) GetAddress() core.IOperand {
  return self.Entity.GetAddress()
}

func (self *Addr) GetAsmValue() core.IImmediateValue {
  panic("#GetAsmValue called")
}

func (self *Addr) GetMemref() core.IMemoryReference {
  return self.Entity.GetMemref()
}

func (self Addr) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}

func (self Addr) GetEntityForce() core.IEntity {
  return self.Entity
}
