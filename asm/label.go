package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type Label struct {
  ClassName string
  Symbol core.ISymbol
}

func NewLabel(sym core.ISymbol) *Label {
  return &Label { "asm.Label", sym }
}

func NewUnnamedLabel() *Label {
  return NewLabel(NewUnnamedSymbol())
}

func (self *Label) AsAssembly() core.IAssembly {
  return self
}

func (self Label) IsInstruction() bool {
  return false
}

func (self Label) IsLabel() bool {
  return true
}

func (self Label) IsDirective() bool {
  return false
}

func (self Label) IsComment() bool {
  return false
}

func (self Label) GetSymbol() core.ISymbol {
  return self.Symbol
}

func (self *Label) CollectStatistics(stats core.IStatistics) {
  // does nothing by default
}
