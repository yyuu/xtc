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
    sVariables[i] = fmt.Sprintf("%s", self.Variables[i])
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
