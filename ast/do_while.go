package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// DoWhileNode
type DoWhileNode struct {
  ClassName string
  Location core.Location
  Body core.IStmtNode
  Cond core.IExprNode
}

func NewDoWhileNode(loc core.Location, body core.IStmtNode, cond core.IExprNode) *DoWhileNode {
  if body == nil { panic("body is nil") }
  if cond == nil { panic("cond is nil") }
  return &DoWhileNode { "ast.DoWhileNode", loc, body, cond }
}

func (self DoWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

func (self DoWhileNode) IsStmtNode() bool {
  return true
}

func (self DoWhileNode) GetLocation() core.Location {
  return self.Location
}

func (self DoWhileNode) GetBody() core.IStmtNode {
  return self.Body
}

func (self DoWhileNode) GetCond() core.IExprNode {
  return self.Cond
}
