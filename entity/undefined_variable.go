package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

type UndefinedVariable struct {
  ClassName string
  Private bool
  Name string
  TypeNode core.ITypeNode
  numRefered int
  memref core.IMemoryReference
  address core.IOperand
}

func NewUndefinedVariable(t core.ITypeNode, name string) *UndefinedVariable {
  return &UndefinedVariable {
    ClassName: "entity.UndefinedVariable",
    Private: false,
    Name: name,
    TypeNode: t,
    numRefered: 0,
    memref: nil,
    address: nil,
  }
}

func NewUndefinedVariables(xs...*UndefinedVariable) []*UndefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []*UndefinedVariable { }
  }
}

func AsUndefinedVariable(x core.IEntity) *UndefinedVariable {
  return x.(*UndefinedVariable)
}

func (self *UndefinedVariable) String() string {
  return fmt.Sprintf("extern %s %s; // ref=%d", self.TypeNode.GetTypeRef(), self.Name, self.numRefered)
}

func (self *UndefinedVariable) IsDefined() bool {
  return false
}

func (self *UndefinedVariable) IsConstant() bool {
  return false
}

func (self *UndefinedVariable) IsPrivate() bool {
  return true
}

func (self *UndefinedVariable) IsParameter() bool {
  return false
}

func (self *UndefinedVariable) IsVariable() bool {
  return true
}

func (self *UndefinedVariable) GetNumRefered() int {
  return self.numRefered
}

func (self *UndefinedVariable) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *UndefinedVariable) Refered() {
  self.numRefered++
}

func (self *UndefinedVariable) GetName() string {
  return self.Name
}

func (self *UndefinedVariable) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *UndefinedVariable) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *UndefinedVariable) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *UndefinedVariable) GetValue() core.IExprNode {
  panic("UndefinedVariable#GetValue called")
}

func (self *UndefinedVariable) SymbolString() string {
  return self.Name
}

func (self *UndefinedVariable) GetMemref() core.IMemoryReference {
  checkAddress(self, self.memref, self.address)
  return self.memref
}

func (self *UndefinedVariable) SetMemref(memref core.IMemoryReference) {
  self.memref = memref
}

func (self *UndefinedVariable) GetAddress() core.IOperand {
  checkAddress(self, self.memref, self.address)
  return self.address
}

func (self *UndefinedVariable) SetAddress(address core.IOperand) {
  self.address = address
}
