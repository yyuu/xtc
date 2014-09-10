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

type IVariable interface {
  IEntity
  IsVariable() bool
  GetTypeNode() ITypeNode
}

type IDefinedVariable interface {
  IVariable
  IsDefinedVariable() bool
  GetInitializer() IExprNode
  SetInitializer(IExprNode) IDefinedVariable
  HasInitializer() bool
  GetNumRefered() int
//Refered()
}

type IFunction interface {
  IEntity
  IsFunction() bool
  GetTypeNode() ITypeNode
  GetParams() IParams
}

type IDefinedFunction interface {
  IFunction
  IsDefinedFunction() bool
  GetBody() IStmtNode
  SetBody(IStmtNode) IDefinedFunction
  GetScope() IVariableScope
  SetScope(IVariableScope) IDefinedFunction
  ListParameters() []IDefinedVariable
}

type IParams interface {
  IEntity
  GetLocation() Location
  GetParamDescs() []IParameter
}

type IParameter interface {
  IEntity
  GetTypeNode() ITypeNode
}

type IVariableScope interface {
  IsToplevel() bool
  GetToplevel() IVariableScope
  GetParent() IVariableScope
  GetByName(string) *IEntity
}
