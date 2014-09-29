package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type ImmediateValue struct {
  ClassName string
  Expr core.ILiteral
}

func NewImmediateValue(val core.ILiteral) *ImmediateValue {
  return &ImmediateValue { "asm.ImmediateValue", val }
}

func (self ImmediateValue) IsRegister() bool {
  return false
}

func (self ImmediateValue) IsMemoryReference() bool {
  return false
}

func (self ImmediateValue) GetExpr() core.ILiteral {
  return self.Expr
}
