package typesys

import (
  "bitbucket.org/yyuu/bs/core"
)

// UserType
type UserType struct {
  ClassName string
  Location core.Location
  Name string
  Real core.ITypeNode
}

func NewUserType(name string, real core.ITypeNode, loc core.Location) *UserType {
  return &UserType { "typesys.UserType", loc, name, real }
}

func (self UserType) Key() string {
  return self.Name
}

func (self UserType) String() string {
  return self.Key()
}

func (self UserType) Size() int {
  panic("UserType#Size called")
}

func (self UserType) AllocSize() int {
  panic("UserType#AllocSize called")
}

func (self UserType) Alignment() int {
  panic("UserType#Alignment called")
}

func (self UserType) IsVoid() bool {
  return false
}

func (self UserType) IsInteger() bool {
  return false
}

func (self UserType) IsSigned() bool {
  return false
}

func (self UserType) IsPointer() bool {
  return false
}

func (self UserType) IsArray() bool {
  return false
}

func (self UserType) IsCompositeType() bool {
  return false
}

func (self UserType) IsStruct() bool {
  return true
}

func (self UserType) IsUnion() bool {
  return false
}

func (self UserType) IsUserType() bool {
  return false
}

func (self UserType) IsFunction() bool {
  return false
}

func (self UserType) IsCallable() bool {
  return false
}

func (self UserType) GetName() string {
  return self.Name
}

func (self UserType) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}
