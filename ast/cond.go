package ast

import (
  "fmt"
  "strings"
)

// CondExprNode
type condExprNode struct {
  location ILocation
  Cond IExprNode
  ThenExpr IExprNode
  ElseExpr IExprNode
}

func CondExprNode(location ILocation, cond IExprNode, thenExpr IExprNode, elseExpr IExprNode) condExprNode {
  return condExprNode { location, cond, thenExpr, elseExpr }
}

func (self condExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

func (self condExprNode) IsExpr() bool {
  return true
}

func (self condExprNode) Location() ILocation {
  return self.location
}

// CaseNode
type caseNode struct {
  location ILocation
  Values []IExprNode
  Body IStmtNode
}

func CaseNode(location ILocation, values []IExprNode, body IStmtNode) caseNode {
  return caseNode { location, values, body }
}

func (self caseNode) String() string {
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

func (self caseNode) IsStmt() bool {
  return true
}

func (self caseNode) Location() ILocation {
  return self.location
}

// IfNode
type ifNode struct {
  location ILocation
  Cond IExprNode
  ThenBody IStmtNode
  ElseBody IStmtNode
}

func IfNode(location ILocation, cond IExprNode, thenBody IStmtNode, elseBody IStmtNode) ifNode {
  return ifNode { location, cond, thenBody, elseBody }
}

func (self ifNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenBody, self.ElseBody)
}

func (self ifNode) IsStmt() bool {
  return true
}

func (self ifNode) Location() ILocation {
  return self.location
}

// SwitchNode
type switchNode struct {
  location ILocation
  Cond IExprNode
  Cases []caseNode
}

func SwitchNode(location ILocation, cond IExprNode, _cases []IStmtNode) switchNode {
  cases := make([]caseNode, len(_cases))
  for i := range _cases {
    cases[i] = _cases[i].(caseNode)
  }
  return switchNode { location, cond, cases }
}

func (self switchNode) String() string {
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

func (self switchNode) IsStmt() bool {
  return true
}

func (self switchNode) Location() ILocation {
  return self.location
}
