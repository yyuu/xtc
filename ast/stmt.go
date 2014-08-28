package ast

import (
  "fmt"
  "strings"
)

type blockNode struct {
// Variables []DefinedVariable
  Variables []IExprNode
  Stmts []IStmtNode
}

func BlockNode(_variables []INode, _stmts []INode) blockNode {
  variables := make([]IExprNode, len(_variables))
  stmts := make([]IStmtNode, len(_stmts))
  for i := range _variables {
    variables[i] = _variables[i].(IExprNode)
  }
  for j := range _stmts {
    stmts[j] = _stmts[j].(IStmtNode)
  }
  return blockNode { variables, stmts }
}

func (self blockNode) String() string {
  sVariables := make([]string, len(self.Variables))
  for i := range self.Variables {
    sVariables[i] = fmt.Sprintf("(a%d %s)", i, self.Variables[i])
  }
  sStmts := make([]string, len(self.Stmts))
  for j := range self.Stmts {
    sStmts[j] = fmt.Sprintf("%s", self.Stmts[j])
  }
  switch len(sStmts) {
    case 0:  return fmt.Sprintf("(let* (%s))", strings.Join(sVariables, " "))
    case 1:  return fmt.Sprintf("(let* (%s) %s)", strings.Join(sVariables, " "), sStmts[0])
    default: return fmt.Sprintf("(let* (%s) (begin %s))", strings.Join(sVariables, " "), strings.Join(sStmts, " "))
  }
}

type caseNode struct {
  Values []IExprNode
  Body blockNode
}

func (self caseNode) String() string {
  sValues := make([]string, len(self.Values))
  for i := range self.Values {
    sValues[i] = fmt.Sprintf("(= a %s)", self.Values[i])
  }
  switch len(sValues) {
    case 0:  return fmt.Sprintf("(() %s)", self.Body)
    case 1:  return fmt.Sprintf("(%s %s)", sValues[0], self.Body)
    default: return fmt.Sprintf("((or %s) %s)", strings.Join(sValues, " "), self.Body)
  }
}

type continueNode struct {
}

type doWhileNode struct {
  Body IStmtNode
  Cond IExprNode
}

func (self doWhileNode) String() string {
  return fmt.Sprintf("(let loop () (begin %s (if %s (loop))))", self.Body, self.Cond)
}

type exprStmtNode struct {
  Expr IExprNode
}

type forNode struct {
  Init IExprNode
  Cond IExprNode
  Incr IExprNode
  Body IStmtNode
}

type gotoNode struct {
  Target string
}

type ifNode struct {
  Cond IExprNode
  ThenBody IStmtNode
  ElseBody IStmtNode
}

func (self ifNode) String() string {
  return fmt.Sprintf("(if %s %s %s)", self.Cond, self.ThenBody, self.ElseBody)
}

type labelNode struct {
  Name string
  Stmt IStmtNode
}

type returnNode struct {
  Expr IExprNode
}

func (self returnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

type switchNode struct {
  Cond IExprNode
  Cases []caseNode
}

func (self switchNode) String() string {
  sCases := make([]string, len(self.Cases))
  for i := range self.Cases {
    sCases[i] = fmt.Sprintf("%s", self.Cases[i])
  }
  if len(sCases) == 0 {
    return fmt.Sprintf("(let ((a %s)) ())", self.Cond)
  } else {
    return fmt.Sprintf("(let ((a %s)) (cond %s))", self.Cond, strings.Join(sCases, " "))
  }
}

type whileNode struct {
  Cond IExprNode
  Body IStmtNode
}

func (self whileNode) String() string {
  return fmt.Sprintf("(let loop ((a %s)) (if a (begin %s (loop %s))))", self.Cond, self.Body, self.Cond)
}
