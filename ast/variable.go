package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

// VariableNode
type VariableNode struct {
  ClassName string
  Location core.Location
  Name string
  entity core.IEntity
  t core.IType
}

func NewVariableNode(loc core.Location, name string) *VariableNode {
  return &VariableNode { "ast.VariableNode", loc, name, nil, nil }
}

func (self VariableNode) String() string {
  return self.Name
}

func (self VariableNode) IsExprNode() bool {
  return true
}

func (self VariableNode) GetLocation() core.Location {
  return self.Location
}

func (self VariableNode) GetName() string {
  return self.Name
}

func (self *VariableNode) SetEntity(ent core.IEntity) {
  self.entity = ent
}

func (self VariableNode) GetEntity() core.IEntity {
  return self.entity
}

func (self VariableNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *VariableNode) SetType(t core.IType) {
  self.t = t
}
