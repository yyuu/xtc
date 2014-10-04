package asm

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

type SuffixedSymbol struct {
  ClassName string
  Base core.ISymbol
  Suffix string
}

func NewSuffixedSymbol(base core.ISymbol, suffix string) *SuffixedSymbol {
  return &SuffixedSymbol { "asm.SuffixedSymbol", base, suffix }
}

func (self *SuffixedSymbol) AsLiteral() core.ILiteral {
  return self
}

func (self *SuffixedSymbol) AsSymbol() core.ISymbol {
  return self
}

func (self SuffixedSymbol) IsZero() bool {
  return false
}

func (self SuffixedSymbol) GetName() string {
  return self.Base.GetName()
}

func (self SuffixedSymbol) String() string {
  return self.Base.String() + self.Suffix
}

func (self *SuffixedSymbol) CollectStatistics(stats core.IStatistics) {
  self.Base.CollectStatistics(stats)
}

func (self *SuffixedSymbol) ToSource(table core.ISymbolTable) string {
  return fmt.Sprintf("%s%s", self.Base.ToSource(table), self.Suffix)
}
