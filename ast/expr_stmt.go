package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// ExprStmtNode
type ExprStmtNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewExprStmtNode(loc core.Location, expr core.IExprNode) *ExprStmtNode {
  if expr == nil { panic("expr is nil") }
  return &ExprStmtNode { "ast.ExprStmtNode", loc, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self *ExprStmtNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self ExprStmtNode) GetLocation() core.Location {
  return self.Location
}

func (self *ExprStmtNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self ExprStmtNode) GetScope() core.IScope {
  panic("#GetScope called")
}
