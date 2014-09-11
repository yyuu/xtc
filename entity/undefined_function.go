package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type UndefinedFunction struct {
  ClassName string
  TypeNode core.ITypeNode
  Name string
  Params *Params
}

func NewUndefinedFunction(t core.ITypeNode, name string, params *Params) *UndefinedFunction {
  return &UndefinedFunction { "entity.UndefinedFunction", t, name, params }
}

func NewUndefinedFunctions(xs...*UndefinedFunction) []*UndefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []*UndefinedFunction { }
  }
}

func (self UndefinedFunction) String() string {
  return fmt.Sprintf("<entity.UndefinedFunction Name=%s TypeNode=%s Params=%s>", self.Name, self.TypeNode, self.Params)
}

func (self UndefinedFunction) IsDefined() bool {
  return false
}

func (self UndefinedFunction) IsConstant() bool {
  return false
}

func (self UndefinedFunction) IsPrivate() bool {
  return true
}

func (self UndefinedFunction) IsRefered() bool {
  return true // FIXME: count up references
}

func (self UndefinedFunction) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self UndefinedFunction) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self UndefinedFunction) GetName() string {
  return self.Name
}

func (self UndefinedFunction) GetParams() *Params {
  return self.Params
}

func (self UndefinedFunction) GetParameters() []*Parameter {
  return self.Params.ParamDescs
}
