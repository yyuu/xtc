package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// Slot
type Slot struct {
  ClassName string
  TypeNode core.ITypeNode
  Name string
  Offset int
}

func NewSlot(t core.ITypeNode, n string) *Slot {
  if t == nil { panic("t is nil") }
  return &Slot { "ast.Slot", t, n, -1 }
}

func AsSlot(x core.INode) core.ISlot {
  return x.(core.ISlot)
}

func AsSlots(xs []core.INode) []core.ISlot {
  ys := make([]core.ISlot, len(xs))
  for i := range xs {
    ys[i] = xs[i].(core.ISlot)
  }
  return ys
}

func (self Slot) String() string {
  return fmt.Sprintf("<ast.Slot Name=%s TypeNode=%s Offset=%d>", self.Name, self.TypeNode, self.Offset)
}

func (self *Slot) GetName() string {
  return self.Name
}

func (self *Slot) GetOffset() int {
  return self.Offset
}

func (self Slot) GetLocation() core.Location {
  panic("Slot#GetLocation called")
}

func (self *Slot) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *Slot) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *Slot) GetType() core.IType {
  return self.TypeNode.GetType()
}
