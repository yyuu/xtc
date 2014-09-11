package core

type IEntity interface {
  String() string
  GetName() string
  IsDefined() bool
  IsPrivate() bool
  IsConstant() bool
  IsRefered() bool
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
}

type IFunction interface {
  IEntity
}

type IVariable interface {
  IEntity
}
