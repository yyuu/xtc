package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// ArrayTypeRef
type ArrayTypeRef struct {
  ClassName string
  Location core.Location
  BaseType core.ITypeRef
  Length int
}

func NewArrayTypeRef(baseType core.ITypeRef, length int) *ArrayTypeRef {
  return &ArrayTypeRef { "typesys.ArrayTypeRef", baseType.GetLocation(), baseType, length }
}

func (self ArrayTypeRef) String() string {
  return fmt.Sprintf("<typesys.ArrayTypeRef Location=%s BaseType=%s Length=%d>", self.Location, self.BaseType, self.Length)
}

func (self ArrayTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self ArrayTypeRef) IsTypeRef() bool {
  return true
}
