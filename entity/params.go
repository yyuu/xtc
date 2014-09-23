package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type Params struct {
  ClassName string
  Location core.Location
  ParamDescs []*Parameter
}

func NewParams(loc core.Location, paramDescs []*Parameter) *Params {
  return &Params { "entity.Params", loc, paramDescs }
}

func (self Params) String() string {
  return fmt.Sprintf("<entity.Params Location=%s ParamDescs=%s>", self.Location, self.ParamDescs)
}

func (self Params) IsDefined() bool {
  return false
}

func (self Params) IsConstant() bool {
  return false
}

func (self Params) IsPrivate() bool {
  return true
}

func (self Params) IsParameter() bool {
  return false
}

func (self Params) GetNumRefered() int {
  return -1
}

func (self Params) IsRefered() bool {
  return false
}

func (self *Params) Refered() {
  // nop
}

func (self Params) GetLocation() core.Location {
  return self.Location
}

func (self Params) GetParamDescs() []*Parameter {
  return self.ParamDescs
}

func (self Params) GetName() string {
  panic("Params#GetName called")
}

func (self Params) GetTypeNode() core.ITypeNode {
  panic("Prams#GetTypeNode called")
}

func (self Params) GetTypeRef() core.ITypeRef {
  panic("Prams#GetTypeRef called")
}
