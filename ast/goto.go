package ast

import (
  "fmt"
)

// BreakNode
type breakNode struct {
}

func BreakNode() breakNode {
  return breakNode { }
}

func (self breakNode) String() string {
  return "(break)"
}

func (self breakNode) IsStmt() bool {
  return true
}

// ContinueNode
type continueNode struct {
}

func ContinueNode() continueNode {
  return continueNode { }
}

func (self continueNode) String() string {
  return "(continue)"
}

func (self continueNode) IsStmt() bool {
  return true
}

// ExprStmtNode
type exprStmtNode struct {
  Expr IExprNode
}

func ExprStmtNode(expr IExprNode) exprStmtNode {
  return exprStmtNode { expr }
}

func (self exprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self exprStmtNode) IsStmt() bool {
  return true
}

// GotoNode
type gotoNode struct {
  Target string
}

func GotoNode(target string) gotoNode {
  return gotoNode { target }
}

func (self gotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self gotoNode) IsStmt() bool {
  return true
}

// LabelNode
type labelNode struct {
  Name string
  Stmt IStmtNode
}

func LabelNode(name string, stmt IStmtNode) labelNode {
  return labelNode { name, stmt }
}

func (self labelNode) String() string {
  panic("not implemented")
}

func (self labelNode) IsStmt() bool {
  return true
}

// ReturnNode
type returnNode struct {
  Expr IExprNode
}

func ReturnNode(expr IExprNode) returnNode {
  return returnNode { expr }
}

func (self returnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self returnNode) IsStmt() bool {
  return true
}
