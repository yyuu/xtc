package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// ReturnNode
type ReturnNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewReturnNode(loc core.Location, expr core.IExprNode) *ReturnNode {
  if expr == nil { panic("expr is nil") }
  return &ReturnNode { "ast.ReturnNode", loc, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self *ReturnNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self ReturnNode) GetLocation() core.Location {
  return self.Location
}

func (self ReturnNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self *ReturnNode) SetExpr(expr core.IExprNode) {
  self.Expr = expr
}
