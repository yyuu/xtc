package ast

import (
  "fmt"
)

// DoWhileNode
type DoWhileNode struct {
  location Location
  Body IStmtNode
  Cond IExprNode
}

func NewDoWhileNode(location Location, body IStmtNode, cond IExprNode) DoWhileNode {
  return DoWhileNode { location, body, cond }
}

func (self DoWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

func (self DoWhileNode) IsStmt() bool {
  return true
}

func (self DoWhileNode) GetLocation() Location {
  return self.location
}

// ForNode
type ForNode struct {
  location Location
  Init IExprNode
  Cond IExprNode
  Incr IExprNode
  Body IStmtNode
}

func NewForNode(location Location, init IExprNode, cond IExprNode, incr IExprNode, body IStmtNode) ForNode {
  return ForNode { location, init, cond, incr, body }
}

func (self ForNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.Init, self.Cond, self.Body, self.Incr)
}

func (self ForNode) IsStmt() bool {
  return true
}

func (self ForNode) GetLocation() Location {
  return self.location
}

// WhileNode
type WhileNode struct {
  location Location
  Cond IExprNode
  Body IStmtNode
}

func NewWhileNode(location Location, cond IExprNode, body IStmtNode) WhileNode {
  return WhileNode { location, cond, body }
}

func (self WhileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}

func (self WhileNode) IsStmt() bool {
  return true
}

func (self WhileNode) GetLocation() Location {
  return self.location
}
