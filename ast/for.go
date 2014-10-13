package ast

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
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
  if body == nil { panic("body is nil") }
  return &ForNode { "ast.ForNode", loc, init, cond, incr, body }
}

func (self ForNode) String() string {
  init := "()"
  if self.Init != nil { init = fmt.Sprint(self.Init) }
  cond := "()"
  if self.Cond != nil { cond = fmt.Sprint(self.Cond) }
  incr := "()"
  if self.Incr != nil { incr = fmt.Sprint(self.Incr) }
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", init, cond, self.Body, incr)
}

func (self *ForNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self ForNode) GetLocation() core.Location {
  return self.Location
}

func (self ForNode) HasInit() bool {
  return self.Init != nil
}

func (self *ForNode) GetInit() core.IExprNode {
  return self.Init
}

func (self ForNode) HasCond() bool {
  return self.Cond != nil
}

func (self *ForNode) GetCond() core.IExprNode {
  return self.Cond
}

func (self ForNode) HasIncr() bool {
  return self.Incr != nil
}

func (self *ForNode) GetIncr() core.IExprNode {
  return self.Incr
}

func (self *ForNode) GetBody() core.IStmtNode {
  return self.Body
}

func (self ForNode) GetScope() core.IScope {
  panic("#GetScope called")
}
