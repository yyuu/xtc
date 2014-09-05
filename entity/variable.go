package entity

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type DefinedVariable struct {
  private bool
  name string
  typeNode duck.ITypeNode
  numRefered int
  initializer duck.IExprNode
}

func NewDefinedVariable(isPrivate bool, t duck.ITypeNode, name string, init duck.IExprNode) DefinedVariable {
  return DefinedVariable { isPrivate, name, t, 0, init }
}

func (self DefinedVariable) String() string {
  return fmt.Sprintf("<entity.DefinedVariable Name=%s Private=%v TypeNode=%s NumRefered=%d Initializer=%s>", self.name, self.private, self.typeNode, self.numRefered, self.initializer)
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
  x.Private = self.private
  x.Name = self.name
  x.TypeNode = self.typeNode
  x.NumRefered = self.numRefered
  x.Initializer = self.initializer
  return json.Marshal(x)
}

func (self DefinedVariable) IsEntity() bool {
  return true
}

func (self DefinedVariable) IsVariable() bool {
  return true
}

func (self DefinedVariable) IsDefinedVariable() bool {
  return true
}

func (self DefinedVariable) HasInitializer() bool {
  return self.initializer != nil
}

func (self DefinedVariable) IsPrivate() bool {
  return self.private
}

func (self DefinedVariable) GetName() string {
  return self.name
}

func (self DefinedVariable) GetTypeNode() duck.ITypeNode {
  return self.typeNode
}

func (self DefinedVariable) NumRefered() int {
  return self.numRefered
}

func (self DefinedVariable) GetInitializer() duck.IExprNode {
  return self.initializer
}

type UndefinedVariable struct {
  private bool
  name string
  typeNode duck.ITypeNode
}

func NewUndefinedVariable(t duck.ITypeNode, name string) UndefinedVariable {
  return UndefinedVariable { false, name, t }
}

func (self UndefinedVariable) String() string {
  return fmt.Sprintf("<entity.UndefinedVariable Name=%s Private=%v TypeNode=%s>", self.name, self.private, self.typeNode)
}

func (self UndefinedVariable) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Private bool
    Name string
    TypeNode duck.ITypeNode
  }
  x.ClassName = "entity.UndefinedVariable"
  x.Private = self.private
  x.Name = self.name
  x.TypeNode = self.typeNode
  return json.Marshal(x)
}

func (self UndefinedVariable) IsEntity() bool {
  return true
}

func (self UndefinedVariable) IsVariable() bool {
  return true
}

func (self UndefinedVariable) IsUndefinedVariable() bool {
  return true
}
