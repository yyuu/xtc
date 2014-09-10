package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// ForNode
type ForNode struct {
  ClassName string
  Location core.Location
  Init core.IExprNode
  Cond core.IExprNode
  Incr core.IExprNode
  Body core.IStmtNode
}

func NewForNode(loc core.Location, init core.IExprNode, cond core.IExprNode, incr core.IExprNode, body core.IStmtNode) *ForNode {
  if init == nil { panic("init is nil") }
  if cond == nil { panic("cond is nil") }
  if incr == nil { panic("incr is nil") }
  if body == nil { panic("body is nil") }
  return &ForNode { "ast.ForNode", loc, init, cond, incr, body }
}

func (self ForNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.Init, self.Cond, self.Body, self.Incr)
}

func (self ForNode) IsStmtNode() bool {
  return true
}

func (self ForNode) GetLocation() core.Location {
  return self.Location
}
