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
  t core.IType
}

func NewCastNode(loc core.Location, t core.ITypeNode, expr core.IExprNode) *CastNode {
  if t == nil { panic("t is nil") }
  if expr == nil { panic("expr is nil") }
  return &CastNode { "ast.CastNode", loc, t, expr, nil }
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

func (self CastNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self CastNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *CastNode) SetType(t core.IType) {
  self.t = t
}

func (self CastNode) IsConstant() bool {
  return false
}

func (self CastNode) IsParameter() bool {
  return false
}

func (self CastNode) IsLvalue() bool {
  return self.Expr.IsLvalue()
}

func (self CastNode) IsAssignable() bool {
  return true
}

func (self CastNode) IsLoadable() bool {
  return false
}

func (self CastNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self CastNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
