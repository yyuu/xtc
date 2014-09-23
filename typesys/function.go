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

func (self FunctionType) String() string {
  return fmt.Sprintf("%s(%s)", self.ReturnType, self.ParamTypes)
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

func (self FunctionType) IsCallable() bool {
  return true
}

func (self FunctionType) GetReturnType() core.IType {
  return self.ReturnType
}

func (self FunctionType) GetParamTypes() *ParamTypes {
  return self.ParamTypes
}

func (self FunctionType) GetBaseType() core.IType {
  panic("#baseType called for undereferable type")
}
