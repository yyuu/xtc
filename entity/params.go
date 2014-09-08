package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Params struct {
  ClassName string
  Location duck.ILocation
  ParamDescs []duck.IParameter
}

func NewParams(loc duck.ILocation, paramDescs []duck.IParameter) Params {
  return Params { "entity.Params", loc, paramDescs }
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

func (self Params) GetLocation() duck.ILocation {
  return self.Location
}

func (self Params) GetParamDescs() []duck.IParameter {
  return self.ParamDescs
}

func (self Params) GetName() string {
  panic("Parameter#GetName called")
}

type Parameter struct {
  ClassName string
  TypeNode duck.ITypeNode
  Name string
}

func NewParameter(t duck.ITypeNode, name string) Parameter {
  return Parameter { "entity.Parameter", t, name }
}

func (self Parameter) String() string {
  return fmt.Sprintf("<entity.Parameter Name=%s TypeNode=%s>", self.Name, self.TypeNode)
}

func (self Parameter) IsEntity() bool {
  return true
}

func (self Parameter) IsDefined() bool {
  return true
}

func (self Parameter) GetTypeNode() duck.ITypeNode {
  return self.TypeNode
}

func (self Parameter) GetName() string {
  return self.Name
}
