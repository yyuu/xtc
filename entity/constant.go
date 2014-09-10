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
}

func NewConstant(t core.ITypeNode, name string, value core.IExprNode) Constant {
  return Constant { "entity.Constant", name, t, value }
}

func (self Constant) String() string {
  return fmt.Sprintf("<entity.Constant Name=%s TypeNode=%s Value=%s>", self.Name, self.TypeNode, self.Value)
}

func (self Constant) IsEntity() bool {
  return true
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

func (self Constant) IsRefered() bool {
  return true // FIXME: count references
}

func (self Constant) GetName() string {
  return self.Name
}

func (self Constant) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self Constant) GetValue() core.IExprNode {
  return self.Value
}

func (self Constant) SetValue(val core.IExprNode) core.IConstant {
  self.Value = val
  return self
}
