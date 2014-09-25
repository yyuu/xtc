package asm

type Label struct {
  ClassName string
  Symbol ISymbol
}

func NewLabel(sym ISymbol) Label {
  return Label { "asm.Label", sym }
}

func NewUnnamedLabel() Label {
  return NewLabel(NewUnnamedSymbol())
}

func (self Label) IsLabel() bool {
  return true
}

func (self Label) GetSymbol() ISymbol {
  return self.Symbol
}
