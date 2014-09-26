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
  Entity core.IEntity
}

func NewVariableNode(loc core.Location, name string) *VariableNode {
  return &VariableNode { "ast.VariableNode", loc, name, nil }
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

func (self *VariableNode) IsResolved() bool {
  return self.Entity != nil
}

func (self *VariableNode) GetName() string {
  return self.Name
}

func (self *VariableNode) SetEntity(ent core.IEntity) {
  self.Entity = ent
}

func (self *VariableNode) GetEntity() core.IEntity {
  if self.Entity == nil {
    panic(fmt.Errorf("%s entity is nil: %s", self.Location, self.Name))
  }
  return self.Entity
}

func (self *VariableNode) GetType() core.IType {
  return self.GetOrigType()
}

func (self *VariableNode) SetType(t core.IType) {
//// FIXME: uncomment following causes "type mismatch: int != int*"
//if ! self.GetType().IsCompatible(t) {
//  panic(fmt.Sprintf("type mismatch: %s != %s", self.GetEntity().GetType(), t))
//}
}

func (self *VariableNode) GetOrigType() core.IType {
  return self.GetEntity().GetType()
}

func (self *VariableNode) IsConstant() bool {
  return false
}

func (self *VariableNode) IsParameter() bool {
  return self.GetEntity().IsParameter()
}

func (self *VariableNode) IsLvalue() bool {
  return true
}

func (self *VariableNode) IsAssignable() bool {
  return true
}

func (self *VariableNode) IsLoadable() bool {
  t := self.GetOrigType()
  return !t.IsArray() && !t.IsFunction()
}

func (self *VariableNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *VariableNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
