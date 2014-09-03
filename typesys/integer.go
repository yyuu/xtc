package typesys

// IntegerType
type IntegerType struct {
  IntegerSize int
  Signed bool
  Name string
}

func NewIntegerType(size int, isSigned bool, name string) IntegerType {
  return IntegerType { size, isSigned, name }
}

func (self IntegerType) Size() int {
  return self.IntegerSize
}

func (self IntegerType) AllocSize() int {
  return self.Size()
}

func (self IntegerType) Alignment() int {
  return self.AllocSize()
}

func (self IntegerType) IsVoid() bool {
  return false
}

func (self IntegerType) IsInteger() bool {
  return true
}

func (self IntegerType) IsSigned() bool {
  return self.Signed
}

func (self IntegerType) IsPointer() bool {
  return false
}

func (self IntegerType) IsArray() bool {
  return false
}

func (self IntegerType) IsCompositeType() bool {
  return false
}

func (self IntegerType) IsStruct() bool {
  return false
}

func (self IntegerType) IsUnion() bool {
  return false
}

func (self IntegerType) IsUserType() bool {
  return false
}

func (self IntegerType) IsFunction() bool {
  return false
}

// IntegerTypeRef
type IntegerTypeRef struct {
  Location ILocation
  Name string
}

func NewIntegerTypeRef(location ILocation, name string) IntegerTypeRef {
  return IntegerTypeRef { location, name }
}

func (self IntegerTypeRef) GetLocation() ILocation {
  return self.Location
}

func (self IntegerTypeRef) IsTypeRef() bool {
  return true
}
