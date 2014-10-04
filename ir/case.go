package ir

import (
  "bitbucket.org/yyuu/xtc/asm"
)

type Case struct {
  ClassName string
  Value int64
  Label *asm.Label
}

func NewCase(value int64, label *asm.Label) *Case {
  return &Case { "ir.Case", value, label }
}

func (self Case) GetValue() int64 {
  return self.Value
}

func (self Case) GetLabel() *asm.Label {
  return self.Label
}
