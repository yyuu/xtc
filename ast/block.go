package ast

import (
  "encoding/json"
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// BlockNode
type BlockNode struct {
  Location duck.ILocation
// Variables []DefinedVariable
  Variables []duck.IExprNode
  Stmts []duck.IStmtNode
}

func NewBlockNode(location duck.ILocation, variables []duck.IExprNode, stmts []duck.IStmtNode) BlockNode {
  return BlockNode { location, variables, stmts }
}

func (self BlockNode) String() string {
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

func (self BlockNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Variables []duck.IExprNode
    Stmts []duck.IStmtNode
  }
  x.ClassName = "ast.BlockNode"
  x.Location = self.Location
  x.Variables = self.Variables
  x.Stmts = self.Stmts
  return json.Marshal(x)
}

func (self BlockNode) IsStmt() bool {
  return true
}

func (self BlockNode) GetLocation() duck.ILocation {
  return self.Location
}
