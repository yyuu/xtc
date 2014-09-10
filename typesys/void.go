package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// VoidType
type VoidType struct {
  ClassName string
}

func NewVoidType() *VoidType {
  return &VoidType { "typesys.VoidType" }
}

func (self VoidType) String() string {
  return fmt.Sprintf("<typesys.VoidType>")
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
  ClassName string
  Location core.Location
}

func NewVoidTypeRef(loc core.Location) VoidTypeRef {
  return VoidTypeRef { "typesys.VoidTypeRef", loc }
}

func (self VoidTypeRef) String() string {
  return fmt.Sprintf("<typesys.VoidTypeRef Location=%s>", self.Location)
}

func (self VoidTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self VoidTypeRef) IsTypeRef() bool {
  return true
}
