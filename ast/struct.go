package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// StructNode
type StructNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Name string
  Members []core.ISlot
}

func NewStructNode(loc core.Location, ref core.ITypeRef, name string, membs []core.ISlot) *StructNode {
  if ref == nil { panic("ref is nil") }
  return &StructNode { "ast.StructNode", loc, NewTypeNode(loc, ref), name, membs }
}

func NewStructNodes(xs...*StructNode) []*StructNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []*StructNode { }
  }
}

func (self StructNode) String() string {
  return fmt.Sprintf("<ast.StructNode Name=%s location=%s typeNode=%s members=%s>", self.Name, self.Location, self.TypeNode, self.Members)
}

func (self StructNode) IsTypeDefinition() bool {
  return true
}

func (self StructNode) GetLocation() core.Location {
  return self.Location
}

func (self StructNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self StructNode) DefiningType() core.IType {
  var membs []core.ISlot
  for i := range self.Members {
    membs[i] = self.Members[i]
  }
  return typesys.NewStructType(self.Name, membs, self.Location)
}
