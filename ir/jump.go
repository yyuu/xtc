package ir

import (
  "bitbucket.org/yyuu/bs/asm"
  "bitbucket.org/yyuu/bs/core"
)

type Jump struct {
  ClassName string
  Location core.Location
  Label *asm.Label
}

func NewJump(loc core.Location, label *asm.Label) *Jump {
  return &Jump { "ir.Jump", loc, label }
}

func (self Jump) AsStmt() core.IStmt {
  return self
}

func (self Jump) GetLocation() core.Location {
  return self.Location
}
