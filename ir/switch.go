package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type Switch struct {
  ClassName string
  Location core.Location
  Cond core.IExpr
  Cases []*Case
  DefaultLabel string // FIXME
  EndLabel string // FIXME
}

func NewSwitch(loc core.Location, cond core.IExpr, cases []*Case, defaultLabel string, endLabel string) *Switch {
  return &Switch { "ir.Switch", loc, cond, cases, defaultLabel, endLabel }
}

func (self Switch) AsStmt() core.IStmt {
  return self
}

func (self Switch) GetLocation() core.Location {
  return self.Location
}
