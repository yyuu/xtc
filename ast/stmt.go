package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

func NewStmtNodes(xs...core.IStmtNode) []core.IStmtNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IStmtNode { }
  }
}
