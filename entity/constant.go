package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type Constant struct {
  ClassName string
  Name string
  TypeNode core.ITypeNode
  Value core.IExprNode
  numRefered int
}

func NewConstant(t core.ITypeNode, name string, value core.IExprNode) *Constant {
  return &Constant { "entity.Constant", name, t, value, 0 }
}

func NewConstants(xs...*Constant) []*Constant {
  if 0 < len(xs) {
    return xs
  } else {
    return []*Constant { }
  }
}

func (self Constant) String() string {
  return fmt.Sprintf("<entity.Constant Name=%s TypeNode=%s Value=%s>", self.Name, self.TypeNode, self.Value)
}

func (self Constant) IsConstant() bool {
  return true
}

func (self Constant) IsDefined() bool {
  return true
}

func (self Constant) IsPrivate() bool {
  return false
}

func (self Constant) IsParameter() bool {
  return false
}

func (self Constant) GetNumRefered() int {
  return self.numRefered
}

func (self Constant) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *Constant) Refered() {
  self.numRefered++
}

func (self Constant) GetName() string {
  return self.Name
}

func (self Constant) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self Constant) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self Constant) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self Constant) GetValue() core.IExprNode {
  return self.Value
}

func (self *Constant) SetValue(val core.IExprNode) {
  self.Value = val
}
