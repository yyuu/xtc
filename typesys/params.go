package typesys

import (
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

// ParamTypes
type ParamTypes struct {
  ClassName string
  Location core.Location
  ParamDescs []core.IType
  Vararg bool
}

func NewParamTypes(loc core.Location, paramDescs []core.IType, vararg bool) *ParamTypes {
  return &ParamTypes { "typesys.ParamTypes", loc, paramDescs, vararg }
}

func (self ParamTypes) String() string {
  params := make([]string, len(self.ParamDescs))
  for i := range self.ParamDescs {
    params[i] = self.ParamDescs[i].String()
  }
  return strings.Join(params, ",")
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

func (self ParamTypes) IsCallable() bool {
  return false
}

func (self ParamTypes) GetParamDescs() []core.IType {
  return self.ParamDescs
}

func (self ParamTypes) IsVararg() bool {
  return self.Vararg
}
