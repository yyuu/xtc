package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// FunctionType
type FunctionType struct {
  ClassName string
  ReturnType duck.IType
  ParamTypes ParamTypes
}

func NewFunctionType(ret duck.IType, paramTypes ParamTypes) FunctionType {
  return FunctionType { "typesys.FunctionType", ret, paramTypes }
}

func (self FunctionType) String() string {
  return fmt.Sprintf("<typesys.FunctionType ReturnType=%s ParamTypes=%s>", self.ReturnType, self.ParamTypes)
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

// FunctionTypeRef
type FunctionTypeRef struct {
  ClassName string
  Location duck.ILocation
  ReturnType duck.ITypeRef
  Params ParamTypeRefs
}

func NewFunctionTypeRef(returnType duck.ITypeRef, params duck.ITypeRef) FunctionTypeRef {
  return FunctionTypeRef { "typesys.FunctionTypeRef", returnType.GetLocation(), returnType, params.(ParamTypeRefs) }
}

func (self FunctionTypeRef) String() string {
  return fmt.Sprintf("<typesys.FunctionTypeRef Location=%s ReturnType=%s Params=%s>", self.Location, self.ReturnType, self.Params)
}

func (self FunctionTypeRef) GetLocation() duck.ILocation {
  return self.Location
}

func (self FunctionTypeRef) IsTypeRef() bool {
  return true
}
