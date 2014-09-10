package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// DoWhileNode
type DoWhileNode struct {
  ClassName string
  Location core.Location
  Body core.IStmtNode
  Cond core.IExprNode
}

func NewDoWhileNode(loc core.Location, body core.IStmtNode, cond core.IExprNode) DoWhileNode {
  if body == nil { panic("body is nil") }
  if cond == nil { panic("cond is nil") }
  return DoWhileNode { "ast.DoWhileNode", loc, body, cond }
}

func (self DoWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

func (self DoWhileNode) IsStmtNode() bool {
  return true
}

func (self DoWhileNode) GetLocation() core.Location {
  return self.Location
}

// ForNode
type ForNode struct {
  ClassName string
  Location core.Location
  Init core.IExprNode
  Cond core.IExprNode
  Incr core.IExprNode
  Body core.IStmtNode
}

func NewForNode(loc core.Location, init core.IExprNode, cond core.IExprNode, incr core.IExprNode, body core.IStmtNode) ForNode {
  if init == nil { panic("init is nil") }
  if cond == nil { panic("cond is nil") }
  if incr == nil { panic("incr is nil") }
  if body == nil { panic("body is nil") }
  return ForNode { "ast.ForNode", loc, init, cond, incr, body }
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

// WhileNode
type WhileNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  Body core.IStmtNode
}

func NewWhileNode(loc core.Location, cond core.IExprNode, body core.IStmtNode) WhileNode {
  if cond == nil { panic("cond is nil") }
  if body == nil { panic("body is nil") }
  return WhileNode { "ast.WhileNode", loc, cond, body }
}

func (self WhileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}

func (self WhileNode) IsStmtNode() bool {
  return true
}

func (self WhileNode) GetLocation() core.Location {
  return self.Location
}
