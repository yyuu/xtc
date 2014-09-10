package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// BlockNode
type BlockNode struct {
  ClassName string
  Location duck.ILocation
  Variables []duck.IDefinedVariable
  Stmts []duck.IStmtNode
  scope duck.IVariableScope
}

func NewBlockNode(loc duck.ILocation, variables []duck.IDefinedVariable, stmts []duck.IStmtNode) BlockNode {
  if loc == nil { panic("location is nil") }
  return BlockNode { "ast.BlockNode", loc, variables, stmts, nil }
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

func (self BlockNode) IsStmtNode() bool {
  return true
}

func (self BlockNode) GetLocation() duck.ILocation {
  return self.Location
}

func (self BlockNode) GetVariables() []duck.IDefinedVariable {
  return self.Variables
}

func (self BlockNode) GetStmts() []duck.IStmtNode {
  return self.Stmts
}

func (self BlockNode) GetScope() duck.IVariableScope {
  return self.scope
}

func (self *BlockNode) SetScope(scope duck.IVariableScope) {
  self.scope = scope
}
