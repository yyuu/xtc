package ast

import (
  "fmt"
)

// DoWhileNode
type doWhileNode struct {
  Location ILocation
  Body IStmtNode
  Cond IExprNode
}

func DoWhileNode(body IStmtNode, cond IExprNode) doWhileNode {
  return doWhileNode { body, cond }
}

func (self doWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

func (self doWhileNode) IsStmt() bool {
  return true
}

// ForNode
type forNode struct {
  Location ILocation
  Init IExprNode
  Cond IExprNode
  Incr IExprNode
  Body IStmtNode
}

func ForNode(init IExprNode, cond IExprNode, incr IExprNode, body IStmtNode) forNode {
  return forNode { init, cond, incr, body }
}

func (self forNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.Init, self.Cond, self.Body, self.Incr)
}

func (self forNode) IsStmt() bool {
  return true
}

// WhileNode
type whileNode struct {
  Location ILocation
  Cond IExprNode
  Body IStmtNode
}

func WhileNode(cond IExprNode, body IStmtNode) whileNode {
  return whileNode { cond, body }
}

func (self whileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}

func (self whileNode) IsStmt() bool {
  return true
}
