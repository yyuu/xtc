package duck

type IEntity interface {
  String() string
  IsEntity() bool
  GetName() string
  IsDefined() bool
}

type IVariable interface {
  IEntity
  IsVariable() bool
  IsPrivate() bool
  GetTypeNode() ITypeNode
}

type IDefinedVariable interface {
  IVariable
  IsDefinedVariable() bool
  GetNumRefered() int
  GetInitializer() IExprNode
}

type IUndefinedVariable interface {
  IVariable
  IsUndefinedVariable() bool
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
  IsPrivate() bool
  GetBody() IStmtNode
}

type IUndefinedFunction interface {
  IFunction
  IsUndefinedFunction() bool
}

type IParams interface {
  IEntity
  GetLocation() ILocation
  GetParamDescs() []IParameter
}

type IParameter interface {
  IEntity
  GetTypeNode() ITypeNode
}

type IConstant interface {
  IEntity
  IsConstant() bool
  GetTypeNode() ITypeNode
  GetValue() IExprNode
}

type IScope interface {
  IsToplevel() bool
  GetToplevel() IScope
  GetParent() IScope
}
