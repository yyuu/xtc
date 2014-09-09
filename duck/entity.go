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
  GetInitializer() IExprNode
  SetInitializer(IExprNode) IDefinedVariable
  HasInitializer() bool
  GetNumRefered() int
//Refered()
  IsRefered() bool
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
  SetBody(IStmtNode) IDefinedFunction
  GetScope() IVariableScope
  SetScope(IVariableScope) IDefinedFunction
  ListParameters() []IDefinedVariable
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
  SetValue(IExprNode) IConstant
}

type IVariableScope interface {
  IsToplevel() bool
  GetToplevel() IVariableScope
  GetParent() IVariableScope
  GetByName(string) *IEntity
}

type IConstantTable interface {
  IsConstantTable() bool
}

type IConstantEntry interface {
  IsConstantEntry() bool
}
