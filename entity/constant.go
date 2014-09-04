package entity

import (
  "bitbucket.org/yyuu/bs/duck"
)

type Constant struct {
  Name string
  TypeNode duck.ITypeNode
  Value duck.IExprNode
}

func NewConstant(name string, t duck.ITypeNode, value duck.IExprNode) Constant {
  return Constant {
    Name: name,
    TypeNode: t,
    Value: value,
  }
}

func (self Constant) IsEntity() bool {
  return true
}
