package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// BreakNode
type BreakNode struct {
  location duck.ILocation
}

func NewBreakNode(loc duck.ILocation) BreakNode {
  if loc == nil { panic("location is nil") }
  return BreakNode { loc }
}

func (self BreakNode) String() string {
  return "(break)"
}

func (self BreakNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
  }
  x.ClassName = "ast.BreakNode"
  x.Location = self.location
  return json.Marshal(x)
}

func (self BreakNode) IsStmtNode() bool {
  return true
}

func (self BreakNode) GetLocation() duck.ILocation {
  return self.location
}

// ContinueNode
type ContinueNode struct {
  location duck.ILocation
}

func NewContinueNode(loc duck.ILocation) ContinueNode {
  if loc == nil { panic("location is nil") }
  return ContinueNode { loc }
}

func (self ContinueNode) String() string {
  return "(continue)"
}

func (self ContinueNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
  }
  x.ClassName = "ast.ContinueNode"
  x.Location = self.location
  return json.Marshal(x)
}

func (self ContinueNode) IsStmtNode() bool {
  return true
}

func (self ContinueNode) GetLocation() duck.ILocation {
  return self.location
}

// ExprStmtNode
type ExprStmtNode struct {
  location duck.ILocation
  expr duck.IExprNode
}

func NewExprStmtNode(loc duck.ILocation, expr duck.IExprNode) ExprStmtNode {
  if loc == nil { panic("location is nil") }
  if expr == nil { panic("expr is nil") }
  return ExprStmtNode { loc, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.expr)
}

func (self ExprStmtNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
  }
  x.ClassName = "ast.ExprStmtNode"
  x.Location = self.location
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self ExprStmtNode) IsStmtNode() bool {
  return true
}

func (self ExprStmtNode) GetLocation() duck.ILocation {
  return self.location
}

// GotoNode
type GotoNode struct {
  location duck.ILocation
  target string
}

func NewGotoNode(loc duck.ILocation, target string) GotoNode {
  if loc == nil { panic("location is nil") }
  return GotoNode { loc, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.target)
}

func (self GotoNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Target string
  }
  x.ClassName = "ast.GotoNode"
  x.Location = self.location
  x.Target = self.target
  return json.Marshal(x)
}

func (self GotoNode) IsStmtNode() bool {
  return true
}

func (self GotoNode) GetLocation() duck.ILocation {
  return self.location
}

// LabelNode
type LabelNode struct {
  location duck.ILocation
  name string
  stmt duck.IStmtNode
}

func NewLabelNode(loc duck.ILocation, name string, stmt duck.IStmtNode) LabelNode {
  if loc == nil { panic("location is nil") }
  if stmt == nil { panic("stmt is nil") }
  return LabelNode { loc, name, stmt }
}

func (self LabelNode) String() string {
  panic("not implemented")
}

func (self LabelNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Name string
    Stmt duck.IStmtNode
  }
  x.ClassName = "ast.LabelNode"
  x.Location = self.location
  x.Name = self.name
  x.Stmt = self.stmt
  return json.Marshal(x)
}

func (self LabelNode) IsStmtNode() bool {
  return true
}

func (self LabelNode) GetLocation() duck.ILocation {
  return self.location
}

// ReturnNode
type ReturnNode struct {
  location duck.ILocation
  expr duck.IExprNode
}

func NewReturnNode(loc duck.ILocation, expr duck.IExprNode) ReturnNode {
  if loc == nil { panic("location is nil") }
  if expr == nil { panic("expr is nil") }
  return ReturnNode { loc, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.expr)
}

func (self ReturnNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
  }
  x.ClassName = "ast.ReturnNode"
  x.Location = self.location
  x.Expr = self.expr
  return json.Marshal(x)
}

func (self ReturnNode) IsStmtNode() bool {
  return true
}

func (self ReturnNode) GetLocation() duck.ILocation {
  return self.location
}

func (self ReturnNode) GetExpr() duck.IExprNode {
  return self.expr
}
