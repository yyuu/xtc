package entity

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Params struct {
  location duck.ILocation
  paramDescs []duck.IParameter
}

func NewParams(loc duck.ILocation, paramDescs []duck.IParameter) Params {
  return Params { loc, paramDescs }
}

func (self Params) String() string {
  return fmt.Sprintf("<entity.Params Location=%s ParamDescs=%s>", self.location, self.paramDescs)
}

func (self Params) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    ParamDescs []duck.IParameter
  }
  x.ClassName = "entity.Params"
  x.Location = self.location
  x.ParamDescs = self.paramDescs
  return json.Marshal(x)
}

func (self Params) IsEntity() bool {
  return true
}

func (self Params) GetLocation() duck.ILocation {
  return self.location
}

func (self Params) GetParamDescs() []duck.IParameter {
  return self.paramDescs
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

func (self Parameter) GetTypeNode() duck.ITypeNode {
  return self.typeNode
}

func (self Parameter) GetName() string {
  return self.name
}
