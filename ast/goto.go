package ast

import (
  "fmt"
)

// BreakNode
type breakNode struct {
  location Location
}

func NewBreakNode(location Location) breakNode {
  return breakNode { location }
}

func (self breakNode) String() string {
  return "(break)"
}

func (self breakNode) IsStmt() bool {
  return true
}

func (self breakNode) GetLocation() Location {
  return self.location
}

// ContinueNode
type continueNode struct {
  location Location
}

func NewContinueNode(location Location) continueNode {
  return continueNode { location }
}

func (self continueNode) String() string {
  return "(continue)"
}

func (self continueNode) IsStmt() bool {
  return true
}

func (self continueNode) GetLocation() Location {
  return self.location
}

// ExprStmtNode
type exprStmtNode struct {
  location Location
  Expr IExprNode
}

func NewExprStmtNode(location Location, expr IExprNode) exprStmtNode {
  return exprStmtNode { location, expr }
}

func (self exprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self exprStmtNode) IsStmt() bool {
  return true
}

func (self exprStmtNode) GetLocation() Location {
  return self.location
}

// GotoNode
type gotoNode struct {
  location Location
  Target string
}

func NewGotoNode(location Location, target string) gotoNode {
  return gotoNode { location, target }
}

func (self gotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self gotoNode) IsStmt() bool {
  return true
}

func (self gotoNode) GetLocation() Location {
  return self.location
}

// LabelNode
type labelNode struct {
  location Location
  Name string
  Stmt IStmtNode
}

func NewLabelNode(location Location, name string, stmt IStmtNode) labelNode {
  return labelNode { location, name, stmt }
}

func (self labelNode) String() string {
  panic("not implemented")
}

func (self labelNode) IsStmt() bool {
  return true
}

func (self labelNode) GetLocation() Location {
  return self.location
}

// ReturnNode
type returnNode struct {
  location Location
  Expr IExprNode
}

func NewReturnNode(location Location, expr IExprNode) returnNode {
  return returnNode { location, expr }
}

func (self returnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self returnNode) IsStmt() bool {
  return true
}

func (self returnNode) GetLocation() Location {
  return self.location
}
