package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

// IfNode
type IfNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  ThenBody core.IStmtNode
  ElseBody core.IStmtNode
}

func NewIfNode(loc core.Location, cond core.IExprNode, thenBody core.IStmtNode, elseBody core.IStmtNode) *IfNode {
  if cond == nil { panic("cond is nil") }
  if thenBody == nil { panic("thenBody is nil") }
  return &IfNode { "ast.IfNode", loc, cond, thenBody, elseBody }
}

func (self IfNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenBody, self.ElseBody)
}

func (self *IfNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self IfNode) GetLocation() core.Location {
  return self.Location
}

func (self *IfNode) GetCond() core.IExprNode {
  return self.Cond
}

func (self *IfNode) GetThenBody() core.IStmtNode {
  return self.ThenBody
}

func (self *IfNode) GetElseBody() core.IStmtNode {
  return self.ElseBody
}

func (self IfNode) GetScope() core.IScope {
  panic("#GetScope called")
}

func (self *IfNode) HasElseBody() bool {
  return self.ElseBody != nil
}
