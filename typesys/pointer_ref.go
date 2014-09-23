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

func (self PointerTypeRef) Key() string {
  return fmt.Sprintf("%s*", self.BaseType)
}

func (self PointerTypeRef) String() string {
  return self.Key()
}

func (self PointerTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self PointerTypeRef) IsTypeRef() bool {
  return true
}

func (self PointerTypeRef) GetBaseType() core.ITypeRef {
  return self.BaseType
}
