package ast

// Slot
type Slot struct {
  TypeNode ITypeNode
  Name string
  Offset int
}

func NewSlot(t ITypeNode, n string) Slot {
  return Slot { t, n, -1 }
}

func (self Slot) GetName() string {
  return self.Name
}

func (self Slot) GetOffset() int {
  return self.Offset
}
