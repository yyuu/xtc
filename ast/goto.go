package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// BreakNode
type BreakNode struct {
  ClassName string
  Location duck.Location
}

func NewBreakNode(loc duck.Location) BreakNode {
  return BreakNode { "ast.BreakNode", loc }
}

func (self BreakNode) String() string {
  return "(break)"
}

func (self BreakNode) IsStmtNode() bool {
  return true
}

func (self BreakNode) GetLocation() duck.Location {
  return self.Location
}

// ContinueNode
type ContinueNode struct {
  ClassName string
  Location duck.Location
}

func NewContinueNode(loc duck.Location) ContinueNode {
  return ContinueNode { "ast.ContinueNode", loc }
}

func (self ContinueNode) String() string {
  return "(continue)"
}

func (self ContinueNode) IsStmtNode() bool {
  return true
}

func (self ContinueNode) GetLocation() duck.Location {
  return self.Location
}

// ExprStmtNode
type ExprStmtNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
}

func NewExprStmtNode(loc duck.Location, expr duck.IExprNode) ExprStmtNode {
  if expr == nil { panic("expr is nil") }
  return ExprStmtNode { "ast.ExprStmtNode", loc, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ExprStmtNode) IsStmtNode() bool {
  return true
}

func (self ExprStmtNode) GetLocation() duck.Location {
  return self.Location
}

// GotoNode
type GotoNode struct {
  ClassName string
  Location duck.Location
  Target string
}

func NewGotoNode(loc duck.Location, target string) GotoNode {
  return GotoNode { "ast.GotoNode", loc, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self GotoNode) IsStmtNode() bool {
  return true
}

func (self GotoNode) GetLocation() duck.Location {
  return self.Location
}

// LabelNode
type LabelNode struct {
  ClassName string
  Location duck.Location
  Name string
  Stmt duck.IStmtNode
}

func NewLabelNode(loc duck.Location, name string, stmt duck.IStmtNode) LabelNode {
  if stmt == nil { panic("stmt is nil") }
  return LabelNode { "ast.LabelNode", loc, name, stmt }
}

func (self LabelNode) String() string {
  panic("not implemented")
}

func (self LabelNode) IsStmtNode() bool {
  return true
}

func (self LabelNode) GetLocation() duck.Location {
  return self.Location
}

// ReturnNode
type ReturnNode struct {
  ClassName string
  Location duck.Location
  Expr duck.IExprNode
}

func NewReturnNode(loc duck.Location, expr duck.IExprNode) ReturnNode {
  if expr == nil { panic("expr is nil") }
  return ReturnNode { "ast.ReturnNode", loc, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ReturnNode) IsStmtNode() bool {
  return true
}

func (self ReturnNode) GetLocation() duck.Location {
  return self.Location
}

func (self ReturnNode) GetExpr() duck.IExprNode {
  return self.Expr
}
