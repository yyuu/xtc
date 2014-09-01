package ast

import (
  "encoding/json"
  "fmt"
)

// BreakNode
type BreakNode struct {
  Location Location
}

func NewBreakNode(location Location) BreakNode {
  return BreakNode { location }
}

func (self BreakNode) String() string {
  return "(break)"
}

func (self BreakNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
  }
  x.ClassName = "ast.BreakNode"
  x.Location = self.Location
  return json.Marshal(x)
}

func (self BreakNode) IsStmt() bool {
  return true
}

func (self BreakNode) GetLocation() Location {
  return self.Location
}

// ContinueNode
type ContinueNode struct {
  Location Location
}

func NewContinueNode(location Location) ContinueNode {
  return ContinueNode { location }
}

func (self ContinueNode) String() string {
  return "(continue)"
}

func (self ContinueNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
  }
  x.ClassName = "ast.ContinueNode"
  x.Location = self.Location
  return json.Marshal(x)
}

func (self ContinueNode) IsStmt() bool {
  return true
}

func (self ContinueNode) GetLocation() Location {
  return self.Location
}

// ExprStmtNode
type ExprStmtNode struct {
  Location Location
  Expr IExprNode
}

func NewExprStmtNode(location Location, expr IExprNode) ExprStmtNode {
  return ExprStmtNode { location, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ExprStmtNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
  }
  x.ClassName = "ast.ExprStmtNode"
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self ExprStmtNode) IsStmt() bool {
  return true
}

func (self ExprStmtNode) GetLocation() Location {
  return self.Location
}

// GotoNode
type GotoNode struct {
  Location Location
  Target string
}

func NewGotoNode(location Location, target string) GotoNode {
  return GotoNode { location, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self GotoNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Target string
  }
  x.ClassName = "ast.GotoNode"
  x.Location = self.Location
  x.Target = self.Target
  return json.Marshal(x)
}

func (self GotoNode) IsStmt() bool {
  return true
}

func (self GotoNode) GetLocation() Location {
  return self.Location
}

// LabelNode
type LabelNode struct {
  Location Location
  Name string
  Stmt IStmtNode
}

func NewLabelNode(location Location, name string, stmt IStmtNode) LabelNode {
  return LabelNode { location, name, stmt }
}

func (self LabelNode) String() string {
  panic("not implemented")
}

func (self LabelNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Name string
    Stmt IStmtNode
  }
  x.ClassName = "ast.LabelNode"
  x.Location = self.Location
  x.Name = self.Name
  x.Stmt = self.Stmt
  return json.Marshal(x)
}

func (self LabelNode) IsStmt() bool {
  return true
}

func (self LabelNode) GetLocation() Location {
  return self.Location
}

// ReturnNode
type ReturnNode struct {
  Location Location
  Expr IExprNode
}

func NewReturnNode(location Location, expr IExprNode) ReturnNode {
  return ReturnNode { location, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ReturnNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location Location
    Expr IExprNode
  }
  x.ClassName = "ast.ReturnNode"
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self ReturnNode) IsStmt() bool {
  return true
}

func (self ReturnNode) GetLocation() Location {
  return self.Location
}
