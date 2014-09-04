package typesys

import (
  "bitbucket.org/yyuu/bs/duck"
)

// PointerType
type PointerType struct {
  PointerSize int
  BaseType duck.IType
}

func NewPointerType(size int, baseType duck.IType) PointerType {
  return PointerType { size, baseType }
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
  Location duck.ILocation
  BaseType duck.ITypeRef
}

func NewPointerTypeRef(baseType duck.ITypeRef) PointerTypeRef {
  return PointerTypeRef { baseType.GetLocation(), baseType }
}

func (self PointerTypeRef) GetLocation() duck.ILocation {
  return self.Location
}

func (self PointerTypeRef) IsTypeRef() bool {
  return true
}
