package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// UnionNode
type UnionNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Name string
  Members []core.ISlot
}

func NewUnionNode(loc core.Location, ref core.ITypeRef, name string, membs []core.ISlot) *UnionNode {
  if ref == nil { panic("ref is nil") }
  return &UnionNode { "ast.UnionNode", loc, NewTypeNode(loc, ref), name, membs }
}

func (self UnionNode) String() string {
  return fmt.Sprintf("<ast.UnionNode name=%s location=%s typeNode=%s members=%s>", self.Name, self.Location, self.TypeNode, self.Members)
}

func (self UnionNode) IsTypeDefinition() bool {
  return true
}

func (self UnionNode) GetLocation() core.Location {
  return self.Location
}

func (self UnionNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self UnionNode) DefiningType() core.IType {
  var membs []core.ISlot
  for i := range self.Members {
    membs[i] = self.Members[i]
  }
  return typesys.NewUnionType(self.Name, membs, self.Location)
}
