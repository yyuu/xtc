package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// PtrMemberNode
type PtrMemberNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Member string
  Type core.IType
}

func NewPtrMemberNode(loc core.Location, expr core.IExprNode, member string) *PtrMemberNode {
  if expr == nil { panic("expr is nil") }
  return &PtrMemberNode { "ast.PtrMemberNode", loc, expr, member, nil }
}

func (self PtrMemberNode) String() string {
  return fmt.Sprintf("(slot-ref %s '%s)", self.Expr, self.Member)
}

func (self *PtrMemberNode) AsExprNode() core.IExprNode {
  return self
}

func (self PtrMemberNode) GetLocation() core.Location {
  return self.Location
}

func (self PtrMemberNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self PtrMemberNode) GetMember() string {
  return self.Member
}

func (self PtrMemberNode) GetType() core.IType {
  if self.Type == nil {
    self.Type = self.GetOrigType()
  }
  return self.Type
}

func (self *PtrMemberNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self PtrMemberNode) GetOrigType() core.IType {
  return self.GetDereferedCompositeType().GetMemberType(self.Member)
}

func (self PtrMemberNode) IsConstant() bool {
  return false
}

func (self PtrMemberNode) IsParameter() bool {
  return false
}

func (self PtrMemberNode) IsLvalue() bool {
  return true
}

func (self PtrMemberNode) IsAssignable() bool {
  return true
}

func (self PtrMemberNode) IsLoadable() bool {
  return false
}

func (self PtrMemberNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self PtrMemberNode) IsPointer() bool {
  return self.GetType().IsPointer()
}

func (self PtrMemberNode) GetDereferedCompositeType() core.ICompositeType {
  pt := self.Expr.GetType().(*typesys.PointerType)
  return pt.GetBaseType().(core.ICompositeType)
}

func (self PtrMemberNode) GetDereferedType() core.IType {
  pt := self.Expr.GetType().(*typesys.PointerType)
  return pt.GetBaseType()
}

func (self PtrMemberNode) GetOffset() int {
  return self.GetDereferedCompositeType().GetMemberOffset(self.Member)
}
