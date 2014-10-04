package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// StructTypeRef
type StructTypeRef struct {
  ClassName string
  Location core.Location
  Name string
}

func NewStructTypeRef(loc core.Location, name string) *StructTypeRef {
  return &StructTypeRef { "typesys.StructTypeRef", loc, name }
}

func (self StructTypeRef) Key() string {
  return fmt.Sprintf("struct %s", self.Name)
}

func (self StructTypeRef) String() string {
  return self.Key()
}

func (self StructTypeRef) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self StructTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self StructTypeRef) IsTypeRef() bool {
  return true
}
