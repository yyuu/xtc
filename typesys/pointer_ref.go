package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// PointerTypeRef
type PointerTypeRef struct {
  ClassName string
  Location core.Location
  BaseType core.ITypeRef
}

func NewPointerTypeRef(baseType core.ITypeRef) *PointerTypeRef {
  return &PointerTypeRef { "typesys.PointerTypeRef", baseType.GetLocation(), baseType }
}

func (self PointerTypeRef) String() string {
  return fmt.Sprintf("<typesys.PointerTypeRef Location=%s BaseType=%s>", self.Location, self.BaseType)
}

func (self PointerTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self PointerTypeRef) IsTypeRef() bool {
  return true
}
