package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type Params struct {
  ClassName string
  Location core.Location
  ParamDescs []core.IParameter
}

func NewParams(loc core.Location, paramDescs []core.IParameter) Params {
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

func (self Params) GetLocation() core.Location {
  return self.Location
}

func (self Params) GetParamDescs() []core.IParameter {
  return self.ParamDescs
}

func (self Params) GetName() string {
  panic("Parameter#GetName called")
}

type Parameter struct {
  ClassName string
  TypeNode core.ITypeNode
  Name string
}

func NewParameter(t core.ITypeNode, name string) Parameter {
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

func (self Parameter) GetInitializer() core.IExprNode {
  return nil
}

func (self Parameter) SetInitializer(e core.IExprNode) core.IDefinedVariable {
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

func (self Parameter) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self Parameter) GetName() string {
  return self.Name
}
