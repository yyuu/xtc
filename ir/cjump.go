package ir

import (
  "bitbucket.org/yyuu/xtc/asm"
  "bitbucket.org/yyuu/xtc/core"
)

type CJump struct {
  ClassName string
  Location core.Location
  Cond core.IExpr
  ThenLabel *asm.Label
  ElseLabel *asm.Label
}

func NewCJump(loc core.Location, cond core.IExpr, thenLabel *asm.Label, elseLabel *asm.Label) *CJump {
  return &CJump { "ir.CJump", loc, cond, thenLabel, elseLabel }
}

func (self *CJump) AsStmt() core.IStmt {
  return self
}

func (self CJump) GetLocation() core.Location {
  return self.Location
}

func (self CJump) GetCond() core.IExpr {
  return self.Cond
}

func (self CJump) GetThenLabel() *asm.Label {
  return self.ThenLabel
}

func (self CJump) GetElseLabel() *asm.Label {
  return self.ElseLabel
}
