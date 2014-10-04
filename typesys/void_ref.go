package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// VoidTypeRef
type VoidTypeRef struct {
  ClassName string
  Location core.Location
}

func NewVoidTypeRef(loc core.Location) *VoidTypeRef {
  return &VoidTypeRef { "typesys.VoidTypeRef", loc }
}

func (self VoidTypeRef) Key() string {
  return "void"
}

func (self VoidTypeRef) String() string {
  return self.Key()
}

func (self VoidTypeRef) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self VoidTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self VoidTypeRef) IsTypeRef() bool {
  return true
}
