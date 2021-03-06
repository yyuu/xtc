package typesys

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// FunctionTypeRef
type FunctionTypeRef struct {
  ClassName string
  Location core.Location
  ReturnType core.ITypeRef
  Params *ParamTypeRefs
}

func NewFunctionTypeRef(returnType core.ITypeRef, params core.ITypeRef) *FunctionTypeRef {
  return &FunctionTypeRef { "typesys.FunctionTypeRef", returnType.GetLocation(), returnType, params.(*ParamTypeRefs) }
}

func (self FunctionTypeRef) Key() string {
  return fmt.Sprintf("%s(%s)", self.ReturnType, self.Params)
}

func (self FunctionTypeRef) String() string {
  return self.Key()
}

func (self FunctionTypeRef) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.Key())
  return []byte(s), nil
}

func (self FunctionTypeRef) GetLocation() core.Location {
  return self.Location
}

func (self FunctionTypeRef) IsTypeRef() bool {
  return true
}

func (self FunctionTypeRef) GetReturnType() core.ITypeRef {
  return self.ReturnType
}

func (self FunctionTypeRef) GetParams() *ParamTypeRefs {
  return self.Params
}
