package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// ParamTypes
type ParamTypes struct {
  ClassName string
  Location duck.Location
  ParamDescs []duck.IType
  Vararg bool
}

func NewParamTypes(loc duck.Location, paramDescs []duck.IType, vararg bool) *ParamTypes {
  return &ParamTypes { "typesys.ParamTypes", loc, paramDescs, vararg }
}

func (self ParamTypes) String() string {
  return fmt.Sprintf("<typesys.ParamTypes Location=%s ParamDescs=%s Vararg=%v>", self.Location, self.ParamDescs, self.Vararg)
}

func (self ParamTypes) Size() int {
  panic("ParamTypes#Size called")
}

func (self ParamTypes) AllocSize() int {
  panic("ParamTypes#AllocSize called")
}

func (self ParamTypes) Alignment() int {
  panic("ParamTypes#Alignment called")
}

func (self ParamTypes) IsVoid() bool {
  return false
}

func (self ParamTypes) IsInteger() bool {
  return false
}

func (self ParamTypes) IsSigned() bool {
  return false
}

func (self ParamTypes) IsPointer() bool {
  return false
}

func (self ParamTypes) IsArray() bool {
  return false
}

func (self ParamTypes) IsCompositeType() bool {
  return false
}

func (self ParamTypes) IsStruct() bool {
  return false
}

func (self ParamTypes) IsUnion() bool {
  return false
}

func (self ParamTypes) IsUserType() bool {
  return false
}

func (self ParamTypes) IsFunction() bool {
  return false
}

// ParamTypeRefs
type ParamTypeRefs struct {
  ClassName string
  Location duck.Location
  ParamDescs []duck.ITypeRef
  Vararg bool
}

func NewParamTypeRefs(loc duck.Location, paramDescs []duck.ITypeRef, vararg bool) ParamTypeRefs {
  return ParamTypeRefs { "typesys.ParamTypeRefs", loc, paramDescs, vararg }
}

func (self ParamTypeRefs) String() string {
  return fmt.Sprintf("<typesys.ParamTypeRefs Location=%s ParamDescs=%s Vararg=%v>", self.Location, self.ParamDescs, self.Vararg)
}

func (self ParamTypeRefs) GetLocation() duck.Location {
  return self.Location
}

func (self ParamTypeRefs) IsTypeRef() bool {
  return true
}
