package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// ArefNode
type ArefNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Index core.IExprNode
  Type core.IType
}

func NewArefNode(loc core.Location, expr core.IExprNode, index core.IExprNode) *ArefNode {
  if expr == nil { panic("expr is nil") }
  if index == nil { panic("index is nil") }
  return &ArefNode { "ast.ArefNode", loc, expr, index, nil }
}

func (self ArefNode) String() string {
  return fmt.Sprintf("(vector-ref %s %s)", self.Expr, self.Index)
}

func (self *ArefNode) AsExprNode() core.IExprNode {
  return self
}

func (self ArefNode) GetLocation() core.Location {
  return self.Location
}

func (self ArefNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self ArefNode) GetIndex() core.IExprNode {
  return self.Index
}

func (self ArefNode) GetType() core.IType {
  if self.Type == nil {
    self.Type = self.GetOrigType()
  }
  return self.Type
}

func (self *ArefNode) SetType(t core.IType) {
  if self.Type != nil {
    panic("#SetType called twice")
  }
  self.Type = t
}

func (self ArefNode) GetOrigType() core.IType {
  return self.Expr.GetOrigType().GetBaseType()
}

func (self ArefNode) IsConstant() bool {
  return false
}

func (self ArefNode) IsParameter() bool {
  return false
}

func (self ArefNode) IsLvalue() bool {
  return true
}

func (self ArefNode) IsAssignable() bool {
  return true
}

func (self ArefNode) IsLoadable() bool {
  return false
}

func (self ArefNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self ArefNode) IsPointer() bool {
  return self.GetType().IsPointer()
}

func (self ArefNode) IsMultiDimension() bool {
  _, ok := self.Expr.(*ArefNode)
  return ok
}

func (self ArefNode) GetBaseExpr() core.IExprNode {
  if self.IsMultiDimension() {
    return self.Expr.(*ArefNode).GetBaseExpr()
  } else {
    return self.Expr
  }
}

func (self ArefNode) GetElementSize() int {
  return self.GetOrigType().AllocSize()
}

func (self ArefNode) GetLength() int {
  t := self.Expr.GetOrigType().(*typesys.ArrayType)
  return t.GetLength()
}
