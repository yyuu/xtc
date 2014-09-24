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

func (self ParamTypes) Key() string {
  params := make([]string, len(self.ParamDescs))
  for i := range self.ParamDescs {
    params[i] = self.ParamDescs[i].String()
  }
  return strings.Join(params, ",")
}

func (self ParamTypes) String() string {
  return self.Key()
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

func (self ParamTypes) IsSameType(other core.IType) bool {
  t, ok := other.(ParamTypes)
  if ! ok {
    return false
  }
  if self.Vararg != t.IsVararg() {
    return false
  }
  ps := t.GetParamDescs()
  if len(self.ParamDescs) != len(ps) {
    return false
  }
  for i := range self.ParamDescs {
    if ! self.ParamDescs[i].IsSameType(ps[i]) {
      return false
    }
  }
  return true
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

func (self ParamTypes) IsAllocatedArray() bool {
  return false
}

func (self ParamTypes) IsIncompleteArray() bool {
  return false
}

func (self ParamTypes) IsScalar() bool {
  return false
}

func (self ParamTypes) IsCallable() bool {
  return false
}

func (self ParamTypes) IsCompatible(target core.IType) bool {
  return false
}

func (self ParamTypes) IsCastableTo(target core.IType) bool {
  return false
}

func (self ParamTypes) GetParamDescs() []core.IType {
  return self.ParamDescs
}

func (self ParamTypes) Argc() int {
  if self.Vararg {
    panic("must not happen")
  }
  return len(self.ParamDescs)
}

func (self ParamTypes) MinArgc() int {
  return len(self.ParamDescs)
}

func (self ParamTypes) IsVararg() bool {
  return self.Vararg
}

func (self ParamTypes) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}
