package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

func NewNodes(xs...core.INode) []core.INode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.INode { }
  }
}
