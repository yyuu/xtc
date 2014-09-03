package typesys

// ParamTypes
type ParamTypes struct {
  Location ILocation
  ParamDescs []IType
  VarArg bool
}

func NewParamTypes(loc ILocation, paramDescs []IType, vararg bool) ParamTypes {
  return ParamTypes { loc, paramDescs, vararg }
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
  Location ILocation
  ParamDescs []ITypeRef
  VarArg bool
}

func NewParamTypeRefs(loc ILocation, paramDescs []ITypeRef, vararg bool) ParamTypeRefs {
  return ParamTypeRefs { loc, paramDescs, vararg }
}

func (self ParamTypeRefs) GetLocation() ILocation {
  return self.Location
}

func (self ParamTypeRefs) IsTypeRef() bool {
  return true
}
