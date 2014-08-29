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

  stmts := ""
  switch len(sStmts) {
    case 0:  stmts = ""
    case 1:  stmts = fmt.Sprintf("%s", sStmts[0])
    default: stmts = fmt.Sprintf("(begin %s)", strings.Join(sStmts, " "))
  }

  switch len(sVariables) {
    case 0:  return stmts
    case 1:  return fmt.Sprintf("(let (%s) %s)", strings.Join(sVariables, " "), stmts)
    default: return fmt.Sprintf("(let* (%s) (begin %s))", strings.Join(sVariables, " "), strings.Join(sStmts, " "))
  }
}

type breakNode struct {
}

func BreakNode() breakNode {
  return breakNode { }
}

func (self breakNode) String() string {
  return "(break)"
}

type continueNode struct {
}

func ContinueNode() continueNode {
  return continueNode { }
}

func (self continueNode) String() string {
  return "(continue)"
}

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
    sValues[i] = fmt.Sprintf("(= a %s)", self.Values[i])
  }
  switch len(sValues) {
    case 0:  return fmt.Sprintf("(() %s)", self.Body)
    case 1:  return fmt.Sprintf("(%s %s)", sValues[0], self.Body)
    default: return fmt.Sprintf("((or %s) %s)", strings.Join(sValues, " "), self.Body)
  }
}

type doWhileNode struct {
  Body IStmtNode
  Cond IExprNode
}

func DoWhileNode(body INode, cond INode) doWhileNode {
  return doWhileNode { body.(IStmtNode), cond.(IExprNode) }
}

func (self doWhileNode) String() string {
  return fmt.Sprintf("(let do-while-loop () (begin %s (if %s (do-while-loop))))", self.Body, self.Cond)
}

type exprStmtNode struct {
  Expr IExprNode
}

func ExprStmtNode(expr INode) exprStmtNode {
  return exprStmtNode { expr.(IExprNode) }
}

func (self exprStmtNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
}

type forNode struct {
  Init IExprNode
  Cond IExprNode
  Incr IExprNode
  Body IStmtNode
}

func ForNode(init INode, cond INode, incr INode, body INode) forNode {
  return forNode { init.(IExprNode), cond.(IExprNode), incr.(IExprNode), body.(IExprNode) }
}

func (self forNode) String() string {
  return fmt.Sprintf("(let for-loop (%s) (if %s (begin %s (for-loop %s))))", self.Init, self.Cond, self.Body, self.Incr)
}

type gotoNode struct {
  Target string
}

func GotoNode(target string) gotoNode {
  return gotoNode { target }
}

func (self gotoNode) String() string {
  return fmt.Sprintf("(goto %s)", self.Target)
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

type labelNode struct {
  Name string
  Stmt IStmtNode
}

func LabelNode(name string, stmt INode) labelNode {
  return labelNode { name, stmt.(IStmtNode) }
}

func (self labelNode) String() string {
  panic("not implemented")
}

type returnNode struct {
  Expr IExprNode
}

func ReturnNode(expr INode) returnNode {
  return returnNode { expr.(IExprNode) }
}

func (self returnNode) String() string {
  return fmt.Sprintf("%s", self.Expr)
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

type whileNode struct {
  Cond IExprNode
  Body IStmtNode
}

func WhileNode(cond INode, body INode) whileNode {
  return whileNode { cond.(IExprNode), body.(IStmtNode) }
}

func (self whileNode) String() string {
  return fmt.Sprintf("(let while-loop ((while-cond %s)) (if while-cond (begin %s (while-loop %s))))", self.Cond, self.Body, self.Cond)
}
