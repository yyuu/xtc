package typesys

import (
  "bitbucket.org/yyuu/bs/core"
)

// VoidType
type VoidType struct {
  ClassName string
}

func NewVoidType() *VoidType {
  return &VoidType { "typesys.VoidType" }
}

func (self VoidType) Key() string {
  return "void"
}

func (self VoidType) String() string {
  return self.Key()
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

func (self VoidType) IsAllocatedArray() bool {
  return false
}

func (self VoidType) IsIncompleteArray() bool {
  return false
}

func (self VoidType) IsScalar() bool {
  return false
}

func (self VoidType) IsCallable() bool {
  return false
}

func (self VoidType) IsCompatible(target core.IType) bool {
  return target.IsVoid()
}

func (self VoidType) IsCastableTo(target core.IType) bool {
  return target.IsVoid()
}

func (self VoidType) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}
