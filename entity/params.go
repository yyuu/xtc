package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Params struct {
  ClassName string
  Location duck.Location
  ParamDescs []duck.IParameter
}

func NewParams(loc duck.Location, paramDescs []duck.IParameter) Params {
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

func (self Params) IsConstant() bool {
  return false
}

func (self Params) IsPrivate() bool {
  return true
}

func (self Params) IsRefered() bool {
  return true // FIXME: count up references
}

func (self Params) GetLocation() duck.Location {
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

func (self Parameter) IsPrivate() bool {
  return false
}

func (self Parameter) IsVariable() bool {
  return true
}

func (self Parameter) IsDefinedVariable() bool {
  return true
}

func (self Parameter) IsConstant() bool {
  return false
}

func (self Parameter) GetInitializer() duck.IExprNode {
  return nil
}

func (self Parameter) SetInitializer(e duck.IExprNode) duck.IDefinedVariable {
  return self
}

func (self Parameter) HasInitializer() bool {
  return false
}

func (self Parameter) GetNumRefered() int {
  return 0
}

func (self Parameter) IsRefered() bool {
  return false
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
