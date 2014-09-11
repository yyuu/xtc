package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// PointerType
type PointerType struct {
  ClassName string
  PointerSize int
  BaseType core.IType
}

func NewPointerType(size int, baseType core.IType) *PointerType {
  return &PointerType { "typesys.PointerType", size, baseType }
}

func (self PointerType) String() string {
  return fmt.Sprintf("%s*", self.BaseType)
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
