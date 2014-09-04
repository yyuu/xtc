package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// DoWhileNode
type DoWhileNode struct {
  Location duck.ILocation
  Body duck.IStmtNode
  Cond duck.IExprNode
}

func NewDoWhileNode(location duck.ILocation, body duck.IStmtNode, cond duck.IExprNode) DoWhileNode {
  return DoWhileNode { location, body, cond }
}

func (self DoWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

func (self DoWhileNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Body duck.IStmtNode
    Cond duck.IExprNode
  }
  x.ClassName = "ast.DoWhileNode"
  x.Location = self.Location
  x.Body = self.Body
  x.Cond = self.Cond
  return json.Marshal(x)
}

func (self DoWhileNode) IsStmt() bool {
  return true
}

func (self DoWhileNode) GetLocation() duck.ILocation {
  return self.Location
}

// ForNode
type ForNode struct {
  Location duck.ILocation
  Init duck.IExprNode
  Cond duck.IExprNode
  Incr duck.IExprNode
  Body duck.IStmtNode
}

func NewForNode(location duck.ILocation, init duck.IExprNode, cond duck.IExprNode, incr duck.IExprNode, body duck.IStmtNode) ForNode {
  return ForNode { location, init, cond, incr, body }
}

func (self ForNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.Init, self.Cond, self.Body, self.Incr)
}

func (self ForNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Init duck.IExprNode
    Cond duck.IExprNode
    Incr duck.IExprNode
    Body duck.IStmtNode
  }
  x.ClassName = "ast.ForNode"
  x.Location = self.Location
  x.Init = self.Init
  x.Cond = self.Cond
  x.Incr = self.Incr
  x.Body = self.Body
  return json.Marshal(x)
}

func (self ForNode) IsStmt() bool {
  return true
}

func (self ForNode) GetLocation() duck.ILocation {
  return self.Location
}

// WhileNode
type WhileNode struct {
  Location duck.ILocation
  Cond duck.IExprNode
  Body duck.IStmtNode
}

func NewWhileNode(location duck.ILocation, cond duck.IExprNode, body duck.IStmtNode) WhileNode {
  return WhileNode { location, cond, body }
}

func (self WhileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}

func (self WhileNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Cond duck.IExprNode
    Body duck.IStmtNode
  }
  x.ClassName = "ast.WhileNode"
  x.Location = self.Location
  x.Cond = self.Cond
  x.Body = self.Body
  return json.Marshal(x)
}

func (self WhileNode) IsStmt() bool {
  return true
}

func (self WhileNode) GetLocation() duck.ILocation {
  return self.Location
}
