package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// CastNode
type CastNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Expr core.IExprNode
}

func NewCastNode(loc core.Location, t core.ITypeNode, expr core.IExprNode) CastNode {
  if t == nil { panic("t is nil") }
  if expr == nil { panic("expr is nil") }
  return CastNode { "ast.CastNode", loc, t, expr }
}

func (self CastNode) String() string {
  return fmt.Sprintf("(%s %s)", self.TypeNode, self.Expr)
}

func (self CastNode) IsExprNode() bool {
  return true
}

func (self CastNode) GetLocation() core.Location {
  return self.Location
}

func (self CastNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self CastNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}
