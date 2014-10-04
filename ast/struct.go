package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/typesys"
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

func AsStructNode(x core.INode) *StructNode {
  return x.(*StructNode)
}

func (self StructNode) String() string {
  return fmt.Sprintf("<ast.StructNode Name=%s location=%s typeNode=%s members=%s>", self.Name, self.Location, self.TypeNode, self.Members)
}

func (self *StructNode) IsTypeDefinition() bool {
  return true
}

func (self *StructNode) IsCompositeType() bool {
  return true
}

func (self StructNode) GetLocation() core.Location {
  return self.Location
}

func (self *StructNode) GetName() string {
  return self.Name
}

func (self *StructNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *StructNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *StructNode) GetMembers() []core.ISlot {
  return self.Members
}

func (self *StructNode) DefiningType() core.IType {
  membs := make([]core.ISlot, len(self.Members))
  for i := range self.Members {
    membs[i] = self.Members[i]
  }
  return typesys.NewStructType(self.Name, membs, self.Location)
}
