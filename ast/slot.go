package ast

import (
  "encoding/json"
  "fmt"
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
  return fmt.Sprintf("<ast.Slot Name=%s TypeNode=%s Offset=%d>", self.Name, self.TypeNode, self.Offset)
}

func (self Slot) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    TypeNode duck.ITypeNode
    Name string
    Offset int
  }
  x.ClassName = "ast.Slot"
  x.TypeNode = self.TypeNode
  x.Name = self.Name
  x.Offset = self.Offset
  return json.Marshal(x)
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
