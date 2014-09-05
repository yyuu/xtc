package typesys

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// StructType
type StructType struct {
  Location duck.ILocation
  Name string
  Members []duck.ISlot
}

func NewStructType(name string, membs []duck.ISlot, loc duck.ILocation) StructType {
  return StructType { loc, name, membs }
}

func (self StructType) String() string {
  return fmt.Sprintf("<typesys.StructType Name=%s Location=%s Members=%s>", self.Name, self.Location, self.Members)
}

func (self StructType) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
    Members []duck.ISlot
  }
  x.ClassName = "typesys.StructType"
  x.Location = self.Location
  x.Name = self.Name
  x.Members = self.Members
  return json.Marshal(x)
}

func (self StructType) Size() int {
  panic("StructType#Size called")
}

func (self StructType) AllocSize() int {
  panic("StructType#AllocSize called")
}

func (self StructType) Alignment() int {
  panic("StructType#Alignment called")
}

func (self StructType) IsVoid() bool {
  return false
}

func (self StructType) IsInteger() bool {
  return false
}

func (self StructType) IsSigned() bool {
  return false
}

func (self StructType) IsPointer() bool {
  return false
}

func (self StructType) IsArray() bool {
  return false
}

func (self StructType) IsCompositeType() bool {
  return false
}

func (self StructType) IsStruct() bool {
  return true
}

func (self StructType) IsUnion() bool {
  return false
}

func (self StructType) IsUserType() bool {
  return false
}

func (self StructType) IsFunction() bool {
  return false
}

// StructTypeRef
type StructTypeRef struct {
  Location duck.ILocation
  Name string
}

func NewStructTypeRef(loc duck.ILocation, name string) StructTypeRef {
  return StructTypeRef { loc, name }
}

func (self StructTypeRef) String() string {
  return fmt.Sprintf("<typesys.StructTypeRef Name=%s Location=%s>", self.Name, self.Location)
}

func (self StructTypeRef) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
  }
  x.ClassName = "typesys.StructTypeRef"
  x.Location = self.Location
  x.Name = self.Name
  return json.Marshal(x)
}

func (self StructTypeRef) GetLocation() duck.ILocation {
  return self.Location
}

func (self StructTypeRef) IsTypeRef() bool {
  return true
}
