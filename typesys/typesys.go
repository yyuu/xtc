package typesys

const (
  TYPE_VOID = 1 << iota
  TYPE_INTEGER
  TYPE_SIGNED
  TYPE_POINTER
  TYPE_ARRAY
  TYPE_STRUCT
  TYPE_UNION
  TYPE_USERTYPE
  TYPE_FUNCTION
)

type ILocation interface {
  GetSourceName() string
  GetLineNumber() int
  GetLineOffset() int
}

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