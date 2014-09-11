package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

type INodeVisitor interface {
  VisitNode(core.INode)
}

func VisitNode(v INodeVisitor, node core.INode) {
  v.VisitNode(node)
}
