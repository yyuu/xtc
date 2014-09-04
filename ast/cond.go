package ast

import (
  "encoding/json"
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// CondExprNode
type CondExprNode struct {
  Location duck.ILocation
  Cond duck.IExprNode
  ThenExpr duck.IExprNode
  ElseExpr duck.IExprNode
}

func NewCondExprNode(location duck.ILocation, cond duck.IExprNode, thenExpr duck.IExprNode, elseExpr duck.IExprNode) CondExprNode {
  return CondExprNode { location, cond, thenExpr, elseExpr }
}

func (self CondExprNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenExpr, self.ElseExpr)
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
  x.Location = self.Location
  x.Cond = self.Cond
  x.ThenExpr = self.ThenExpr
  x.ElseExpr = self.ElseExpr
  return json.Marshal(x)
}

func (self CondExprNode) IsExpr() bool {
  return true
}

func (self CondExprNode) GetLocation() duck.ILocation {
  return self.Location
}

// CaseNode
type CaseNode struct {
  Location duck.ILocation
  Values []duck.IExprNode
  Body duck.IStmtNode
}

func NewCaseNode(location duck.ILocation, values []duck.IExprNode, body duck.IStmtNode) CaseNode {
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

func (self CaseNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Values []duck.IExprNode
    Body duck.IStmtNode
  }
  x.ClassName = "ast.CaseNode"
  x.Location = self.Location
  x.Values = self.Values
  x.Body = self.Body
  return json.Marshal(x)
}

func (self CaseNode) IsStmt() bool {
  return true
}

func (self CaseNode) GetLocation() duck.ILocation {
  return self.Location
}

// IfNode
type IfNode struct {
  Location duck.ILocation
  Cond duck.IExprNode
  ThenBody duck.IStmtNode
  ElseBody duck.IStmtNode
}

func NewIfNode(location duck.ILocation, cond duck.IExprNode, thenBody duck.IStmtNode, elseBody duck.IStmtNode) IfNode {
  return IfNode { location, cond, thenBody, elseBody }
}

func (self IfNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenBody, self.ElseBody)
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
  x.Location = self.Location
  x.Cond = self.Cond
  x.ThenBody = self.ThenBody
  x.ElseBody = self.ElseBody
  return json.Marshal(x)
}

func (self IfNode) IsStmt() bool {
  return true
}

func (self IfNode) GetLocation() duck.ILocation {
  return self.Location
}

// SwitchNode
type SwitchNode struct {
  Location duck.ILocation
  Cond duck.IExprNode
  Cases []CaseNode
}

func NewSwitchNode(location duck.ILocation, cond duck.IExprNode, _cases []duck.IStmtNode) SwitchNode {
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

func (self SwitchNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Cond duck.IExprNode
    Cases []CaseNode
  }
  x.ClassName = "ast.SwitchNode"
  x.Location = self.Location
  x.Cond = self.Cond
  x.Cases = self.Cases
  return json.Marshal(x)
}

func (self SwitchNode) IsStmt() bool {
  return true
}

func (self SwitchNode) GetLocation() duck.ILocation {
  return self.Location
}
