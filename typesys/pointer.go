package typesys

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// PointerType
type PointerType struct {
  pointerSize int
  baseType duck.IType
}

func NewPointerType(size int, baseType duck.IType) PointerType {
  return PointerType { size, baseType }
}

func (self PointerType) String() string {
  return fmt.Sprintf("<typesys.PointerType PointerSize=%d BaseType=%s>", self.pointerSize, self.baseType)
}

func (self PointerType) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    PointerSize int
    BaseType duck.IType
  }
  x.ClassName = "typesys.PointerType"
  x.PointerSize = self.pointerSize
  x.BaseType = self.baseType
  return json.Marshal(x)
}

func (self PointerType) Size() int {
  return self.pointerSize
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
  location duck.ILocation
  baseType duck.ITypeRef
}

func NewPointerTypeRef(baseType duck.ITypeRef) PointerTypeRef {
  return PointerTypeRef { baseType.GetLocation(), baseType }
}

func (self PointerTypeRef) String() string {
  return fmt.Sprintf("<typesys.PointerTypeRef Location=%s BaseType=%s>", self.location, self.baseType)
}

func (self PointerTypeRef) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    BaseType duck.ITypeRef
  }
  x.ClassName = "typesys.PointerTypeRef"
  x.Location = self.location
  x.BaseType = self.baseType
  return json.Marshal(x)
}

func (self PointerTypeRef) GetLocation() duck.ILocation {
  return self.location
}

func (self PointerTypeRef) IsTypeRef() bool {
  return true
}
