package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

// ReturnNode
type ReturnNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewReturnNode(loc core.Location, expr core.IExprNode) *ReturnNode {
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

func (self ReturnNode) HasExpr() bool {
  return self.Expr != nil
}

func (self *ReturnNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self *ReturnNode) SetExpr(expr core.IExprNode) {
  self.Expr = expr
}

func (self ReturnNode) GetScope() core.IScope {
  panic("#GetScope called")
}
