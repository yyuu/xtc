package ast

type INode interface {
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
