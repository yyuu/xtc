package ast

import (
  "bitbucket.org/yyuu/xtc/core"
)

func NewExprNodes(xs...core.IExprNode) []core.IExprNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IExprNode { }
  }
}

func AsExprNode(x core.INode) core.IExprNode {
  if x == nil {
    return nil
  } else {
    return x.(core.IExprNode)
  }
}

func AsExprNodes(xs []core.INode) []core.IExprNode {
  ys := make([]core.IExprNode, len(xs))
  for i := range xs {
    if xs[i] == nil {
      ys[i] = nil
    } else {
      ys[i] = xs[i].(core.IExprNode)
    }
  }
  return ys
}
