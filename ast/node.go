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

func AsExprNode(a INode) IExprNode {
  switch b := a.(type) {
    case IExprNode: {
      return b
    }
    default: {
      panic("syntax error")
    }
  }
}

func AsExprNodeList(a []INode) []IExprNode {
  b := make([]IExprNode, len(a))
  for i := range a {
    b[i] = AsExprNode(a[i])
  }
  return b
}
