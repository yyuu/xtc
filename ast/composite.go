package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// StructNode
type StructNode struct {
  location duck.ILocation
  typeNode duck.ITypeNode
  name string
  members []Slot
}

func NewStructNode(loc duck.ILocation, ref duck.ITypeRef, name string, membs []Slot) StructNode {
  if loc == nil { panic("location is nil") }
  if ref == nil { panic("ref is nil") }
  return StructNode { loc, NewTypeNode(loc, ref), name, membs }
}

func (self StructNode) String() string {
  return fmt.Sprintf("<ast.StructNode Name=%s location=%s typeNode=%s members=%s>", self.name, self.location, self.typeNode, self.members)
}

func (self StructNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    TypeNode duck.ITypeNode
    Name string
    Members []Slot
  }
  x.ClassName = "ast.StructNode"
  x.Location = self.location
  x.TypeNode = self.typeNode
  x.Name = self.name
  x.Members = self.members
  return json.Marshal(x)
}

func (self StructNode) IsTypeDefinition() bool {
  return true
}

func (self StructNode) GetLocation() duck.ILocation {
  return self.location
}

// UnionNode
type UnionNode struct {
  location duck.ILocation
  typeNode duck.ITypeNode
  name string
  members []Slot
}

func NewUnionNode(loc duck.ILocation, ref duck.ITypeRef, name string, membs []Slot) UnionNode {
  if loc == nil { panic("location is nil") }
  if ref == nil { panic("ref is nil") }
  return UnionNode { loc, NewTypeNode(loc, ref), name, membs }
}

func (self UnionNode) String() string {
  return fmt.Sprintf("<ast.UnionNode name=%s location=%s typeNode=%s members=%s>", self.name, self.location, self.typeNode, self.members)
}

func (self UnionNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    TypeNode duck.ITypeNode
    Name string
    Members []Slot
  }
  x.ClassName = "ast.UnionNode"
  x.Location = self.location
  x.TypeNode = self.typeNode
  x.Name = self.name
  x.Members = self.members
  return json.Marshal(x)
}

func (self UnionNode) IsTypeDefinition() bool {
  return true
}

func (self UnionNode) GetLocation() duck.ILocation {
  return self.location
}
