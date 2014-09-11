package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type DefinedVariable struct {
  ClassName string
  Private bool
  Name string
  TypeNode core.ITypeNode
  NumRefered int
  Initializer core.IExprNode
}

func NewDefinedVariable(isPrivate bool, t core.ITypeNode, name string, init core.IExprNode) *DefinedVariable {
  return &DefinedVariable { "entity.DefinedVariable", isPrivate, name, t, 0, init }
}

func NewDefinedVariables(xs...*DefinedVariable) []*DefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []*DefinedVariable { }
  }
}

func (self DefinedVariable) String() string {
  return fmt.Sprintf("<entity.DefinedVariable Name=%s Private=%v TypeNode=%s NumRefered=%d Initializer=%s>", self.Name, self.Private, self.TypeNode, self.NumRefered, self.Initializer)
}

func (self DefinedVariable) IsDefined() bool {
  return true
}

func (self DefinedVariable) HasInitializer() bool {
  return self.Initializer != nil
}

func (self DefinedVariable) IsPrivate() bool {
  return self.Private
}

func (self DefinedVariable) IsConstant() bool {
  return false
}

func (self DefinedVariable) GetName() string {
  return self.Name
}

func (self DefinedVariable) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self DefinedVariable) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self DefinedVariable) GetNumRefered() int {
  return self.NumRefered
}

func (self DefinedVariable) IsRefered() bool {
  return 0 < self.NumRefered
}

func (self *DefinedVariable) Refered() {
  self.NumRefered++
}

func (self DefinedVariable) GetInitializer() core.IExprNode {
  return self.Initializer
}

func (self *DefinedVariable) SetInitializer(init core.IExprNode) {
  self.Initializer = init
}
