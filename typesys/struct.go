package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// StructType
type StructType struct {
  ClassName string
  Location core.Location
  Name string
  Members []core.ISlot
}

func NewStructType(name string, membs []core.ISlot, loc core.Location) *StructType {
  return &StructType { "typesys.StructType", loc, name, membs }
}

func (self StructType) String() string {
  return fmt.Sprintf("<typesys.StructType Name=%s Location=%s Members=%s>", self.Name, self.Location, self.Members)
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
