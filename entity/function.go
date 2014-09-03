package entity

type DefinedFunction struct {
  Private bool
  TypeNode ITypeNode
  Name string
  Params Params
  Body IStmtNode
}

func NewDefinedFunction(priv bool, t ITypeNode, name string, params Params, body IStmtNode) DefinedFunction {
  return DefinedFunction {
    Private: priv,
    TypeNode: t,
    Name: name,
    Params: params,
    Body: body,
  }
}

func (self DefinedFunction) IsEntity() bool {
  return true
}

type UndefinedFunction struct {
  TypeNode ITypeNode
  Name string
  Params Params
}

func NewUndefinedFunction(t ITypeNode, name string, params Params) UndefinedFunction {
  return UndefinedFunction {
    TypeNode: t,
    Name: name,
    Params: params,
  }
}

func (self UndefinedFunction) IsEntity() bool {
  return true
}
