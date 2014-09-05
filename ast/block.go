package ast

import (
  "encoding/json"
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// BlockNode
type BlockNode struct {
  location duck.ILocation
  variables []duck.IDefinedVariable
  stmts []duck.IStmtNode
}

func NewBlockNode(loc duck.ILocation, variables []duck.IDefinedVariable, stmts []duck.IStmtNode) BlockNode {
  return BlockNode { loc, variables, stmts }
}

func (self BlockNode) String() string {
  sVariables := make([]string, len(self.variables))
  for i := range self.variables {
    sVariables[i] = fmt.Sprintf("%s", self.variables[i])
  }
  sStmts := make([]string, len(self.stmts))
  for j := range self.stmts {
    sStmts[j] = fmt.Sprintf("%s", self.stmts[j])
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

func (self BlockNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Variables []duck.IDefinedVariable
    Stmts []duck.IStmtNode
  }
  x.ClassName = "ast.BlockNode"
  x.Location = self.location
  x.Variables = self.variables
  x.Stmts = self.stmts
  return json.Marshal(x)
}

func (self BlockNode) IsStmt() bool {
  return true
}

func (self BlockNode) GetLocation() duck.ILocation {
  return self.location
}

func (self BlockNode) GetVariables() []duck.IDefinedVariable {
  return self.variables
}

func (self BlockNode) GetStmts() []duck.IStmtNode {
  return self.stmts
}
