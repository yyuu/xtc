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

func (self PointerType) Key() string {
  return fmt.Sprintf("%s*", self.BaseType)
}

func (self PointerType) String() string {
  return self.Key()
}

func (self PointerType) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
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

func (self PointerType) IsSameType(other core.IType) bool {
  if ! other.IsPointer() {
    return false
  } else {
    return self.BaseType.IsSameType(other.GetBaseType())
  }
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

func (self PointerType) IsAllocatedArray() bool {
  return false
}

func (self PointerType) IsIncompleteArray() bool {
  return false
}

func (self PointerType) IsScalar() bool {
  return true
}

func (self PointerType) IsCallable() bool {
  return false
}

func (self PointerType) IsCompatible(target core.IType) bool {
  if !target.IsPointer() {
    return false
  } else {
    if self.BaseType.IsVoid() {
      return true
    }
    if target.GetBaseType().IsVoid() {
      return true
    }
    return self.BaseType.IsCompatible(target.GetBaseType())
  }
}

func (self PointerType) IsCastableTo(target core.IType) bool {
  return target.IsPointer() || target.IsInteger()
}

func (self PointerType) GetBaseType() core.IType {
  return self.BaseType
}
