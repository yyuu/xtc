package ast

import (
  "fmt"
  "strings"
)

type BlockNode struct {
// Variables []DefinedVariable
  Variables []IExprNode
  Stmts []IStmtNode
}

func (self BlockNode) DumpString() string {
  sVariables := make([]string, len(self.Variables))
  for i := range self.Variables {
    sVariables[i] = fmt.Sprintf("(a%d %s)", i, self.Variables[i].DumpString())
  }
  sStmts := make([]string, len(self.Stmts))
  for j := range self.Stmts {
    sStmts[j] = self.Stmts[j].DumpString()
  }
  switch len(sStmts) {
    case 0:  return fmt.Sprintf("(let* (%s))", strings.Join(sVariables, " "))
    case 1:  return fmt.Sprintf("(let* (%s) %s)", strings.Join(sVariables, " "), sStmts[0])
    default: return fmt.Sprintf("(let* (%s) (begin %s))", strings.Join(sVariables, " "), strings.Join(sStmts, " "))
  }
}

type CaseNode struct {
  Values []IExprNode
  Body BlockNode
}

func (self CaseNode) DumpString() string {
  sValues := make([]string, len(self.Values))
  for i := range self.Values {
    sValues[i] = fmt.Sprintf("(= a %s)", self.Values[i].DumpString())
  }
  switch len(sValues) {
    case 0:  return fmt.Sprintf("(() %s)", self.Body.DumpString())
    case 1:  return fmt.Sprintf("(%s %s)", sValues[0], self.Body.DumpString())
    default: return fmt.Sprintf("((or %s) %s)", strings.Join(sValues, " "), self.Body.DumpString())
  }
}

type ContinueNode struct {
}

type DoWhileNode struct {
  Body IStmtNode
  Cond IExprNode
}

func (self DoWhileNode) DumpString() string {
  return fmt.Sprintf("(let loop () (begin %s (if %s (loop))))", self.Body.DumpString(), self.Cond.DumpString())
}

type ExprStmtNode struct {
  Expr IExprNode
}

type ForNode struct {
  Init IExprNode
  Cond IExprNode
  Incr IExprNode
  Body IStmtNode
}

type GotoNode struct {
  Target string
}

type IfNode struct {
  Cond IExprNode
  ThenBody IStmtNode
  ElseBody IStmtNode
}

func (self IfNode) DumpString() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond.DumpString(), self.ThenBody.DumpString(), self.ElseBody.DumpString())
}

type LabelNode struct {
  Name string
  Stmt IStmtNode
}

type ReturnNode struct {
  Expr IExprNode
}

func (self ReturnNode) DumpString() string {
  return self.Expr.DumpString()
}

type SwitchNode struct {
  Cond IExprNode
  Cases []CaseNode
}

func (self SwitchNode) DumpString() string {
  sCases := make([]string, len(self.Cases))
  for i := range self.Cases {
    sCases[i] = self.Cases[i].DumpString()
  }
  if len(sCases) == 0 {
    return fmt.Sprintf("(let ((a %s)) ())", self.Cond.DumpString())
  } else {
    return fmt.Sprintf("(let ((a %s)) (cond %s))", self.Cond.DumpString(), strings.Join(sCases, " "))
  }
}

type WhileNode struct {
  Cond IExprNode
  Body IStmtNode
}

func (self WhileNode) DumpString() string {
  return fmt.Sprintf("(let loop ((a %s)) (if a (begin %s (loop %s))))", self.Cond.DumpString(), self.Body.DumpString(), self.Cond.DumpString())
}
