package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// BreakNode
type BreakNode struct {
  Location duck.ILocation
}

func NewBreakNode(location duck.ILocation) BreakNode {
  return BreakNode { location }
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
  x.Location = self.Location
  return json.Marshal(x)
}

func (self BreakNode) IsStmt() bool {
  return true
}

func (self BreakNode) GetLocation() duck.ILocation {
  return self.Location
}

// ContinueNode
type ContinueNode struct {
  Location duck.ILocation
}

func NewContinueNode(location duck.ILocation) ContinueNode {
  return ContinueNode { location }
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
  x.Location = self.Location
  return json.Marshal(x)
}

func (self ContinueNode) IsStmt() bool {
  return true
}

func (self ContinueNode) GetLocation() duck.ILocation {
  return self.Location
}

// ExprStmtNode
type ExprStmtNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
}

func NewExprStmtNode(location duck.ILocation, expr duck.IExprNode) ExprStmtNode {
  return ExprStmtNode { location, expr }
}

func (self ExprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ExprStmtNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
  }
  x.ClassName = "ast.ExprStmtNode"
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self ExprStmtNode) IsStmt() bool {
  return true
}

func (self ExprStmtNode) GetLocation() duck.ILocation {
  return self.Location
}

// GotoNode
type GotoNode struct {
  Location duck.ILocation
  Target string
}

func NewGotoNode(location duck.ILocation, target string) GotoNode {
  return GotoNode { location, target }
}

func (self GotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
}

func (self GotoNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
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

func (self GotoNode) GetLocation() duck.ILocation {
  return self.Location
}

// LabelNode
type LabelNode struct {
  Location duck.ILocation
  Name string
  Stmt duck.IStmtNode
}

func NewLabelNode(location duck.ILocation, name string, stmt duck.IStmtNode) LabelNode {
  return LabelNode { location, name, stmt }
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
  x.Location = self.Location
  x.Name = self.Name
  x.Stmt = self.Stmt
  return json.Marshal(x)
}

func (self LabelNode) IsStmt() bool {
  return true
}

func (self LabelNode) GetLocation() duck.ILocation {
  return self.Location
}

// ReturnNode
type ReturnNode struct {
  Location duck.ILocation
  Expr duck.IExprNode
}

func NewReturnNode(location duck.ILocation, expr duck.IExprNode) ReturnNode {
  return ReturnNode { location, expr }
}

func (self ReturnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

func (self ReturnNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Expr duck.IExprNode
  }
  x.ClassName = "ast.ReturnNode"
  x.Location = self.Location
  x.Expr = self.Expr
  return json.Marshal(x)
}

func (self ReturnNode) IsStmt() bool {
  return true
}

func (self ReturnNode) GetLocation() duck.ILocation {
  return self.Location
}
