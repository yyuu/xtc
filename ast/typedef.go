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
  TypeNode core.ITypeNode
  RealTypeNode core.ITypeNode
  Name string
}

func NewTypedefNode(loc core.Location, real core.ITypeRef, name string) *TypedefNode {
  if real == nil { panic("real is nil") }
  t := NewTypeNode(loc, typesys.NewUserTypeRef(loc, name))
  return &TypedefNode { "ast.TypedefNode", loc, t, NewTypeNode(loc, real), name }
}

func NewTypedefNodes(xs...*TypedefNode) []*TypedefNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []*TypedefNode { }
  }
}

func (self TypedefNode) String() string {
  return fmt.Sprintf("(typedef %s %s)", self.Name, self.RealTypeNode)
}

func (self TypedefNode) IsTypeDefinition() bool {
  return true
}

func (self TypedefNode) GetLocation() core.Location {
  return self.Location
}

func (self TypedefNode) GetName() string {
  return self.Name
}

func (self TypedefNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self TypedefNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self TypedefNode) GetRealTypeNode() core.ITypeNode {
  return self.RealTypeNode
}

func (self TypedefNode) GetRealTypeRef() core.ITypeRef {
  return self.RealTypeNode.GetTypeRef()
}

func (self TypedefNode) DefiningType() core.IType {
  return typesys.NewUserType(self.Name, self.RealTypeNode, self.Location)
}
