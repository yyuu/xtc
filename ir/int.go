package ir

import (
  "bitbucket.org/yyuu/bs/asm"
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

func (self *Int) AsExpr() core.IExpr {
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

func (self *Int) GetAddress() core.IOperand {
  panic("#GetAddress called")
}

func (self *Int) GetAsmValue() core.IImmediateValue {
  return asm.NewImmediateValue(asm.NewIntegerLiteral(self.Value))
}

func (self *Int) GetMemref() core.IMemoryReference {
  panic("#GetMemref called")
}

func (self Int) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}

func (self Int) GetValue() int64 {
  return self.Value
}

func (self Int) GetEntityForce() core.IEntity {
  return nil
}
