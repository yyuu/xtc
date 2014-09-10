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
}

func NewVariableNode(loc core.Location, name string) VariableNode {
  return VariableNode { "ast.VariableNode", loc, name, nil }
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
