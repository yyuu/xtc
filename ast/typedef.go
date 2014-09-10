package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// TypedefNode
type TypedefNode struct {
  ClassName string
  Location core.Location
  NewType core.ITypeRef
  Real core.ITypeRef
  Name string
}

func NewTypedefNode(loc core.Location, real core.ITypeRef, name string) *TypedefNode {
  if real == nil { panic("real is nil") }
  newType := real
  return &TypedefNode { "ast.TypedefNode", loc, newType, real, name }
}

func (self TypedefNode) String() string {
  return fmt.Sprintf("(typedef %s %s)", self.Name, self.Real)
}

func (self TypedefNode) IsTypeDefinition() bool {
  return true
}

func (self TypedefNode) GetLocation() core.Location {
  return self.Location
}

func (self TypedefNode) GetTypeRef() core.ITypeRef {
  return self.NewType
}

func (self TypedefNode) DefiningType() core.IType {
  realTypeNode := NewTypeNode(self.Location, self.Real)
  return typesys.NewUserType(self.Name, realTypeNode, self.Location)
}
