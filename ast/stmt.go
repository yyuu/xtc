package ast

import (
  "bitbucket.org/yyuu/xtc/core"
)

func NewStmtNodes(xs...core.IStmtNode) []core.IStmtNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IStmtNode { }
  }
}

func AsStmtNode(x core.INode) core.IStmtNode {
  return x.(core.IStmtNode)
}

func AsStmtNodes(xs []core.INode) []core.IStmtNode {
  ys := make([]core.IStmtNode, len(xs))
  for i := range xs {
    ys[i] = xs[i].(core.IStmtNode)
  }
  return ys
}
