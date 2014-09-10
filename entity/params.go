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

func (self Params) IsEntity() bool {
  return true
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

func (self Params) IsRefered() bool {
  return true // FIXME: count up references
}

func (self Params) GetLocation() core.Location {
  return self.Location
}

func (self Params) GetParamDescs() []*Parameter {
  return self.ParamDescs
}

func (self Params) GetName() string {
  panic("Parameter#GetName called")
}
