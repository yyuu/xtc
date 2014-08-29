package ast

type INode interface {
  String() string
}

type IExprNode interface {
  INode
}

type IStmtNode interface {
  INode
}

type ITypeNode interface {
  INode
}
