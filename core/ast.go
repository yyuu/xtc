package core

type INode interface {
  String() string
  GetLocation() Location
}

type IExprNode interface {
  INode
  IsExprNode() bool
}

type IStmtNode interface {
  INode
  IsStmtNode() bool
}

type ITypeNode interface {
  INode
  IsTypeNode() bool
  GetTypeRef() ITypeRef
}

type ITypeDefinition interface {
  INode
  IsTypeDefinition() bool
  GetTypeRef() ITypeRef
  DefiningType() IType
}

type IStructNode interface {
  ITypeDefinition
}

type IUnionNode interface {
  ITypeDefinition
}

type ITypedefNode interface {
  ITypeDefinition
}
