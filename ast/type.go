package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// TypeNode
type TypeNode struct {
  ClassName string
  Location core.Location
  TypeRef core.ITypeRef
  Type core.IType
}

func NewTypeNode(loc core.Location, ref core.ITypeRef) *TypeNode {
  if ref == nil { panic("ref is nil") }
  return &TypeNode { "ast.TypeNode", loc, ref, nil }
}

func (self TypeNode) String() string {
  return fmt.Sprintf("(type %s)", self.TypeRef)
}

func (self TypeNode) GetTypeRef() core.ITypeRef {
  return self.TypeRef
}

func (self *TypeNode) AsTypeNode() core.ITypeNode {
  return self
}

func (self TypeNode) GetLocation() core.Location {
  return self.Location
}

func (self TypeNode) IsResolved() bool {
  return self.Type != nil
}

func (self TypeNode) GetType() core.IType {
  if self.Type == nil {
    panic(fmt.Errorf("%s type not resolved", self.Location))
  }
  return self.Type
}

func (self *TypeNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}
