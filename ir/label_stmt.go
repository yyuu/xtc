package ir

import (
  "bitbucket.org/yyuu/bs/asm"
  "bitbucket.org/yyuu/bs/core"
)

type LabelStmt struct {
  ClassName string
  Location core.Location
  Label *asm.Label
}

func NewLabelStmt(loc core.Location, label *asm.Label) *LabelStmt {
  return &LabelStmt { "ir.LabelStmt", loc, label }
}

func (self *LabelStmt) AsStmt() core.IStmt {
  return self
}

func (self LabelStmt) GetLocation() core.Location {
  return self.Location
}

func (self LabelStmt) GetLabel() *asm.Label {
  return self.Label
}
