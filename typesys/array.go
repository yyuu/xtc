package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// ArrayType
type ArrayType struct {
  ClassName string
  BaseType core.IType
  Length int
  PointerSize int
}

func NewArrayType(baseType core.IType, length int, pointerSize int) *ArrayType {
  return &ArrayType { "typesys.ArrayType", baseType, length, pointerSize }
}

func (self ArrayType) String() string {
  return fmt.Sprintf("<typesys.ArrayType BaseType=%s Length=%d PointerSize=%d>", self.BaseType, self.Length, self.PointerSize)
}

func (self ArrayType) Size() int {
  return self.PointerSize
}

func (self ArrayType) allocSize() int {
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
