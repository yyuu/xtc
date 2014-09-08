package duck

type IEntity interface {
  String() string
  IsEntity() bool
}

type IVariable interface {
  IEntity
  IsVariable() bool
  IsPrivate() bool
  GetName() string
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
  GetName() string
}

type IParams interface {
  IEntity
  GetLocation() ILocation
  GetParamDescs() []IParameter
}

type IParameter interface {
  IEntity
  GetTypeNode() ITypeNode
  GetName() string
}

type IConstant interface {
  IEntity
  IsConstant() bool
  GetName() string
  GetTypeNode() ITypeNode
  GetValue() IExprNode
}
