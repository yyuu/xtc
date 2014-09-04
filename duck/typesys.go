package duck

// IType
type IType interface {
  Size() int
  AllocSize() int
  Alignment() int

  IsVoid() bool
  IsInteger() bool
  IsSigned() bool
  IsPointer() bool
  IsArray() bool
  IsCompositeType() bool
  IsStruct() bool
  IsUnion() bool
  IsUserType() bool
  IsFunction() bool
}

// ITypeRef
type ITypeRef interface {
  GetLocation() ILocation
  IsTypeRef() bool
}

type ISlot interface {
  GetName() string
  GetOffset() int
}
