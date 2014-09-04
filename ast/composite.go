package ast

import (
  "bitbucket.org/yyuu/bs/duck"
)

// StructNode
type StructNode struct {
  Location duck.ILocation
  TypeNode duck.ITypeNode
  Name string
  Members []Slot
}

func NewStructNode(loc duck.ILocation, ref duck.ITypeRef, name string, membs []Slot) StructNode {
  return StructNode { loc, NewTypeNode(loc, ref), name, membs }
}

func (self StructNode) String() string {
  panic("not implemented")
}

func (self StructNode) MarshalJSON() ([]byte, error) {
  panic("not implemented")
}

func (self StructNode) IsTypeDefinition() bool {
  return true
}

func (self StructNode) GetLocation() duck.ILocation {
  return self.Location
}

// UnionNode
type UnionNode struct {
  Location duck.ILocation
  TypeNode duck.ITypeNode
  Name string
  Members []Slot
}

func NewUnionNode(loc duck.ILocation, ref duck.ITypeRef, name string, membs []Slot) UnionNode {
  return UnionNode { loc, NewTypeNode(loc, ref), name, membs }
}

func (self UnionNode) String() string {
  panic("not implemented")
}

func (self UnionNode) MarshalJSON() ([]byte, error) {
  panic("not implemented")
}

func (self UnionNode) IsTypeDefinition() bool {
  return true
}

func (self UnionNode) GetLocation() duck.ILocation {
  return self.Location
}
