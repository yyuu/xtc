package asm

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

type ImmediateValue struct {
  ClassName string
  Expr core.ILiteral
}

func NewImmediateValue(val core.ILiteral) *ImmediateValue {
  return &ImmediateValue { "asm.ImmediateValue", val }
}

func (self *ImmediateValue) AsOperand() core.IOperand {
  return self
}

func (self *ImmediateValue) AsImmediateValue() core.IImmediateValue {
  return self
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

func (self *ImmediateValue) CollectStatistics(stats core.IStatistics) {
  // does nothing
}

func (self *ImmediateValue) ToSource(table core.ISymbolTable) string {
  return fmt.Sprintf("$%s", self.Expr.ToSource(table))
}
