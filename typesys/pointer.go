package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// PointerType
type PointerType struct {
  ClassName string
  PointerSize int
  BaseType duck.IType
}

func NewPointerType(size int, baseType duck.IType) *PointerType {
  return &PointerType { "typesys.PointerType", size, baseType }
}

func (self PointerType) String() string {
  return fmt.Sprintf("<typesys.PointerType PointerSize=%d BaseType=%s>", self.PointerSize, self.BaseType)
}

func (self PointerType) Size() int {
  return self.PointerSize
}

func (self PointerType) AllocSize() int {
  return self.Size()
}

func (self PointerType) Alignment() int {
  return self.AllocSize()
}

func (self PointerType) IsVoid() bool {
  return false
}

func (self PointerType) IsInteger() bool {
  return false
}

func (self PointerType) IsSigned() bool {
  return false
}

func (self PointerType) IsPointer() bool {
  return true
}

func (self PointerType) IsArray() bool {
  return false
}

func (self PointerType) IsCompositeType() bool {
  return false
}

func (self PointerType) IsStruct() bool {
  return false
}

func (self PointerType) IsUnion() bool {
  return false
}

func (self PointerType) IsUserType() bool {
  return false
}

func (self PointerType) IsFunction() bool {
  return false
}

// PointerTypeRef
type PointerTypeRef struct {
  ClassName string
  Location duck.Location
  BaseType duck.ITypeRef
}

func NewPointerTypeRef(baseType duck.ITypeRef) PointerTypeRef {
  return PointerTypeRef { "typesys.PointerTypeRef", baseType.GetLocation(), baseType }
}

func (self PointerTypeRef) String() string {
  return fmt.Sprintf("<typesys.PointerTypeRef Location=%s BaseType=%s>", self.Location, self.BaseType)
}

func (self PointerTypeRef) GetLocation() duck.Location {
  return self.Location
}

func (self PointerTypeRef) IsTypeRef() bool {
  return true
}
