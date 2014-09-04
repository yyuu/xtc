package duck

type INode interface {
  String() string
  MarshalJSON() ([]byte, error)
  GetLocation() ILocation
}

type IExprNode interface {
  INode
  IsExpr() bool
}

type IStmtNode interface {
  INode
  IsStmt() bool
}

type ITypeNode interface {
  INode
  IsType() bool
  GetTypeRef() ITypeRef
}

type ITypeDefinition interface {
  INode
  IsTypeDefinition() bool
}

type ILocation interface {
  GetSourceName() string
  GetLineNumber() int
  GetLineOffset() int
}
