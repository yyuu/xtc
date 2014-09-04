package typesys

import (
  "encoding/json"
  "bitbucket.org/yyuu/bs/duck"
)

// VoidType
type VoidType struct {
}

func NewVoidType() VoidType {
  return VoidType { }
}

func (self VoidType) String() string {
  panic("VoidType#String called")
}

func (self VoidType) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
  }
  x.ClassName = "typesys.VoidType"
  return json.Marshal(x)
}

func (self VoidType) Size() int {
  return 1
}

func (self VoidType) AllocSize() int {
  return self.Size()
}

func (self VoidType) Alignment() int {
  return self.AllocSize()
}

func (self VoidType) IsVoid() bool {
  return true
}

func (self VoidType) IsInteger() bool {
  return false
}

func (self VoidType) IsSigned() bool {
  return false
}

func (self VoidType) IsPointer() bool {
  return false
}

func (self VoidType) IsArray() bool {
  return false
}

func (self VoidType) IsCompositeType() bool {
  return false
}

func (self VoidType) IsStruct() bool {
  return false
}

func (self VoidType) IsUnion() bool {
  return false
}

func (self VoidType) IsUserType() bool {
  return false
}

func (self VoidType) IsFunction() bool {
  return false
}

// VoidTypeRef
type VoidTypeRef struct {
  Location duck.ILocation
}

func NewVoidTypeRef(location duck.ILocation) VoidTypeRef {
  return VoidTypeRef { location }
}

func (self VoidTypeRef) String() string {
  panic("VoidTypeRev#String called")
}

func (self VoidTypeRef) MarshalJSON() ([]byte, error) {
  var x struct {
    Classname string
    Location duck.ILocation
  }
  x.Classname = "typesys.VoidTypeRef"
  x.Location = self.Location
  return json.Marshal(x)
}

func (self VoidTypeRef) GetLocation() duck.ILocation {
  return self.Location
}

func (self VoidTypeRef) IsTypeRef() bool {
  return true
}
