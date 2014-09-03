package entity

import (
  "bitbucket.org/yyuu/bs/typesys"
)

type Params struct {
  Location ILocation
  ParamDescs []Parameter
}

func NewParams(loc ILocation, paramDescs []Parameter) Params {
  return Params {
    Location: loc,
    ParamDescs: paramDescs,
  }
}

func (self Params) IsEntity() bool {
  return true
}

func (self Params) ParametersTypeRef() typesys.ParamTypeRefs {
  ps := make([]typesys.ITypeRef, len(self.ParamDescs))
  for i := range self.ParamDescs {
    ps[i] = self.ParamDescs[i].TypeNode.GetTypeRef()
  }
  return typesys.NewParamTypeRefs(self.Location, ps, false)
}

type Parameter struct {
  TypeNode ITypeNode
  Name string
}

func NewParameter(t ITypeNode, name string) Parameter {
  return Parameter {
    TypeNode: t,
    Name: name,
  }
}

func (self Parameter) IsEntity() bool {
  return true
}
