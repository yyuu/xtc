package ir

import (
  "bitbucket.org/yyuu/bs/asm"
  "bitbucket.org/yyuu/bs/core"
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

func (self CJump) AsStmt() core.IStmt {
  return self
}

func (self CJump) GetLocation() core.Location {
  return self.Location
}
