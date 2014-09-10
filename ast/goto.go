package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// BreakNode
type BreakNode struct {
  ClassName string
  Location core.Location
}

func NewBreakNode(loc core.Location) BreakNode {
  return BreakNode { "ast.BreakNode", loc }
}

func (self BreakNode) String() string {
  return "(break)"
}

func (self BreakNode) IsStmtNode() bool {
  return true
}

func (self BreakNode) GetLocation() core.Location {
  return self.Location
}

// ContinueNode
type ContinueNode struct {
  ClassName string
  Location core.Location
}

func NewContinueNode(loc core.Location) ContinueNode {
  return ContinueNode { "ast.ContinueNode", loc }
}

func (self ContinueNode) String() string {
  return "(continue)"
}

func (self ContinueNode) IsStmtNode() bool {
  return true
}

func (self ContinueNode) GetLocation() core.Location {
  return self.Location
}

// ExprStmtNode
type ExprStmtNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewExprStmtNode(loc core.Location, expr core.IExprNode) ExprStmtNode {
  if expr == nil { panic("expr is nil") }
  return ExprStmtNode { "ast.ExprStmtNode", loc, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ExprStmtNode) IsStmtNode() bool {
  return true
}

func (self ExprStmtNode) GetLocation() core.Location {
  return self.Location
}

// GotoNode
type GotoNode struct {
  ClassName string
  Location core.Location
  Target string
}

func NewGotoNode(loc core.Location, target string) GotoNode {
  return GotoNode { "ast.GotoNode", loc, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self GotoNode) IsStmtNode() bool {
  return true
}

func (self GotoNode) GetLocation() core.Location {
  return self.Location
}

// LabelNode
type LabelNode struct {
  ClassName string
  Location core.Location
  Name string
  Stmt core.IStmtNode
}

func NewLabelNode(loc core.Location, name string, stmt core.IStmtNode) LabelNode {
  if stmt == nil { panic("stmt is nil") }
  return LabelNode { "ast.LabelNode", loc, name, stmt }
}

func (self LabelNode) String() string {
  panic("not implemented")
}

func (self LabelNode) IsStmtNode() bool {
  return true
}

func (self LabelNode) GetLocation() core.Location {
  return self.Location
}

// ReturnNode
type ReturnNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
}

func NewReturnNode(loc core.Location, expr core.IExprNode) ReturnNode {
  if expr == nil { panic("expr is nil") }
  return ReturnNode { "ast.ReturnNode", loc, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ReturnNode) IsStmtNode() bool {
  return true
}

func (self ReturnNode) GetLocation() core.Location {
  return self.Location
}

func (self ReturnNode) GetExpr() core.IExprNode {
  return self.Expr
}
