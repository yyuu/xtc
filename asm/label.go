package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type Label struct {
  ClassName string
  Symbol core.ISymbol
}

func NewLabel(sym core.ISymbol) Label {
  return Label { "asm.Label", sym }
}

func NewUnnamedLabel() Label {
  return NewLabel(NewUnnamedSymbol())
}

func (self Label) IsLabel() bool {
  return true
}

func (self Label) GetSymbol() core.ISymbol {
  return self.Symbol
}
