package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// DoWhileNode
type DoWhileNode struct {
  ClassName string
  Location duck.Location
  Body duck.IStmtNode
  Cond duck.IExprNode
}

func NewDoWhileNode(loc duck.Location, body duck.IStmtNode, cond duck.IExprNode) DoWhileNode {
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

func (self DoWhileNode) GetLocation() duck.Location {
  return self.Location
}

// ForNode
type ForNode struct {
  ClassName string
  Location duck.Location
  Init duck.IExprNode
  Cond duck.IExprNode
  Incr duck.IExprNode
  Body duck.IStmtNode
}

func NewForNode(loc duck.Location, init duck.IExprNode, cond duck.IExprNode, incr duck.IExprNode, body duck.IStmtNode) ForNode {
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

func (self ForNode) GetLocation() duck.Location {
  return self.Location
}

// WhileNode
type WhileNode struct {
  ClassName string
  Location duck.Location
  Cond duck.IExprNode
  Body duck.IStmtNode
}

func NewWhileNode(loc duck.Location, cond duck.IExprNode, body duck.IStmtNode) WhileNode {
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

func (self WhileNode) GetLocation() duck.Location {
  return self.Location
}
