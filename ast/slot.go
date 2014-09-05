package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// Slot
type Slot struct {
  typeNode duck.ITypeNode
  name string
  offset int
}

func NewSlot(t duck.ITypeNode, n string) Slot {
  if t == nil { panic("t is nil") }
  return Slot { t, n, -1 }
}

func (self Slot) String() string {
  return fmt.Sprintf("<ast.Slot Name=%s TypeNode=%s Offset=%d>", self.name, self.typeNode, self.offset)
}

func (self Slot) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    TypeNode duck.ITypeNode
    Name string
    Offset int
  }
  x.ClassName = "ast.Slot"
  x.TypeNode = self.typeNode
  x.Name = self.name
  x.Offset = self.offset
  return json.Marshal(x)
}

func (self Slot) GetName() string {
  return self.name
}

func (self Slot) GetOffset() int {
  return self.offset
}

func (self Slot) GetLocation() duck.ILocation {
  panic("Slot#GetLocation called")
}
