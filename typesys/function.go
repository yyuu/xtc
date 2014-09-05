package typesys

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// FunctionType
type FunctionType struct {
  ReturnType duck.IType
  ParamTypes ParamTypes
}

func NewFunctionType(ret duck.IType, paramTypes ParamTypes) FunctionType {
  return FunctionType { ret, paramTypes }
}

func (self FunctionType) String() string {
  return fmt.Sprintf("<typesys.FunctionType ReturnType=%s ParamTypes=%s>", self.ReturnType, self.ParamTypes)
}

func (self FunctionType) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    ReturnType duck.IType
    ParamTypes ParamTypes
  }
  x.ClassName = "typesys.FunctionType"
  x.ReturnType = self.ReturnType
  x.ParamTypes = self.ParamTypes
  return json.Marshal(x)
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
  Location duck.ILocation
  ReturnType duck.ITypeRef
  Params ParamTypeRefs
}

func NewFunctionTypeRef(returnType duck.ITypeRef, params duck.ITypeRef) FunctionTypeRef {
  return FunctionTypeRef { returnType.GetLocation(), returnType, params.(ParamTypeRefs) }
}

func (self FunctionTypeRef) String() string {
  return fmt.Sprintf("<typesys.FunctionTypeRef Location=%s ReturnType=%s Params=%s>", self.Location, self.ReturnType, self.Params)
}

func (self FunctionTypeRef) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    ReturnType duck.ITypeRef
    Params ParamTypeRefs
  }
  x.ClassName = "typesys.FunctionTypeRef"
  x.Location = self.Location
  x.ReturnType = self.ReturnType
  x.Params = self.Params
  return json.Marshal(x)
}

func (self FunctionTypeRef) GetLocation() duck.ILocation {
  return self.Location
}

func (self FunctionTypeRef) IsTypeRef() bool {
  return true
}
