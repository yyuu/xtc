package asm

import (
  "bitbucket.org/yyuu/xtc/core"
)

type UnnamedSymbol struct {
  ClassName string
}

func NewUnnamedSymbol() *UnnamedSymbol {
  return &UnnamedSymbol { "asm.UnnamedSymbol" }
}

func (self *UnnamedSymbol) AsLiteral() core.ILiteral {
  return self
}

func (self *UnnamedSymbol) AsSymbol() core.ISymbol {
  return self
}

func (self UnnamedSymbol) IsZero() bool {
  return false
}

func (self UnnamedSymbol) GetName() string {
  panic("unnamed symbol")
}

func (self UnnamedSymbol) String() string {
  panic("UnnamedSymbol#String() called")
}

func (self *UnnamedSymbol) CollectStatistics(stats core.IStatistics) {
  stats.SymbolUsed(self)
}

func (self *UnnamedSymbol) ToSource(table core.ISymbolTable) string {
  return table.SymbolString(self)
}
