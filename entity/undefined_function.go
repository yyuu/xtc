package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

type UndefinedFunction struct {
  ClassName string
  TypeNode core.ITypeNode
  Name string
  Params *Params
  numRefered int
}

func NewUndefinedFunction(t core.ITypeNode, name string, params *Params) *UndefinedFunction {
  return &UndefinedFunction { "entity.UndefinedFunction", t, name, params, 0 }
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

func (self UndefinedFunction) IsParameter() bool {
  return false
}

func (self UndefinedFunction) GetNumRefered() int {
  return self.numRefered
}

func (self UndefinedFunction) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *UndefinedFunction) Refered() {
  self.numRefered++
}

func (self UndefinedFunction) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self UndefinedFunction) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self UndefinedFunction) GetType() core.IType {
  return self.TypeNode.GetType()
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

func (self UndefinedFunction) GetReturnType() core.IType {
  t := self.GetType().(*typesys.FunctionType)
  return t.GetReturnType()
}
