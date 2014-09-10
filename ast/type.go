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
}

func NewTypeNode(loc core.Location, t core.ITypeRef) TypeNode {
  if t == nil { panic("t is nil") }
  return TypeNode { "ast.TypeNode", loc, t }
}

func (self TypeNode) String() string {
  return fmt.Sprintf("(type %s)", self.TypeRef)
}

func (self TypeNode) GetTypeRef() core.ITypeRef {
  return self.TypeRef
}

func (self TypeNode) IsTypeNode() bool {
  return true
}

func (self TypeNode) GetLocation() core.Location {
  return self.Location
}
