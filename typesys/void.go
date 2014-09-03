package typesys

// VoidType
type VoidType struct {
}

func NewVoidType() VoidType {
  return VoidType { }
}

func (self VoidType) Size() int {
  return 1
}

func (self VoidType) AllocSize() int {
  return self.Size()
}

func (self VoidType) Alignment() int {
  return self.AllocSize()
}

func (self VoidType) IsVoid() bool {
  return true
}

func (self VoidType) IsInteger() bool {
  return false
}

func (self VoidType) IsSigned() bool {
  return false
}

func (self VoidType) IsPointer() bool {
  return false
}

func (self VoidType) IsArray() bool {
  return false
}

func (self VoidType) IsCompositeType() bool {
  return false
}

func (self VoidType) IsStruct() bool {
  return false
}

func (self VoidType) IsUnion() bool {
  return false
}

func (self VoidType) IsUserType() bool {
  return false
}

func (self VoidType) IsFunction() bool {
  return false
}

// VoidTypeRef
type VoidTypeRef struct {
  Location ILocation
}

func NewVoidTypeRef(location ILocation) VoidTypeRef {
  return VoidTypeRef { location }
}

func (self VoidTypeRef) GetLocation() ILocation {
  return self.Location
}

func (self VoidTypeRef) IsTypeRef() bool {
  return true
}