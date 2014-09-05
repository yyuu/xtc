package ast

import (
  "encoding/json"
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// CondExprNode
type CondExprNode struct {
  location duck.ILocation
  cond duck.IExprNode
  thenExpr duck.IExprNode
  elseExpr duck.IExprNode
}

func NewCondExprNode(loc duck.ILocation, cond duck.IExprNode, thenExpr duck.IExprNode, elseExpr duck.IExprNode) CondExprNode {
  return CondExprNode { loc, cond, thenExpr, elseExpr }
}

func (self CondExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.cond, self.thenExpr, self.elseExpr)
}

func (self CondExprNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Cond duck.IExprNode
    ThenExpr duck.IExprNode
    ElseExpr duck.IExprNode
  }
  x.ClassName = "ast.CondExprNode"
  x.Location = self.location
  x.Cond = self.cond
  x.ThenExpr = self.thenExpr
  x.ElseExpr = self.elseExpr
  return json.Marshal(x)
}

func (self CondExprNode) IsExpr() bool {
  return true
}

func (self CondExprNode) GetLocation() duck.ILocation {
  return self.location
}

// CaseNode
type CaseNode struct {
  location duck.ILocation
  values []duck.IExprNode
  body duck.IStmtNode
}

func NewCaseNode(loc duck.ILocation, values []duck.IExprNode, body duck.IStmtNode) CaseNode {
  return CaseNode { loc, values, body }
}

func (self CaseNode) String() string {
  sValues := make([]string, len(self.values))
  for i := range self.values {
    sValues[i] = fmt.Sprintf("(= switch-cond %s)", self.values[i])
  }
  switch len(sValues) {
    case 0:  return fmt.Sprintf("(else %s)", self.body)
    case 1:  return fmt.Sprintf("(%s %s)", sValues[0], self.body)
    default: return fmt.Sprintf("((or %s) %s)", strings.Join(sValues, " "), self.body)
  }
}

func (self CaseNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Values []duck.IExprNode
    Body duck.IStmtNode
  }
  x.ClassName = "ast.CaseNode"
  x.Location = self.location
  x.Values = self.values
  x.Body = self.body
  return json.Marshal(x)
}

func (self CaseNode) IsStmt() bool {
  return true
}

func (self CaseNode) GetLocation() duck.ILocation {
  return self.location
}

// IfNode
type IfNode struct {
  location duck.ILocation
  cond duck.IExprNode
  thenBody duck.IStmtNode
  elseBody duck.IStmtNode
}

func NewIfNode(loc duck.ILocation, cond duck.IExprNode, thenBody duck.IStmtNode, elseBody duck.IStmtNode) IfNode {
  return IfNode { loc, cond, thenBody, elseBody }
}

func (self IfNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.cond, self.thenBody, self.elseBody)
}

func (self IfNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Cond duck.IExprNode
    ThenBody duck.IStmtNode
    ElseBody duck.IStmtNode
  }
  x.ClassName = "ast.IfNode"
  x.Location = self.location
  x.Cond = self.cond
  x.ThenBody = self.thenBody
  x.ElseBody = self.elseBody
  return json.Marshal(x)
}

func (self IfNode) IsStmt() bool {
  return true
}

func (self IfNode) GetLocation() duck.ILocation {
  return self.location
}

// SwitchNode
type SwitchNode struct {
  location duck.ILocation
  cond duck.IExprNode
  cases []CaseNode
}

func NewSwitchNode(loc duck.ILocation, cond duck.IExprNode, _cases []duck.IStmtNode) SwitchNode {
  cases := make([]CaseNode, len(_cases))
  for i := range _cases {
    cases[i] = _cases[i].(CaseNode)
  }
  return SwitchNode { loc, cond, cases }
}

func (self SwitchNode) String() string {
  sCases := make([]string, len(self.cases))
  for i := range self.cases {
    sCases[i] = fmt.Sprintf("%s", self.cases[i])
  }
  if len(sCases) == 0 {
    return fmt.Sprintf("(let ((switch-cond %s)) ())", self.cond)
  } else {
    return fmt.Sprintf("(let ((switch-cond %s)) (cond %s))", self.cond, strings.Join(sCases, " "))
  }
}

func (self SwitchNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Cond duck.IExprNode
    Cases []CaseNode
  }
  x.ClassName = "ast.SwitchNode"
  x.Location = self.location
  x.Cond = self.cond
  x.Cases = self.cases
  return json.Marshal(x)
}

func (self SwitchNode) IsStmt() bool {
  return true
}

func (self SwitchNode) GetLocation() duck.ILocation {
  return self.location
}
