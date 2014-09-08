package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type DefinedFunction struct {
  ClassName string
  Private bool
  TypeNode duck.ITypeNode
  Name string
  Params Params
  Body duck.IStmtNode
}

func NewDefinedFunction(priv bool, t duck.ITypeNode, name string, params Params, body duck.IStmtNode) DefinedFunction {
  return DefinedFunction { "entity.DefinedFunction", priv, t, name, params, body }
}

func (self DefinedFunction) String() string {
  return fmt.Sprintf("<entity.DefinedFunction Name=%s Private=%v TypeNode=%s Params=%s Body=%s>", self.Name, self.Private, self.TypeNode, self.Params, self.Body)
}

func (self DefinedFunction) IsPrivate() bool {
  return self.Private
}

func (self DefinedFunction) GetTypeNode() duck.ITypeNode {
  return self.TypeNode
}

func (self DefinedFunction) GetName() string {
  return self.Name
}

func (self DefinedFunction) IsEntity() bool {
  return true
}

func (self DefinedFunction) IsFunction() bool {
  return true
}

func (self DefinedFunction) IsDefinedFunction() bool {
  return true
}

func (self DefinedFunction) GetParams() duck.IParams {
  return self.Params
}

func (self DefinedFunction) GetBody() duck.IStmtNode {
  return self.Body
}

type UndefinedFunction struct {
  ClassName string
  TypeNode duck.ITypeNode
  Name string
  Params Params
}

func NewUndefinedFunction(t duck.ITypeNode, name string, params Params) UndefinedFunction {
  return UndefinedFunction { "entity.UndefinedFunction", t, name, params }
}

func (self UndefinedFunction) String() string {
  return fmt.Sprintf("<entity.UndefinedFunction Name=%s TypeNode=%s Params=%s>", self.Name, self.TypeNode, self.Params)
}

func (self UndefinedFunction) IsEntity() bool {
  return true
}

func (self UndefinedFunction) IsFunction() bool {
  return true
}

func (self UndefinedFunction) IsUndefinedFunction() bool {
  return true
}

func (self UndefinedFunction) GetTypeNode() duck.ITypeNode {
  return self.TypeNode
}

func (self UndefinedFunction) GetName() string {
  return self.Name
}

func (self UndefinedFunction) GetParams() Params {
  return self.Params
}
