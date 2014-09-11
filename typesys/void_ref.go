package typesys

import (
  "bitbucket.org/yyuu/bs/core"
)

// VoidTypeRef
type VoidTypeRef struct {
  ClassName string
  Location core.Location
}

func NewVoidTypeRef(loc core.Location) *VoidTypeRef {
  return &VoidTypeRef { "typesys.VoidTypeRef", loc }
}

func (self VoidTypeRef) String() string {
  return "void"
}

func (self VoidTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self VoidTypeRef) IsTypeRef() bool {
  return true
}
