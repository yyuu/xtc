package ast

import (
  "fmt"
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

func (self *VariableNode) AsExprNode() core.IExprNode {
  return self
}

func (self VariableNode) GetLocation() core.Location {
  return self.Location
}

func (self VariableNode) IsResolved() bool {
  return self.entity != nil
}

func (self VariableNode) GetName() string {
  return self.Name
}

func (self *VariableNode) SetEntity(ent core.IEntity) {
  self.entity = ent
}

func (self VariableNode) GetEntity() core.IEntity {
  if self.entity == nil {
    panic(fmt.Errorf("%s entity is nil: %s", self.Location, self.Name))
  }
  return self.entity
}

func (self VariableNode) GetType() core.IType {
  return self.GetEntity().GetType()
}

func (self *VariableNode) SetType(t core.IType) {
  self.t = t
}

func (self VariableNode) IsConstant() bool {
  return false
}

func (self VariableNode) IsParameter() bool {
  return self.entity.IsParameter()
}

func (self VariableNode) IsLvalue() bool {
  return true
}

func (self VariableNode) IsAssignable() bool {
  return true
}

func (self VariableNode) IsLoadable() bool {
  return false
}

func (self VariableNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self VariableNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
