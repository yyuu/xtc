package ast

import (
  "fmt"
)

type doWhileNode struct {
  Body IStmtNode
  Cond IExprNode
}

func DoWhileNode(body INode, cond INode) doWhileNode {
  return doWhileNode { body.(IStmtNode), cond.(IExprNode) }
}

func (self doWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

type forNode struct {
  Init IExprNode
  Cond IExprNode
  Incr IExprNode
  Body IStmtNode
}

func ForNode(init INode, cond INode, incr INode, body INode) forNode {
  return forNode { init.(IExprNode), cond.(IExprNode), incr.(IExprNode), body.(IExprNode) }
}

func (self forNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.Init, self.Cond, self.Body, self.Incr)
}

type whileNode struct {
  Cond IExprNode
  Body IStmtNode
}

func WhileNode(cond INode, body INode) whileNode {
  return whileNode { cond.(IExprNode), body.(IStmtNode) }
}

func (self whileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}
