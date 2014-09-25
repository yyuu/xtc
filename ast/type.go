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
  t core.IType
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
  return self.t != nil
}

func (self TypeNode) GetType() core.IType {
  return self.t
}

func (self *TypeNode) SetType(t core.IType) {
  if self.t != nil {
    panic("TypeNode#SetType called twice")
  }
  self.t = t
}
