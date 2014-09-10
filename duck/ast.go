package duck

type INode interface {
  String() string
  GetLocation() ILocation
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

type ILocation interface {
  String() string
  GetSourceName() string
  GetLineNumber() int
  GetLineOffset() int
}
