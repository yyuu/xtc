package core

// IType
type IType interface {
  Key() string
  String() string

  Size() int
  AllocSize() int
  Alignment() int

  IsSameType(IType) bool

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

  IsAllocatedArray() bool
  IsIncompleteArray() bool
  IsScalar() bool
  IsCallable() bool

  IsCompatible(IType) bool
  IsCastableTo(IType) bool

  GetBaseType() IType
}

// ITypeRef
type ITypeRef interface {
  Key() string
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
  GetMemberType(string) IType
  GetMemberOffset(string) int
  GetMemberTypes() []IType
  HasMember(string) bool
}
