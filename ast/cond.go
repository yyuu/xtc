package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

// CondExprNode
type CondExprNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  ThenExpr core.IExprNode
  ElseExpr core.IExprNode
}

func NewCondExprNode(loc core.Location, cond core.IExprNode, thenExpr core.IExprNode, elseExpr core.IExprNode) CondExprNode {
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

func (self CondExprNode) GetLocation() core.Location {
  return self.Location
}

// CaseNode
type CaseNode struct {
  ClassName string
  Location core.Location
  Values []core.IExprNode
  Body core.IStmtNode
}

func NewCaseNode(loc core.Location, values []core.IExprNode, body core.IStmtNode) CaseNode {
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

func (self CaseNode) GetLocation() core.Location {
  return self.Location
}

// IfNode
type IfNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  ThenBody core.IStmtNode
  ElseBody core.IStmtNode
}

func NewIfNode(loc core.Location, cond core.IExprNode, thenBody core.IStmtNode, elseBody core.IStmtNode) IfNode {
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

func (self IfNode) GetLocation() core.Location {
  return self.Location
}

// SwitchNode
type SwitchNode struct {
  ClassName string
  Location core.Location
  Cond core.IExprNode
  Cases []CaseNode
}

func NewSwitchNode(loc core.Location, cond core.IExprNode, _cases []core.IStmtNode) SwitchNode {
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

func (self SwitchNode) GetLocation() core.Location {
  return self.Location
}
