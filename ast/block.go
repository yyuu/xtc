package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

// BlockNode
type BlockNode struct {
  ClassName string
  Location core.Location
  Variables []*entity.DefinedVariable
  Stmts []core.IStmtNode
  scope *entity.LocalScope
}

func NewBlockNode(loc core.Location, variables []*entity.DefinedVariable, stmts []core.IStmtNode) *BlockNode {
  return &BlockNode { "ast.BlockNode", loc, variables, stmts, nil }
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

func (self BlockNode) GetLocation() core.Location {
  return self.Location
}

func (self BlockNode) GetVariables() []*entity.DefinedVariable {
  return self.Variables
}

func (self BlockNode) GetStmts() []core.IStmtNode {
  return self.Stmts
}

func (self BlockNode) GetScope() *entity.LocalScope {
  return self.scope
}

func (self *BlockNode) SetScope(scope *entity.LocalScope) {
  self.scope = scope
}
