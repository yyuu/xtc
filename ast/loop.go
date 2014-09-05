package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

// DoWhileNode
type DoWhileNode struct {
  location duck.ILocation
  body duck.IStmtNode
  cond duck.IExprNode
}

func NewDoWhileNode(loc duck.ILocation, body duck.IStmtNode, cond duck.IExprNode) DoWhileNode {
  return DoWhileNode { loc, body, cond }
}

func (self DoWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.body, self.cond)
}

func (self DoWhileNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Body duck.IStmtNode
    Cond duck.IExprNode
  }
  x.ClassName = "ast.DoWhileNode"
  x.Location = self.location
  x.Body = self.body
  x.Cond = self.cond
  return json.Marshal(x)
}

func (self DoWhileNode) IsStmt() bool {
  return true
}

func (self DoWhileNode) GetLocation() duck.ILocation {
  return self.location
}

// ForNode
type ForNode struct {
  location duck.ILocation
  init duck.IExprNode
  cond duck.IExprNode
  incr duck.IExprNode
  body duck.IStmtNode
}

func NewForNode(loc duck.ILocation, init duck.IExprNode, cond duck.IExprNode, incr duck.IExprNode, body duck.IStmtNode) ForNode {
  return ForNode { loc, init, cond, incr, body }
}

func (self ForNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.init, self.cond, self.body, self.incr)
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
  x.Location = self.location
  x.Init = self.init
  x.Cond = self.cond
  x.Incr = self.incr
  x.Body = self.body
  return json.Marshal(x)
}

func (self ForNode) IsStmt() bool {
  return true
}

func (self ForNode) GetLocation() duck.ILocation {
  return self.location
}

// WhileNode
type WhileNode struct {
  location duck.ILocation
  cond duck.IExprNode
  body duck.IStmtNode
}

func NewWhileNode(loc duck.ILocation, cond duck.IExprNode, body duck.IStmtNode) WhileNode {
  return WhileNode { loc, cond, body }
}

func (self WhileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.cond, self.body, self.cond)
}

func (self WhileNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Cond duck.IExprNode
    Body duck.IStmtNode
  }
  x.ClassName = "ast.WhileNode"
  x.Location = self.location
  x.Cond = self.cond
  x.Body = self.body
  return json.Marshal(x)
}

func (self WhileNode) IsStmt() bool {
  return true
}

func (self WhileNode) GetLocation() duck.ILocation {
  return self.location
}
