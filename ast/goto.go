package ast

import (
  "fmt"
)

// BreakNode
type breakNode struct {
  location ILocation
}

func BreakNode(location ILocation) breakNode {
  return breakNode { location }
}

func (self breakNode) String() string {
  return "(break)"
}

func (self breakNode) IsStmt() bool {
  return true
}

func (self breakNode) GetLocation() ILocation {
  return self.location
}

// ContinueNode
type continueNode struct {
  location ILocation
}

func ContinueNode(location ILocation) continueNode {
  return continueNode { location }
}

func (self continueNode) String() string {
  return "(continue)"
}

func (self continueNode) IsStmt() bool {
  return true
}

func (self continueNode) GetLocation() ILocation {
  return self.location
}

// ExprStmtNode
type exprStmtNode struct {
  location ILocation
  Expr IExprNode
}

func ExprStmtNode(location ILocation, expr IExprNode) exprStmtNode {
  return exprStmtNode { location, expr }
}

func (self exprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self exprStmtNode) IsStmt() bool {
  return true
}

func (self exprStmtNode) GetLocation() ILocation {
  return self.location
}

// GotoNode
type gotoNode struct {
  location ILocation
  Target string
}

func GotoNode(location ILocation, target string) gotoNode {
  return gotoNode { location, target }
}

func (self gotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self gotoNode) IsStmt() bool {
  return true
}

func (self gotoNode) GetLocation() ILocation {
  return self.location
}

// LabelNode
type labelNode struct {
  location ILocation
  Name string
  Stmt IStmtNode
}

func LabelNode(location ILocation, name string, stmt IStmtNode) labelNode {
  return labelNode { location, name, stmt }
}

func (self labelNode) String() string {
  panic("not implemented")
}

func (self labelNode) IsStmt() bool {
  return true
}

func (self labelNode) GetLocation() ILocation {
  return self.location
}

// ReturnNode
type returnNode struct {
  location ILocation
  Expr IExprNode
}

func ReturnNode(location ILocation, expr IExprNode) returnNode {
  return returnNode { location, expr }
}

func (self returnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self returnNode) IsStmt() bool {
  return true
}

func (self returnNode) GetLocation() ILocation {
  return self.location
}
