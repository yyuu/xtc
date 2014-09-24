package core

type IEntity interface {
  String() string
  GetName() string
  IsDefined() bool
  IsPrivate() bool
  IsConstant() bool
  IsParameter() bool
  Refered()
  IsRefered() bool
  GetNumRefered() int
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
  GetType() IType
}

type IFunction interface {
  IEntity
  GetReturnType() IType
}

type IVariable interface {
  IEntity
}

type IScope interface {
  IsToplevel() bool
  GetToplevel() IScope
  GetParent() IScope
  GetByName(string) IEntity
  CheckReferences(*ErrorHandler)
}
