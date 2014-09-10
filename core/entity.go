package core

type IEntity interface {
  String() string
  IsEntity() bool
  GetName() string
  IsDefined() bool
  IsPrivate() bool
  IsConstant() bool
  IsRefered() bool
}

type IDefinedVariable interface {
  IEntity
  IsVariable() bool
  GetTypeNode() ITypeNode
  IsDefinedVariable() bool
  GetInitializer() IExprNode
  SetInitializer(IExprNode) IDefinedVariable
  HasInitializer() bool
  GetNumRefered() int
//Refered()
}
