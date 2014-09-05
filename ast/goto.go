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

func (self BreakNode) IsStmt() bool {
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

func (self ContinueNode) IsStmt() bool {
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

func (self ExprStmtNode) IsStmt() bool {
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

func (self GotoNode) IsStmt() bool {
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

func (self LabelNode) IsStmt() bool {
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

func (self ReturnNode) IsStmt() bool {
  return true
}

func (self ReturnNode) GetLocation() duck.ILocation {
  return self.location
}

func (self ReturnNode) GetExpr() duck.IExprNode {
  return self.expr
}
