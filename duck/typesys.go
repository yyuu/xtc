package duck

// IType
type IType interface {
  String() string
  MarshalJSON() ([]byte, error)

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
  String() string
  MarshalJSON() ([]byte, error)

  GetLocation() ILocation
  IsTypeRef() bool
}

type ISlot interface {
  String() string
  MarshalJSON() ([]byte, error)

  GetName() string
  GetOffset() int
}

