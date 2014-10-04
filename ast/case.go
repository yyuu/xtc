package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/xtc/asm"
  "bitbucket.org/yyuu/xtc/core"
)

// CaseNode
type CaseNode struct {
  ClassName string
  Location core.Location
  Label *asm.Label
  Values []core.IExprNode
  Body core.IStmtNode
}

func NewCaseNode(loc core.Location, values []core.IExprNode, body core.IStmtNode) *CaseNode {
  if body == nil { panic("body is nil") }
  return &CaseNode { "ast.CaseNode", loc, nil, values, body }
}

func (self CaseNode) String() string {
  sValues := make([]string, len(self.Values))
  for i := range self.Values {
    sValues[i] = fmt.Sprintf("(= switch-cond %s)", self.Values[i])
  }
  switch len(sValues) {
    case 0:  return fmt.Sprintf("(else %s)", self.Body)
    case 1:  return fmt.Sprintf("(%s %s)", sValues[0], self.Body)
    default: return fmt.Sprintf("((or %s) %s)", strings.Join(sValues, " "), self.Body)
  }
}

func (self *CaseNode) AsStmtNode() core.IStmtNode {
  return self
}

func (self CaseNode) GetLocation() core.Location {
  return self.Location
}

func (self *CaseNode) GetLabel() *asm.Label {
  return self.Label
}

func (self *CaseNode) SetLabel(label *asm.Label) {
  self.Label = label
}

func (self *CaseNode) GetValues() []core.IExprNode {
  return self.Values
}

func (self *CaseNode) GetBody() core.IStmtNode {
  return self.Body
}

func (self *CaseNode) IsDefault() bool {
  return len(self.Values) == 0
}

func (self CaseNode) GetScope() core.IScope {
  panic("#GetScope called")
}
