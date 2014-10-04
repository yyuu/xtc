package ir

import (
  "bitbucket.org/yyuu/xtc/asm"
  "bitbucket.org/yyuu/xtc/core"
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

func (self Switch) GetCond() core.IExpr {
  return self.Cond
}

func (self Switch) GetCases() []*Case {
  return self.Cases
}

func (self Switch) GetDefaultLabel() *asm.Label {
  return self.DefaultLabel
}

func (self Switch) GetEndLabel() *asm.Label {
  return self.EndLabel
}
