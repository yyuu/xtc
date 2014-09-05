package entity

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/typesys"
)

type Params struct {
  location duck.ILocation
  paramDescs []Parameter
}

func NewParams(loc duck.ILocation, paramDescs []Parameter) Params {
  return Params { loc, paramDescs }
}

func (self Params) String() string {
  return fmt.Sprintf("<entity.Params Location=%s ParamDescs=%s>", self.location, self.paramDescs)
}

func (self Params) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    ParamDescs []Parameter
  }
  x.ClassName = "entity.Params"
  x.Location = self.location
  x.ParamDescs = self.paramDescs
  return json.Marshal(x)
}

func (self Params) IsEntity() bool {
  return true
}

func (self Params) GetParamDescs() []Parameter {
  return self.paramDescs
}

func (self Params) ParametersTypeRef() typesys.ParamTypeRefs {
  ps := make([]duck.ITypeRef, len(self.paramDescs))
  for i := range self.paramDescs {
    ps[i] = self.paramDescs[i].typeNode.GetTypeRef()
  }
  return typesys.NewParamTypeRefs(self.location, ps, false)
}

type Parameter struct {
  typeNode duck.ITypeNode
  name string
}

func NewParameter(t duck.ITypeNode, name string) Parameter {
  return Parameter { t, name }
}

func (self Parameter) String() string {
  return fmt.Sprintf("<entity.Parameter Name=%s TypeNode=%s>", self.name, self.typeNode)
}

func (self Parameter) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    TypeNode duck.ITypeNode
    Name string
  }
  x.ClassName = "entity.Parameter"
  x.TypeNode = self.typeNode
  x.Name = self.name
  return json.Marshal(x)
}

func (self Parameter) IsEntity() bool {
  return true
}
