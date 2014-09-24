package ir

import (
  "bitbucket.org/yyuu/bs/core"
)

type CJump struct {
  ClassName string
  Location core.Location
  Cond core.IExpr
  ThenLabel string // FIXME:
  ElseLabel string // FIXME:
}

func NewCJump(loc core.Location, cond core.IExpr, thenLabel string, elseLabel string) *CJump {
  return &CJump { "ir.CJump", loc, cond, thenLabel, elseLabel }
}

func (self CJump) AsStmt() core.IStmt {
  return self
}

func (self CJump) GetLocation() core.Location {
  return self.Location
}
