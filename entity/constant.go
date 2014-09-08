package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Constant struct {
  ClassName string
  Name string
  TypeNode duck.ITypeNode
  Value duck.IExprNode
}

func NewConstant(t duck.ITypeNode, name string, value duck.IExprNode) Constant {
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

func (self Constant) GetName() string {
  return self.Name
}

func (self Constant) GetTypeNode() duck.ITypeNode {
  return self.TypeNode
}

func (self Constant) GetValue() duck.IExprNode {
  return self.Value
}
