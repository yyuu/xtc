package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// MemberNode
type MemberNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Member string
}

func NewMemberNode(loc core.Location, expr core.IExprNode, member string) *MemberNode {
  if expr == nil { panic("expr is nil") }
  return &MemberNode { "ast.MemberNode", loc, expr, member }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self MemberNode) IsExprNode() bool {
  return true
}

func (self MemberNode) GetLocation() core.Location {
  return self.Location
}

func (self MemberNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self MemberNode) GetMember() string {
  return self.Member
}
