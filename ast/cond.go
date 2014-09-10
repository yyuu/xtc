package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// CondExprNode
type CondExprNode struct {
  ClassName string
  Location duck.Location
  Cond duck.IExprNode
  ThenExpr duck.IExprNode
  ElseExpr duck.IExprNode
}

func NewCondExprNode(loc duck.Location, cond duck.IExprNode, thenExpr duck.IExprNode, elseExpr duck.IExprNode) CondExprNode {
  if cond == nil { panic("cond is nil") }
  if thenExpr == nil { panic("thenExpr is nil") }
  if elseExpr == nil { panic("elseExpr is nil") }
  return CondExprNode { "ast.CondExprNode", loc, cond, thenExpr, elseExpr }
}

func (self CondExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
}

func (self CondExprNode) IsExprNode() bool {
  return true
}

func (self CondExprNode) GetLocation() duck.Location {
  return self.Location
}

// CaseNode
type CaseNode struct {
  ClassName string
  Location duck.Location
  Values []duck.IExprNode
  Body duck.IStmtNode
}

func NewCaseNode(loc duck.Location, values []duck.IExprNode, body duck.IStmtNode) CaseNode {
  if body == nil { panic("body is nil") }
  return CaseNode { "ast.CaseNode", loc, values, body }
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

func (self CaseNode) IsStmtNode() bool {
  return true
}

func (self CaseNode) GetLocation() duck.Location {
  return self.Location
}

// IfNode
type IfNode struct {
  ClassName string
  Location duck.Location
  Cond duck.IExprNode
  ThenBody duck.IStmtNode
  ElseBody duck.IStmtNode
}

func NewIfNode(loc duck.Location, cond duck.IExprNode, thenBody duck.IStmtNode, elseBody duck.IStmtNode) IfNode {
  if cond == nil { panic("cond is nil") }
  if thenBody == nil { panic("thenBody is nil") }
  return IfNode { "ast.IfNode", loc, cond, thenBody, elseBody }
}

func (self IfNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenBody, self.ElseBody)
}

func (self IfNode) IsStmtNode() bool {
  return true
}

func (self IfNode) GetLocation() duck.Location {
  return self.Location
}

// SwitchNode
type SwitchNode struct {
  ClassName string
  Location duck.Location
  Cond duck.IExprNode
  Cases []CaseNode
}

func NewSwitchNode(loc duck.Location, cond duck.IExprNode, _cases []duck.IStmtNode) SwitchNode {
  if cond == nil { panic("cond is nil") }
  cases := make([]CaseNode, len(_cases))
  for i := range _cases {
    cases[i] = _cases[i].(CaseNode)
  }
  return SwitchNode { "ast.SwitchNode", loc, cond, cases }
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

func (self SwitchNode) IsStmtNode() bool {
  return true
}

func (self SwitchNode) GetLocation() duck.Location {
  return self.Location
}
