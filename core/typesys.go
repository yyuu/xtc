package core

// IType
type IType interface {
  String() string

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

  GetLocation() Location
  IsTypeRef() bool
}

type ISlot interface {
  String() string

  GetName() string
  GetOffset() int
}

type ITypeTable interface {
  String() string

  IsTypeTable() bool
  IsDefined(ITypeRef) bool
  GetType(ITypeRef) IType
  PutType(ITypeRef, IType)
}
