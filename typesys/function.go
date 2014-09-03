package typesys

// FunctionType
type FunctionType struct {
  ReturnType IType
  ParamTypes ParamTypes
}

func NewFunctionType(ret IType, paramTypes ParamTypes) FunctionType {
  return FunctionType { ret, paramTypes }
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
  Location ILocation
  ReturnType ITypeRef
  Params ParamTypeRefs
}

func NewFunctionTypeRef(returnType ITypeRef, params ITypeRef) FunctionTypeRef {
  return FunctionTypeRef { returnType.GetLocation(), returnType, params.(ParamTypeRefs) }
}

func (self FunctionTypeRef) GetLocation() ILocation {
  return self.Location
}

func (self FunctionTypeRef) IsTypeRef() bool {
  return true
}
