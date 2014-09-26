package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// FunctionType
type FunctionType struct {
  ClassName string
  ReturnType core.IType
  ParamTypes *ParamTypes
}

func NewFunctionType(ret core.IType, paramTypes *ParamTypes) *FunctionType {
  return &FunctionType { "typesys.FunctionType", ret, paramTypes }
}

func (self FunctionType) Key() string {
  return fmt.Sprintf("%s(%s)", self.ReturnType, self.ParamTypes)
}

func (self FunctionType) String() string {
  return self.Key()
}

func (self FunctionType) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self FunctionType) Size() int {
  panic("FunctionType#Size called")
}

func (self FunctionType) AllocSize() int {
  panic("FunctionType#AllocSize called")
}

func (self FunctionType) Alignment() int {
  panic("FunctionType#Alignment called")
}

func (self FunctionType) IsSameType(other core.IType) bool {
  if !other.IsFunction() {
    return false
  } else {
    t := other.(*FunctionType)
    return t.GetReturnType().IsSameType(self.ReturnType) && t.GetParamTypes().IsSameType(self.ParamTypes)
  }
}

func (self FunctionType) IsVoid() bool {
  return false
}

func (self FunctionType) IsInteger() bool {
  return false
}

func (self FunctionType) IsSigned() bool {
  return false
}

func (self FunctionType) IsPointer() bool {
  return false
}

func (self FunctionType) IsArray() bool {
  return false
}

func (self FunctionType) IsCompositeType() bool {
  return false
}

func (self FunctionType) IsStruct() bool {
  return false
}

func (self FunctionType) IsUnion() bool {
  return false
}

func (self FunctionType) IsUserType() bool {
  return false
}

func (self FunctionType) IsFunction() bool {
  return true
}

func (self FunctionType) IsAllocatedArray() bool {
  return false
}

func (self FunctionType) IsIncompleteArray() bool {
  return false
}

func (self FunctionType) IsScalar() bool {
  return false
}

func (self FunctionType) IsCallable() bool {
  return true
}

func (self FunctionType) IsCompatible(target core.IType) bool {
  if ! target.IsFunction() {
    return false
  } else {
    t := target.(FunctionType)
    return t.GetReturnType().IsCompatible(self.ReturnType) && t.GetParamTypes().IsSameType(self.ParamTypes)
  }
}

func (self FunctionType) IsCastableTo(target core.IType) bool {
  return target.IsFunction()
}

func (self FunctionType) GetReturnType() core.IType {
  if self.ReturnType == nil {
    panic("type is nil")
  }
  return self.ReturnType
}

func (self FunctionType) GetParamTypes() *ParamTypes {
  return self.ParamTypes
}

func (self FunctionType) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}

func (self FunctionType) AcceptsArgc(numArgs int) bool {
  if self.ParamTypes.IsVararg() {
    return numArgs >= self.ParamTypes.MinArgc()
  } else {
    return numArgs == self.ParamTypes.Argc()
  }
}
