package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

func NewExprNodes(xs...core.IExprNode) []core.IExprNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IExprNode { }
  }
}
