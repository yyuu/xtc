package ast

type INode interface {
  DumpString() string
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
