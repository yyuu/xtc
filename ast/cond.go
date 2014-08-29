package ast

import (
  "fmt"
  "strings"
)

type caseNode struct {
  Values []IExprNode
  Body IStmtNode
}

func CaseNode(_values []INode, body INode) caseNode {
  values := make([]IExprNode, len(_values))
  for i := range _values {
    values[i] = _values[i].(IExprNode)
  }
  return caseNode { values, body.(IStmtNode) }
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

type ifNode struct {
  Cond IExprNode
  ThenBody IStmtNode
  ElseBody IStmtNode
}

func IfNode(cond INode, thenBody INode, elseBody INode) ifNode {
  return ifNode { cond.(IExprNode), thenBody.(IStmtNode), elseBody.(IStmtNode) }
}

func (self ifNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenBody, self.ElseBody)
}

type switchNode struct {
  Cond IExprNode
  Cases []caseNode
}

func SwitchNode(cond INode, _cases []INode) switchNode {
  cases := make([]caseNode, len(_cases))
  for i := range _cases {
    cases[i] = _cases[i].(caseNode)
  }
  return switchNode { cond.(IExprNode), cases }
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
