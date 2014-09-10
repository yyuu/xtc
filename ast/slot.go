package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// Slot
type Slot struct {
  ClassName string
  TypeNode duck.ITypeNode
  Name string
  Offset int
}

func NewSlot(t duck.ITypeNode, n string) Slot {
  if t == nil { panic("t is nil") }
  return Slot { "ast.Slot", t, n, -1 }
}

func (self Slot) String() string {
  return fmt.Sprintf("<ast.Slot Name=%s TypeNode=%s Offset=%d>", self.Name, self.TypeNode, self.Offset)
}

func (self Slot) GetName() string {
  return self.Name
}

func (self Slot) GetOffset() int {
  return self.Offset
}

func (self Slot) GetLocation() duck.Location {
  panic("Slot#GetLocation called")
}
