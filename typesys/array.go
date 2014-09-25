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

func (self ArrayType) Key() string {
  if self.Length < 1 {
    return fmt.Sprintf("%s[]", self.BaseType)
  } else {
    return fmt.Sprintf("%s[%d]", self.BaseType, self.Length)
  }
}

func (self ArrayType) String() string {
  return self.Key()
}

func (self ArrayType) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self ArrayType) Size() int {
  return self.PointerSize
}

func (self ArrayType) AllocSize() int {
  return self.BaseType.AllocSize() * self.Length
}

func (self ArrayType) Alignment() int {
  return self.BaseType.AllocSize()
}

func (self ArrayType) IsSameType(other core.IType) bool {
  if !other.IsPointer() && !other.IsArray() {
    return false
  } else {
    return self.BaseType.IsSameType(other.GetBaseType())
  }
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

func (self ArrayType) IsAllocatedArray() bool {
  return self.Length < 1 && (!self.BaseType.IsArray() || self.BaseType.IsAllocatedArray())
}

func (self ArrayType) IsIncompleteArray() bool {
  if self.BaseType.IsArray() {
    return false
  } else {
    return ! self.BaseType.IsAllocatedArray()
  }
}

func (self ArrayType) IsScalar() bool {
  return false
}

func (self ArrayType) IsCallable() bool {
  return false
}

func (self ArrayType) IsCompatible(target core.IType) bool {
  if !target.IsPointer() && !target.IsArray() {
    return false
  } else {
    if target.GetBaseType().IsVoid() {
      return true
    } else {
      return target.GetBaseType().IsCompatible(target.GetBaseType()) && self.BaseType.Size() == target.GetBaseType().Size()
    }
  }
}

func (self ArrayType) IsCastableTo(target core.IType) bool {
  return target.IsPointer() || target.IsArray()
}

func (self ArrayType) GetBaseType() core.IType {
  return self.BaseType
}

func (self ArrayType) GetLength() int {
  return self.Length
}
