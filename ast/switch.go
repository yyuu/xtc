package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/xtc/core"
)

// SwitchNode
type SwitchNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  Cases []core.IStmtNode
}

func NewSwitchNode(loc core.Location, cond core.IExprNode, cases []core.IStmtNode) *SwitchNode {
  if cond == nil { panic("cond is nil") }
  for i := range cases {
    _, ok := cases[i].(*CaseNode)
    if ! ok {
      panic(fmt.Errorf("syntax error: not a case: %s", cases[i]))
    }
  }
  return &SwitchNode { "ast.SwitchNode", loc, cond, cases }
}

func (self SwitchNode) String() string {
  sCases := make([]string, len(self.Cases))
  for i := range self.Cases {
    sCases[i] = fmt.Sprintf("%s", self.Cases[i])
  }
  if len(sCases) == 0 {
    return fmt.Sprintf("(let ((switch-cond %s)) ())", self.Cond)
  } else {
    return fmt.Sprintf("(let ((switch-cond %s)) (cond %s))", self.Cond, strings.Join(sCases, " "))
  }
}

func (self *SwitchNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self SwitchNode) GetLocation() core.Location {
  return self.Location
}

func (self *SwitchNode) GetCond() core.IExprNode {
  return self.Cond
}

func (self *SwitchNode) GetCases() []core.IStmtNode {
  return self.Cases
}

func (self SwitchNode) GetScope() core.IScope {
  panic("#GetScope called")
}
