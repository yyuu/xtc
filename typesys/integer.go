package typesys

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// IntegerType
type IntegerType struct {
  integerSize int
  signed bool
  name string
}

func NewIntegerType(size int, isSigned bool, name string) IntegerType {
  return IntegerType { size, isSigned, name }
}

func (self IntegerType) String() string {
  return fmt.Sprintf("<typesys.IntegerType Name=%s IntegerSize=%d Signed=%v>", self.name, self.integerSize, self.signed)
}

func (self IntegerType) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    IntegerSize int
    Signed bool
    Name string
  }
  x.ClassName = "typesys.IntegerType"
  x.IntegerSize = self.integerSize
  x.Signed = self.signed
  x.Name = self.name
  return json.Marshal(x)
}

func (self IntegerType) Size() int {
  return self.integerSize
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
  return self.signed
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
  location duck.ILocation
  name string
}

func NewIntegerTypeRef(loc duck.ILocation, name string) IntegerTypeRef {
  return IntegerTypeRef { loc, name }
}

func (self IntegerTypeRef) String() string {
  return fmt.Sprintf("<typesys.IntegerTypeRef Name=%s Location=%s>", self.name, self.location)
}

func (self IntegerTypeRef) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
  }
  x.ClassName = "typesys.IntegerTypeRef"
  x.Location = self.location
  x.Name = self.name
  return json.Marshal(x)
}

func (self IntegerTypeRef) GetLocation() duck.ILocation {
  return self.location
}

func (self IntegerTypeRef) IsTypeRef() bool {
  return true
}
