package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
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

func (self StructTypeRef) String() string {
  return fmt.Sprintf("struct %s", self.Name)
}

func (self StructTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self StructTypeRef) IsTypeRef() bool {
  return true
}
