package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// UnionType
type UnionType struct {
  ClassName string
  Location core.Location
  Name string
  Members []core.ISlot
}

func NewUnionType(name string, membs []core.ISlot, loc core.Location) *UnionType {
  return &UnionType { "typesys.UnionType", loc, name, membs }
}

func (self UnionType) String() string {
  return fmt.Sprintf("<typesys.UnionType Name=%s Location=%s Members=%s>", self.Name, self.Location, self.Members)
}

func (self UnionType) Size() int {
  panic("UnionType#Size called")
}

func (self UnionType) AllocSize() int {
  panic("UnionType#AllocSize called")
}

func (self UnionType) Alignment() int {
  panic("UnionType#Alignment called")
}

func (self UnionType) IsVoid() bool {
  return false
}

func (self UnionType) IsInteger() bool {
  return false
}

func (self UnionType) IsSigned() bool {
  return false
}

func (self UnionType) IsPointer() bool {
  return false
}

func (self UnionType) IsArray() bool {
  return false
}

func (self UnionType) IsCompositeType() bool {
  return false
}

func (self UnionType) IsStruct() bool {
  return false
}

func (self UnionType) IsUnion() bool {
  return true
}

func (self UnionType) IsUserType() bool {
  return false
}

func (self UnionType) IsFunction() bool {
  return false
}

// UnionTypeRef
type UnionTypeRef struct {
  ClassName string
  Location core.Location
  Name string
}

func NewUnionTypeRef(loc core.Location, name string) UnionTypeRef {
  return UnionTypeRef { "typesys.UnionTypeRef", loc, name }
}

func (self UnionTypeRef) String() string {
  return fmt.Sprintf("<typesys.UnionTypeRef Name=%s Location=%s>", self.Name, self.Location)
}

func (self UnionTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self UnionTypeRef) IsTypeRef() bool {
  return true
}
