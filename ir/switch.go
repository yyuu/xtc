package ir

import (
  "bitbucket.org/yyuu/bs/asm"
  "bitbucket.org/yyuu/bs/core"
)

type Switch struct {
  ClassName string
  Location core.Location
  Cond core.IExpr
  Cases []*Case
  DefaultLabel *asm.Label
  EndLabel *asm.Label
}

func NewSwitch(loc core.Location, cond core.IExpr, cases []*Case, defaultLabel *asm.Label, endLabel *asm.Label) *Switch {
  return &Switch { "ir.Switch", loc, cond, cases, defaultLabel, endLabel }
}

func (self *Switch) AsStmt() core.IStmt {
  return self
}

func (self Switch) GetLocation() core.Location {
  return self.Location
}
