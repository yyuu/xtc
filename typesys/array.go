package typesys

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// ArrayType
type ArrayType struct {
  BaseType duck.IType
  Length int
  PointerSize int
}

func NewArrayType(baseType duck.IType, length int, pointerSize int) ArrayType {
  return ArrayType { baseType, length, pointerSize }
}

func (self ArrayType) String() string {
  return fmt.Sprintf("<typesys.ArrayType BaseType=%s Length=%d PointerSize=%d>", self.BaseType, self.Length, self.PointerSize)
}

func (self ArrayType) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    BaseType duck.IType
    Length int
    PointerSize int
  }
  x.ClassName = "typesys.ArrayType"
  x.BaseType = self.BaseType
  x.Length = self.Length
  x.PointerSize = self.PointerSize
  return json.Marshal(x)
}

func (self ArrayType) Size() int {
  return self.PointerSize
}

func (self ArrayType) AllocSize() int {
  return self.BaseType.AllocSize() * self.Length
}

func (self ArrayType) Alignment() int {
  return self.BaseType.AllocSize()
}

func (self ArrayType) IsVoid() bool {
  return false
}

func (self ArrayType) IsInteger() bool {
  return false
}

func (self ArrayType) IsSigned() bool {
  return false
}

func (self ArrayType) IsPointer() bool {
  return false
}

func (self ArrayType) IsArray() bool {
  return true
}

func (self ArrayType) IsCompositeType() bool {
  return false
}

func (self ArrayType) IsStruct() bool {
  return false
}

func (self ArrayType) IsUnion() bool {
  return false
}

func (self ArrayType) IsUserType() bool {
  return false
}

func (self ArrayType) IsFunction() bool {
  return false
}

// ArrayTypeRef
type ArrayTypeRef struct {
  Location duck.ILocation
  BaseType duck.ITypeRef
  Length int
}

func NewArrayTypeRef(baseType duck.ITypeRef, length int) ArrayTypeRef {
  return ArrayTypeRef { baseType.GetLocation(), baseType, length }
}

func (self ArrayTypeRef) String() string {
  return fmt.Sprintf("<typesys.ArrayTypeRef Location=%s BaseType=%s Length=%d>", self.Location, self.BaseType, self.Length)
}

func (self ArrayTypeRef) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    BaseType duck.ITypeRef
    Length int
  }
  x.ClassName = "typesys.ArrayTypeRef"
  x.Location = self.Location
  x.BaseType = self.BaseType
  x.Length = self.Length
  return json.Marshal(x)
}

func (self ArrayTypeRef) GetLocation() duck.ILocation {
  return self.Location
}

func (self ArrayTypeRef) IsTypeRef() bool {
  return true
}
