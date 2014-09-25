package ir

import (
  "bitbucket.org/yyuu/bs/asm"
)

type Case struct {
  ClassName string
  Value int64
  Label asm.Label
}

func NewCase(value int64, label asm.Label) *Case {
  return &Case { "ir.Case", value, label }
}
