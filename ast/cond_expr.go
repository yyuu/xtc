package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// CondExprNode
type CondExprNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  ThenExpr core.IExprNode
  ElseExpr core.IExprNode
  t core.IType
}

func NewCondExprNode(loc core.Location, cond core.IExprNode, thenExpr core.IExprNode, elseExpr core.IExprNode) *CondExprNode {
  if cond == nil { panic("cond is nil") }
  if thenExpr == nil { panic("thenExpr is nil") }
  if elseExpr == nil { panic("elseExpr is nil") }
  return &CondExprNode { "ast.CondExprNode", loc, cond, thenExpr, elseExpr, nil }
}

func (self CondExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

func (self CondExprNode) IsExprNode() bool {
  return true
}

func (self CondExprNode) GetLocation() core.Location {
  return self.Location
}

func (self CondExprNode) GetCond() core.IExprNode {
  return self.Cond
}

func (self CondExprNode) GetThenExpr() core.IExprNode {
  return self.ThenExpr
}

func (self CondExprNode) GetElseExpr() core.IExprNode {
  return self.ElseExpr
}

func (self CondExprNode) GetType() core.IType {
  if self.t == nil {
    panic("type is nil")
  }
  return self.t
}

func (self *CondExprNode) SetType(t core.IType) {
  self.t = t
}
