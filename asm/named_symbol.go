package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type NamedSymbol struct {
  ClassName string
  Name string
}

func NewNamedSymbol(name string) *NamedSymbol {
  return &NamedSymbol { "asm.NamedSymbol", name }
}

func (self *NamedSymbol) AsLiteral() core.ILiteral {
  return self
}

func (self *NamedSymbol) AsSymbol() core.ISymbol {
  return self
}

func (self NamedSymbol) IsZero() bool {
  return false
}

func (self NamedSymbol) GetName() string {
  return self.Name
}

func (self NamedSymbol) String() string {
  return self.Name
}

func (self *NamedSymbol) CollectStatistics(stats core.IStatistics) {
  stats.SymbolUsed(self)
}
