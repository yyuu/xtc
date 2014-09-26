package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type ImmediateValue struct {
  ClassName string
  Expr core.ILiteral
}

func NewImmediateValue(val core.ILiteral) ImmediateValue {
  return ImmediateValue { "asm.ImmediateValue", val }
}

func (self ImmediateValue) GetExpr() core.ILiteral {
  return self.Expr
}
