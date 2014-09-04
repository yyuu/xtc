package entity

import (
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/typesys"
)

type Params struct {
  Location duck.ILocation
  ParamDescs []Parameter
}

func NewParams(loc duck.ILocation, paramDescs []Parameter) Params {
  return Params {
    Location: loc,
    ParamDescs: paramDescs,
  }
}

func (self Params) IsEntity() bool {
  return true
}

func (self Params) ParametersTypeRef() typesys.ParamTypeRefs {
  ps := make([]duck.ITypeRef, len(self.ParamDescs))
  for i := range self.ParamDescs {
    ps[i] = self.ParamDescs[i].TypeNode.GetTypeRef()
  }
  return typesys.NewParamTypeRefs(self.Location, ps, false)
}

type Parameter struct {
  TypeNode duck.ITypeNode
  Name string
}

func NewParameter(t duck.ITypeNode, name string) Parameter {
  return Parameter {
    TypeNode: t,
    Name: name,
  }
}

func (self Parameter) IsEntity() bool {
  return true
}
