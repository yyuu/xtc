package entity

import (
  "encoding/json"
  "fmt"
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

func (self DefinedVariable) String() string {
  return fmt.Sprintf("<entity.DefinedVariable Name=%s Private=%v TypeNode=%s NumRefered=%d Initializer=%s>", self.Name, self.Private, self.TypeNode, self.NumRefered, self.Initializer)
}

func (self DefinedVariable) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Private bool
    Name string
    TypeNode duck.ITypeNode
    NumRefered int
    Initializer duck.IExprNode
  }
  x.ClassName = "entity.DefinedVariable"
  x.Private = self.Private
  x.Name = self.Name
  x.TypeNode = self.TypeNode
  x.NumRefered = self.NumRefered
  x.Initializer = self.Initializer
  return json.Marshal(x)
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

func (self UndefinedVariable) String() string {
  return fmt.Sprintf("<entity.UndefinedVariable Name=%s Private=%v TypeNode=%s>", self.Name, self.Private, self.TypeNode)
}

func (self UndefinedVariable) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Private bool
    Name string
    TypeNode duck.ITypeNode
  }
  x.ClassName = "entity.UndefinedVariable"
  x.Private = self.Private
  x.Name = self.Name
  x.TypeNode = self.TypeNode
  return json.Marshal(x)
}

func (self UndefinedVariable) IsEntity() bool {
  return true
}
