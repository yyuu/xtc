package asm

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type IntegerLiteral struct {
  ClassName string
  Value int64
}

func NewIntegerLiteral(n int64) *IntegerLiteral {
  return &IntegerLiteral { "asm.IntegerLiteral", n }
}

func (self *IntegerLiteral) AsLiteral() core.ILiteral {
  return self
}

func (self IntegerLiteral) IsZero() bool {
  return self.Value == 0
}

func (self IntegerLiteral) GetValue() int64 {
  return self.Value
}

func (self IntegerLiteral) String() string {
  return fmt.Sprintf("%d", self.Value)
}

func (self *IntegerLiteral) CollectStatistics(stats core.IStatistics) {
  // does nothing
}

func (self *IntegerLiteral) ToSource(table core.ISymbolTable) string {
  return fmt.Sprintf("%d", self.Value)
}
