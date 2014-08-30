package ast

import (
  "fmt"
  "strings"
)

// CondExprNode
type CondExprNode struct {
  Location Location
  Cond IExprNode
  ThenExpr IExprNode
  ElseExpr IExprNode
}

func NewCondExprNode(location Location, cond IExprNode, thenExpr IExprNode, elseExpr IExprNode) CondExprNode {
  return CondExprNode { location, cond, thenExpr, elseExpr }
}

func (self CondExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

func (self CondExprNode) IsExpr() bool {
  return true
}

func (self CondExprNode) GetLocation() Location {
  return self.Location
}

// CaseNode
type CaseNode struct {
  Location Location
  Values []IExprNode
  Body IStmtNode
}

func NewCaseNode(location Location, values []IExprNode, body IStmtNode) CaseNode {
  return CaseNode { location, values, body }
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

func (self CaseNode) IsStmt() bool {
  return true
}

func (self CaseNode) GetLocation() Location {
  return self.Location
}

// IfNode
type IfNode struct {
  Location Location
  Cond IExprNode
  ThenBody IStmtNode
  ElseBody IStmtNode
}

func NewIfNode(location Location, cond IExprNode, thenBody IStmtNode, elseBody IStmtNode) IfNode {
  return IfNode { location, cond, thenBody, elseBody }
}

func (self IfNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenBody, self.ElseBody)
}

func (self IfNode) IsStmt() bool {
  return true
}

func (self IfNode) GetLocation() Location {
  return self.Location
}

// SwitchNode
type SwitchNode struct {
  Location Location
  Cond IExprNode
  Cases []CaseNode
}

func NewSwitchNode(location Location, cond IExprNode, _cases []IStmtNode) SwitchNode {
  cases := make([]CaseNode, len(_cases))
  for i := range _cases {
    cases[i] = _cases[i].(CaseNode)
  }
  return SwitchNode { location, cond, cases }
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

func (self SwitchNode) IsStmt() bool {
  return true
}

func (self SwitchNode) GetLocation() Location {
  return self.Location
}
