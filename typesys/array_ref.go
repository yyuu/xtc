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

func (self ArrayTypeRef) Key() string {
  if self.Length < 1 {
    return fmt.Sprintf("%s[]", self.BaseType)
  } else {
    return fmt.Sprintf("%s[%d]", self.BaseType, self.Length)
  }
}

func (self ArrayTypeRef) String() string {
  return self.Key()
}

func (self ArrayTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self ArrayTypeRef) IsTypeRef() bool {
  return true
}

func (self ArrayTypeRef) GetBaseType() core.ITypeRef {
  return self.BaseType
}

func (self ArrayTypeRef) GetLength() int {
  return self.Length
}
