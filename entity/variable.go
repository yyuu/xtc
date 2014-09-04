package entity

import (
  "bitbucket.org/yyuu/bs/duck"
)

type DefinedVariable struct {
  Private bool
  Name string
  TypeNode duck.ITypeNode
  NumRefered int
  Initializer duck.IExprNode
}

func NewDefinedVariable(isPrivate bool, t duck.ITypeNode, name string, init duck.IExprNode) DefinedVariable {
  return DefinedVariable {
    Private: isPrivate,
    Name: name,
    TypeNode: t,
    NumRefered: 0,
    Initializer: init,
  }
}

func (self DefinedVariable) IsEntity() bool {
  return true
}

func (self DefinedVariable) IsDefined() bool {
  return true
}

func (self DefinedVariable) HasInitializer() bool {
  return self.Initializer != nil
}

type UndefinedVariable struct {
  Private bool
  Name string
  TypeNode duck.ITypeNode
}

func NewUndefinedVariable(t duck.ITypeNode, name string) UndefinedVariable {
  return UndefinedVariable {
    Private: false,
    Name: name,
    TypeNode: t,
  }
}

func (self UndefinedVariable) IsEntity() bool {
  return true
}
