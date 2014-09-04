package entity

import (
  "bitbucket.org/yyuu/bs/duck"
)

type Constant struct {
  Name string
  TypeNode duck.ITypeNode
  Value duck.IExprNode
}

func NewConstant(t duck.ITypeNode, name string, value duck.IExprNode) Constant {
  return Constant {
    TypeNode: t,
    Name: name,
    Value: value,
  }
}

func (self Constant) IsEntity() bool {
  return true
}
