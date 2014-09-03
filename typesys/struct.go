package typesys

// StructType
type StructType struct {
  Location ILocation
  Name string
  Members []ISlot
}

func NewStructType(name string, membs []ISlot, loc ILocation) StructType {
  return StructType { loc, name, membs }
}

func (self StructType) Size() int {
  panic("StructType#Size called")
}

func (self StructType) AllocSize() int {
  panic("StructType#AllocSize called")
}

func (self StructType) Alignment() int {
  panic("StructType#Alignment called")
}

func (self StructType) IsVoid() bool {
  return false
}

func (self StructType) IsInteger() bool {
  return false
}

func (self StructType) IsSigned() bool {
  return false
}

func (self StructType) IsPointer() bool {
  return false
}

func (self StructType) IsArray() bool {
  return false
}

func (self StructType) IsCompositeType() bool {
  return false
}

func (self StructType) IsStruct() bool {
  return true
}

func (self StructType) IsUnion() bool {
  return false
}

func (self StructType) IsUserType() bool {
  return false
}

func (self StructType) IsFunction() bool {
  return false
}

// StructTypeRef
type StructTypeRef struct {
  Location ILocation
  Name string
}

func NewStructTypeRef(loc ILocation, name string) StructTypeRef {
  return StructTypeRef { loc, name }
}

func (self StructTypeRef) GetLocation() ILocation {
  return self.Location
}

func (self StructTypeRef) IsTypeRef() bool {
  return true
}
