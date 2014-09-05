package typesys

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// IntegerType
type IntegerType struct {
  IntegerSize int
  Signed bool
  Name string
}

func NewIntegerType(size int, isSigned bool, name string) IntegerType {
  return IntegerType { size, isSigned, name }
}

func (self IntegerType) String() string {
  return fmt.Sprintf("<typesys.IntegerType Name=%s IntegerSize=%d Signed=%v>", self.Name, self.IntegerSize, self.Signed)
}

func (self IntegerType) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    IntegerSize int
    Signed bool
    Name string
  }
  x.ClassName = "typesys.IntegerType"
  x.IntegerSize = self.IntegerSize
  x.Signed = self.Signed
  x.Name = self.Name
  return json.Marshal(x)
}

func (self IntegerType) Size() int {
  return self.IntegerSize
}

func (self IntegerType) AllocSize() int {
  return self.Size()
}

func (self IntegerType) Alignment() int {
  return self.AllocSize()
}

func (self IntegerType) IsVoid() bool {
  return false
}

func (self IntegerType) IsInteger() bool {
  return true
}

func (self IntegerType) IsSigned() bool {
  return self.Signed
}

func (self IntegerType) IsPointer() bool {
  return false
}

func (self IntegerType) IsArray() bool {
  return false
}

func (self IntegerType) IsCompositeType() bool {
  return false
}

func (self IntegerType) IsStruct() bool {
  return false
}

func (self IntegerType) IsUnion() bool {
  return false
}

func (self IntegerType) IsUserType() bool {
  return false
}

func (self IntegerType) IsFunction() bool {
  return false
}

// IntegerTypeRef
type IntegerTypeRef struct {
  Location duck.ILocation
  Name string
}

func NewIntegerTypeRef(location duck.ILocation, name string) IntegerTypeRef {
  return IntegerTypeRef { location, name }
}

func (self IntegerTypeRef) String() string {
  return fmt.Sprintf("<typesys.IntegerTypeRef Name=%s Location=%s>", self.Name, self.Location)
}

func (self IntegerTypeRef) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
  }
  x.ClassName = "typesys.IntegerTypeRef"
  x.Location = self.Location
  x.Name = self.Name
  return json.Marshal(x)
}

func (self IntegerTypeRef) GetLocation() duck.ILocation {
  return self.Location
}

func (self IntegerTypeRef) IsTypeRef() bool {
  return true
}
