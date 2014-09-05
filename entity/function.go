package entity

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type DefinedFunction struct {
  private bool
  typeNode duck.ITypeNode
  name string
  params Params
  body duck.IStmtNode
}

func NewDefinedFunction(priv bool, t duck.ITypeNode, name string, params Params, body duck.IStmtNode) DefinedFunction {
  return DefinedFunction { priv, t, name, params, body }
}

func (self DefinedFunction) String() string {
  return fmt.Sprintf("<entity.DefinedFunction Name=%s Private=%v TypeNode=%s Params=%s Body=%s>", self.name, self.private, self.typeNode, self.params, self.body)
}

func (self DefinedFunction) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Private bool
    TypeNode duck.ITypeNode
    Name string
    Params Params
    Body duck.IStmtNode
  }
  x.ClassName = "entity.DefinedFunction"
  x.Private = self.private
  x.TypeNode = self.typeNode
  x.Name = self.name
  x.Params = self.params
  x.Body = self.body
  return json.Marshal(x)
}

func (self DefinedFunction) IsEntity() bool {
  return true
}

type UndefinedFunction struct {
  typeNode duck.ITypeNode
  name string
  params Params
}

func NewUndefinedFunction(t duck.ITypeNode, name string, params Params) UndefinedFunction {
  return UndefinedFunction { t, name, params }
}

func (self UndefinedFunction) String() string {
  return fmt.Sprintf("<entity.UndefinedFunction Name=%s TypeNode=%s Params=%s>", self.name, self.typeNode, self.params)
}

func (self UndefinedFunction) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    TypeNode duck.ITypeNode
    Name string
    Params Params
  }
  x.ClassName = "entity.UndefinedFunction"
  x.TypeNode = self.typeNode
  x.Name = self.name
  x.Params = self.params
  return json.Marshal(x)
}

func (self UndefinedFunction) IsEntity() bool {
  return true
}
