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
  t core.IType
}

func NewMemberNode(loc core.Location, expr core.IExprNode, member string) *MemberNode {
  if expr == nil { panic("expr is nil") }
  return &MemberNode { "ast.MemberNode", loc, expr, member, nil }
}

func (self MemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self *MemberNode) AsExprNode() core.IExprNode {
  return self
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

func (self MemberNode) GetType() core.IType {
  if self.t == nil {
    panic(fmt.Errorf("%s type is nil", self.Location))
  }
  return self.t
}

func (self *MemberNode) SetType(t core.IType) {
  self.t = t
}

func (self MemberNode) IsConstant() bool {
  return false
}

func (self MemberNode) IsParameter() bool {
  return false
}

func (self MemberNode) IsLvalue() bool {
  return true
}

func (self MemberNode) IsAssignable() bool {
  return true
}

func (self MemberNode) IsLoadable() bool {
  return false
}

func (self MemberNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self MemberNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
