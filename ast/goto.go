package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// BreakNode
type BreakNode struct {
  ClassName string
  Location duck.ILocation
}

func NewBreakNode(loc duck.ILocation) BreakNode {
  if loc == nil { panic("location is nil") }
  return BreakNode { "ast.BreakNode", loc }
}

func (self BreakNode) String() string {
  return "(break)"
}

func (self BreakNode) IsStmtNode() bool {
  return true
}

func (self BreakNode) GetLocation() duck.ILocation {
  return self.Location
}

// ContinueNode
type ContinueNode struct {
  ClassName string
  Location duck.ILocation
}

func NewContinueNode(loc duck.ILocation) ContinueNode {
  if loc == nil { panic("location is nil") }
  return ContinueNode { "ast.ContinueNode", loc }
}

func (self ContinueNode) String() string {
  return "(continue)"
}

func (self ContinueNode) IsStmtNode() bool {
  return true
}

func (self ContinueNode) GetLocation() duck.ILocation {
  return self.Location
}

// ExprStmtNode
type ExprStmtNode struct {
  ClassName string
  Location duck.ILocation
  Expr duck.IExprNode
}

func NewExprStmtNode(loc duck.ILocation, expr duck.IExprNode) ExprStmtNode {
  if loc == nil { panic("location is nil") }
  if expr == nil { panic("expr is nil") }
  return ExprStmtNode { "ast.ExprStmtNode", loc, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ExprStmtNode) IsStmtNode() bool {
  return true
}

func (self ExprStmtNode) GetLocation() duck.ILocation {
  return self.Location
}

// GotoNode
type GotoNode struct {
  ClassName string
  Location duck.ILocation
  Target string
}

func NewGotoNode(loc duck.ILocation, target string) GotoNode {
  if loc == nil { panic("location is nil") }
  return GotoNode { "ast.GotoNode", loc, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self GotoNode) IsStmtNode() bool {
  return true
}

func (self GotoNode) GetLocation() duck.ILocation {
  return self.Location
}

// LabelNode
type LabelNode struct {
  ClassName string
  Location duck.ILocation
  Name string
  Stmt duck.IStmtNode
}

func NewLabelNode(loc duck.ILocation, name string, stmt duck.IStmtNode) LabelNode {
  if loc == nil { panic("location is nil") }
  if stmt == nil { panic("stmt is nil") }
  return LabelNode { "ast.LabelNode", loc, name, stmt }
}

func (self LabelNode) String() string {
  panic("not implemented")
}

func (self LabelNode) IsStmtNode() bool {
  return true
}

func (self LabelNode) GetLocation() duck.ILocation {
  return self.Location
}

// ReturnNode
type ReturnNode struct {
  ClassName string
  Location duck.ILocation
  Expr duck.IExprNode
}

func NewReturnNode(loc duck.ILocation, expr duck.IExprNode) ReturnNode {
  if loc == nil { panic("location is nil") }
  if expr == nil { panic("expr is nil") }
  return ReturnNode { "ast.ReturnNode", loc, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ReturnNode) IsStmtNode() bool {
  return true
}

func (self ReturnNode) GetLocation() duck.ILocation {
  return self.Location
}

func (self ReturnNode) GetExpr() duck.IExprNode {
  return self.Expr
}
