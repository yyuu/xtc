package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type Parameter struct {
  ClassName string
  TypeNode core.ITypeNode
  Name string
}

func NewParameter(t core.ITypeNode, name string) *Parameter {
  return &Parameter { "entity.Parameter", t, name }
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
