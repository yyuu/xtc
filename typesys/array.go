package typesys

// ArrayType
type ArrayType struct {
  BaseType IType
  Length int
  PointerSize int
}

func NewArrayType(baseType IType, length int, pointerSize int) ArrayType {
  return ArrayType { baseType, length, pointerSize }
}

func (self ArrayType) Size() int {
  return self.PointerSize
}

func (self ArrayType) AllocSize() int {
  return self.BaseType.AllocSize() * self.Length
}

func (self ArrayType) Alignment() int {
  return self.BaseType.AllocSize()
}

func (self ArrayType) IsVoid() bool {
  return false
}

func (self ArrayType) IsInteger() bool {
  return false
}

func (self ArrayType) IsSigned() bool {
  return false
}

func (self ArrayType) IsPointer() bool {
  return false
}

func (self ArrayType) IsArray() bool {
  return true
}

func (self ArrayType) IsCompositeType() bool {
  return false
}

func (self ArrayType) IsStruct() bool {
  return false
}

func (self ArrayType) IsUnion() bool {
  return false
}

func (self ArrayType) IsUserType() bool {
  return false
}

func (self ArrayType) IsFunction() bool {
  return false
}

// ArrayTypeRef
type ArrayTypeRef struct {
  Location ILocation
  BaseType ITypeRef
  Length int
}

func NewArrayTypeRef(baseType ITypeRef, length int) ArrayTypeRef {
  return ArrayTypeRef { baseType.GetLocation(), baseType, length }
}

func (self ArrayTypeRef) GetLocation() ILocation {
  return self.Location
}

func (self ArrayTypeRef) IsTypeRef() bool {
  return true
}
