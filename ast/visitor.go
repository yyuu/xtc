package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

type IVisitor interface {
  Visit(core.INode)
}

func Visit(v IVisitor, node core.INode) {
  v.Visit(node)
}
