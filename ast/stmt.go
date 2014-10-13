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
  if x == nil {
    return nil
  } else {
    return x.(core.IStmtNode)
  }
}

func AsStmtNodes(xs []core.INode) []core.IStmtNode {
  ys := make([]core.IStmtNode, len(xs))
  for i := range xs {
    if xs[i] == nil {
      ys[i] = nil
    } else {
      ys[i] = xs[i].(core.IStmtNode)
    }
  }
  return ys
}
