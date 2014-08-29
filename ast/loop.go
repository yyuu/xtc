package ast

import (
  "fmt"
)

// DoWhileNode
type doWhileNode struct {
  location ILocation
  Body IStmtNode
  Cond IExprNode
}

func DoWhileNode(location ILocation, body IStmtNode, cond IExprNode) doWhileNode {
  return doWhileNode { location, body, cond }
}

func (self doWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

func (self doWhileNode) IsStmt() bool {
  return true
}

func (self doWhileNode) GetLocation() ILocation {
  return self.location
}

// ForNode
type forNode struct {
  location ILocation
  Init IExprNode
  Cond IExprNode
  Incr IExprNode
  Body IStmtNode
}

func ForNode(location ILocation, init IExprNode, cond IExprNode, incr IExprNode, body IStmtNode) forNode {
  return forNode { location, init, cond, incr, body }
}

func (self forNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.Init, self.Cond, self.Body, self.Incr)
}

func (self forNode) IsStmt() bool {
  return true
}

func (self forNode) GetLocation() ILocation {
  return self.location
}

// WhileNode
type whileNode struct {
  location ILocation
  Cond IExprNode
  Body IStmtNode
}

func WhileNode(location ILocation, cond IExprNode, body IStmtNode) whileNode {
  return whileNode { location, cond, body }
}

func (self whileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}

func (self whileNode) IsStmt() bool {
  return true
}

func (self whileNode) GetLocation() ILocation {
  return self.location
}
