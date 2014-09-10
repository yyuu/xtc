package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// PtrMemberNode
type PtrMemberNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Member string
}

func NewPtrMemberNode(loc core.Location, expr core.IExprNode, member string) *PtrMemberNode {
  if expr == nil { panic("expr is nil") }
  return &PtrMemberNode { "ast.PtrMemberNode", loc, expr, member }
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self PtrMemberNode) IsExprNode() bool {
  return true
}

func (self PtrMemberNode) GetLocation() core.Location {
  return self.Location
}
