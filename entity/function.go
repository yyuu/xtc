package entity

import (
  "encoding/json"
  "fmt"
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

func (self DefinedFunction) String() string {
  return fmt.Sprintf("<entity.DefinedFunction Name=%s Private=%v TypeNode=%s Params=%s Body=%s>", self.Name, self.Private, self.TypeNode, self.Params, self.Body)
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
  x.Private = self.Private
  x.TypeNode = self.TypeNode
  x.Name = self.Name
  x.Params = self.Params
  x.Body = self.Body
  return json.Marshal(x)
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

func (self UndefinedFunction) String() string {
  return fmt.Sprintf("<entity.UndefinedFunction Name=%s TypeNode=%s Params=%s>", self.Name, self.TypeNode, self.Params)
}

func (self UndefinedFunction) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    TypeNode duck.ITypeNode
    Name string
    Params Params
  }
  x.ClassName = "entity.UndefinedFunction"
  x.TypeNode = self.TypeNode
  x.Name = self.Name
  x.Params = self.Params
  return json.Marshal(x)
}

func (self UndefinedFunction) IsEntity() bool {
  return true
}
