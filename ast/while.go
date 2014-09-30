package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// WhileNode
type WhileNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  Body core.IStmtNode
}

func NewWhileNode(loc core.Location, cond core.IExprNode, body core.IStmtNode) *WhileNode {
  if cond == nil { panic("cond is nil") }
  if body == nil { panic("body is nil") }
  return &WhileNode { "ast.WhileNode", loc, cond, body }
}

func (self WhileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}

func (self *WhileNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self WhileNode) GetLocation() core.Location {
  return self.Location
}

func (self *WhileNode) GetCond() core.IExprNode {
  return self.Cond
}

func (self *WhileNode) GetBody() core.IStmtNode {
  return self.Body
}

func (self WhileNode) GetScope() core.IScope {
  panic("#GetScope called")
}
