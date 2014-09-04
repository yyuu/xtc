package ast

import (
  "bitbucket.org/yyuu/bs/duck"
)

// Slot
type Slot struct {
  TypeNode duck.ITypeNode
  Name string
  Offset int
}

func NewSlot(t duck.ITypeNode, n string) Slot {
  return Slot { t, n, -1 }
}

func (self Slot) String() string {
  panic("Slot#String called")
}

func (self Slot) MarshalJSON() ([]byte, error) {
  panic("Slot#MarshalJSON called")
}

func (self Slot) GetName() string {
  return self.Name
}

func (self Slot) GetOffset() int {
  return self.Offset
}

func (self Slot) GetLocation() duck.ILocation {
  panic("Slot#GetLocation called")
}
