package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type LabelStmt struct {
  ClassName string
  Location core.Location
  Label string // FIXME:
}

func NewLabelStmt(loc core.Location, label string) *LabelStmt {
  return &LabelStmt { "ir.LabelStmt", loc, label }
}

func (self LabelStmt) AsStmt() core.IStmt {
  return self
}

func (self LabelStmt) GetLocation() core.Location {
  return self.Location
}
