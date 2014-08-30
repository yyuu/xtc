package ast

import (
  "fmt"
  "strings"
)

// BlockNode
type blockNode struct {
  location Location
// Variables []DefinedVariable
  Variables []IExprNode
  Stmts []IStmtNode
}

func NewBlockNode(location Location, variables []IExprNode, stmts []IStmtNode) blockNode {
  return blockNode { location, variables, stmts }
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

func (self blockNode) IsStmt() bool {
  return true
}

func (self blockNode) GetLocation() Location {
  return self.location
}
