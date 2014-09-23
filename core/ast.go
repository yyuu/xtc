package core

type INode interface {
  String() string
  GetLocation() Location
}

type IExprNode interface {
  INode
  IsExprNode() bool
  GetType() IType
  SetType(IType)
//IsConstant() bool
//IsParameter() bool
//IsLvalue() bool
//IsAssignable() bool
//IsLoadable() bool
//IsCallable() bool
//IsPointer() bool
}

type IStmtNode interface {
  INode
  IsStmtNode() bool
}

type ITypeNode interface {
  INode
  IsTypeNode() bool
  GetTypeRef() ITypeRef
  IsResolved() bool
  GetType() IType
  SetType(IType)
}

type ISlot interface {
  INode
  GetName() string
  GetOffset() int
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
  GetType() IType
}

type ITypeDefinition interface {
  INode
  IsTypeDefinition() bool
  GetName() string
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
  DefiningType() IType
}

type ICompositeTypeDefinition interface {
  ITypeDefinition
  IsCompositeType() bool
  GetMembers() []ISlot
}
