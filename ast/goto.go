package ast

import (
  "fmt"
)

// BreakNode
type BreakNode struct {
  location Location
}

func NewBreakNode(location Location) BreakNode {
  return BreakNode { location }
}

func (self BreakNode) String() string {
  return "(break)"
}

func (self BreakNode) IsStmt() bool {
  return true
}

func (self BreakNode) GetLocation() Location {
  return self.location
}

// ContinueNode
type ContinueNode struct {
  location Location
}

func NewContinueNode(location Location) ContinueNode {
  return ContinueNode { location }
}

func (self ContinueNode) String() string {
  return "(continue)"
}

func (self ContinueNode) IsStmt() bool {
  return true
}

func (self ContinueNode) GetLocation() Location {
  return self.location
}

// ExprStmtNode
type ExprStmtNode struct {
  location Location
  Expr IExprNode
}

func NewExprStmtNode(location Location, expr IExprNode) ExprStmtNode {
  return ExprStmtNode { location, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ExprStmtNode) IsStmt() bool {
  return true
}

func (self ExprStmtNode) GetLocation() Location {
  return self.location
}

// GotoNode
type GotoNode struct {
  location Location
  Target string
}

func NewGotoNode(location Location, target string) GotoNode {
  return GotoNode { location, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self GotoNode) IsStmt() bool {
  return true
}

func (self GotoNode) GetLocation() Location {
  return self.location
}

// LabelNode
type LabelNode struct {
  location Location
  Name string
  Stmt IStmtNode
}

func NewLabelNode(location Location, name string, stmt IStmtNode) LabelNode {
  return LabelNode { location, name, stmt }
}

func (self LabelNode) String() string {
  panic("not implemented")
}

func (self LabelNode) IsStmt() bool {
  return true
}

func (self LabelNode) GetLocation() Location {
  return self.location
}

// ReturnNode
type ReturnNode struct {
  location Location
  Expr IExprNode
}

func NewReturnNode(location Location, expr IExprNode) ReturnNode {
  return ReturnNode { location, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ReturnNode) IsStmt() bool {
  return true
}

func (self ReturnNode) GetLocation() Location {
  return self.location
}
