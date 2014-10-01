package ir

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type Bin struct {
  ClassName string
  TypeId int
  Op int
  Left core.IExpr
  Right core.IExpr
}

func NewBin(t int, op int, left core.IExpr, right core.IExpr) *Bin {
  return &Bin { "ir.Bin", t, op, left, right }
}

func (self *Bin) AsExpr() core.IExpr {
  return self
}

func (self Bin) GetTypeId() int {
  return self.TypeId
}

func (self Bin) GetOp() int {
  return self.Op
}

func (self Bin) GetLeft() core.IExpr {
  return self.Left
}

func (self Bin) GetRight() core.IExpr {
  return self.Right
}

func (self Bin) IsAddr() bool {
  return false
}

func (self Bin) IsConstant() bool {
  return false
}

func (self Bin) IsVar() bool {
  return false
}

func (self *Bin) GetAddress() core.IOperand {
  panic("#GetAddress called")
}

func (self *Bin) GetAsmValue() core.IImmediateValue {
  panic("#GetAsmValue called")
}

func (self *Bin) GetMemref() core.IMemoryReference {
  panic("#GetMemref called")
}

func (self Bin) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}

func (self Bin) GetEntityForce() core.IEntity {
  return nil
}

func (self Bin) String() string {
  return fmt.Sprintf("Bin(%d,%s,%s)", self.Op, self.Left, self.Right)
}
