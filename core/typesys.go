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
  IsCallable() bool
}

// ITypeRef
type ITypeRef interface {
  String() string

  GetLocation() Location
  IsTypeRef() bool
}

type INamedType interface {
  IType
  GetName() string
}

type ICompositeType interface {
  INamedType
  GetMember(string) ISlot
  GetMembers() []ISlot
//GetMemberType(string) IType
//GetMemberTypes() []IType
  HasMember(string) bool
}
