package entity

import (
  "bitbucket.org/yyuu/bs/duck"
)

type DefinedFunction struct {
  Private bool
  TypeNode duck.ITypeNode
  Name string
  Params Params
  Body duck.IStmtNode
}

func NewDefinedFunction(priv bool, t duck.ITypeNode, name string, params Params, body duck.IStmtNode) DefinedFunction {
  return DefinedFunction {
    Private: priv,
    TypeNode: t,
    Name: name,
    Params: params,
    Body: body,
  }
}

func (self DefinedFunction) IsEntity() bool {
  return true
}

type UndefinedFunction struct {
  TypeNode duck.ITypeNode
  Name string
  Params Params
}

func NewUndefinedFunction(t duck.ITypeNode, name string, params Params) UndefinedFunction {
  return UndefinedFunction {
    TypeNode: t,
    Name: name,
    Params: params,
  }
}

func (self UndefinedFunction) IsEntity() bool {
  return true
}
