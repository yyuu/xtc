package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type SuffixSymbol struct {
  ClassName string
  Base core.ISymbol
  Suffix string
}

func NewSuffixSymbol(base core.ISymbol, suffix string) SuffixSymbol {
  return SuffixSymbol { "asm.SuffixSymbol", base, suffix }
}

func (self SuffixSymbol) IsZero() bool {
  return false
}

func (self SuffixSymbol) GetName() string {
  return self.Base.GetName()
}
