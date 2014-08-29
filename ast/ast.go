package ast

type INode interface {
  String() string
  Location() ILocation
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
}

type ILocation interface {
  GetSourceName() string
  GetLineNumber() int
  GetLineOffset() int
  String() string
}
