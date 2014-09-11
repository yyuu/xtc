package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type Parameter struct {
  *DefinedVariable
}

func NewParameter(t core.ITypeNode, name string) *Parameter {
  return &Parameter { &DefinedVariable { "entity.Parameter", true, name, t, 0, nil } }
}

func NewParameters(xs...*Parameter) []*Parameter {
  if 0 < len(xs) {
    return xs
  } else {
    return []*Parameter { }
  }
}

func (self Parameter) String() string {
  return fmt.Sprintf("<entity.Parameter Name=%s TypeNode=%s>", self.DefinedVariable.Name, self.DefinedVariable.TypeNode)
}

func (self Parameter) IsPrivate() bool {
  return false
}

func (self Parameter) IsConstant() bool {
  return false
}

func (self Parameter) GetInitializer() core.IExprNode {
  return nil
}

func (self *Parameter) SetInitializer(e core.IExprNode) {
  // noop
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
  return self.DefinedVariable.TypeNode
}

func (self Parameter) GetName() string {
  return self.DefinedVariable.Name
}
