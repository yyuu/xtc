package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

type Constant struct {
  ClassName string
  Name string
  TypeNode core.ITypeNode
  Value core.IExprNode
  numRefered int
  memref core.IMemoryReference
  address core.IOperand
}

func NewConstant(t core.ITypeNode, name string, value core.IExprNode) *Constant {
  return &Constant {
    ClassName: "entity.Constant",
    Name: name,
    TypeNode: t,
    Value: value,
    numRefered: 0,
    memref: nil,
    address: nil,
  }
}

func NewConstants(xs...*Constant) []*Constant {
  if 0 < len(xs) {
    return xs
  } else {
    return []*Constant { }
  }
}

func AsConstant(x core.IEntity) *Constant {
  return x.(*Constant)
}

func (self *Constant) String() string {
  return fmt.Sprintf("<entity.Constant Name=%s TypeNode=%s Value=%s>", self.Name, self.TypeNode, self.Value)
}

func (self *Constant) IsConstant() bool {
  return true
}

func (self *Constant) IsDefined() bool {
  return true
}

func (self *Constant) IsPrivate() bool {
  return false
}

func (self *Constant) IsParameter() bool {
  return false
}

func (self *Constant) IsVariable() bool {
  return false
}

func (self *Constant) GetNumRefered() int {
  return self.numRefered
}

func (self *Constant) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *Constant) Refered() {
  self.numRefered++
}

func (self *Constant) GetName() string {
  return self.Name
}

func (self *Constant) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *Constant) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *Constant) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *Constant) GetValue() core.IExprNode {
  return self.Value
}

func (self *Constant) SetValue(val core.IExprNode) {
  self.Value = val
}

func (self *Constant) SymbolString() string {
  return self.Name
}

func (self *Constant) GetMemref() core.IMemoryReference {
  checkAddress(self, self.memref, self.address)
  return self.memref
}

func (self *Constant) SetMemref(memref core.IMemoryReference) {
  self.memref = memref
}

func (self *Constant) GetAddress() core.IOperand {
  checkAddress(self, self.memref, self.address)
  return self.address
}

func (self *Constant) SetAddress(address core.IOperand) {
  self.address = address
}
