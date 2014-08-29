package ast

import (
  "fmt"
)

// BreakNode
type breakNode struct {
  Location ILocation
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
  Location ILocation
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
  Location ILocation
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
  Location ILocation
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
  Location ILocation
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
  Location ILocation
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
