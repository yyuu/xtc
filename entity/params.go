package entity

import (
  "encoding/json"
  "fmt"
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

func (self Params) String() string {
  return fmt.Sprintf("<entity.Params Location=%s ParamDescs=%s>", self.Location, self.ParamDescs)
}

func (self Params) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    ParamDescs []Parameter
  }
  x.ClassName = "entity.Params"
  x.Location = self.Location
  x.ParamDescs = self.ParamDescs
  return json.Marshal(x)
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

func (self Parameter) String() string {
  return fmt.Sprintf("<entity.Parameter Name=%s TypeNode=%s>", self.Name, self.TypeNode)
}

func (self Parameter) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    TypeNode duck.ITypeNode
    Name string
  }
  x.ClassName = "entity.Parameter"
  x.TypeNode = self.TypeNode
  x.Name = self.Name
  return json.Marshal(x)
}

func (self Parameter) IsEntity() bool {
  return true
}
