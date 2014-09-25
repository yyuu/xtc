package asm

type SuffixSymbol struct {
  ClassName string
  Base ISymbol
  Suffix string
}

func NewSuffixSymbol(base ISymbol, suffix string) SuffixSymbol {
  return SuffixSymbol { "asm.SuffixSymbol", base, suffix }
}

func (self SuffixSymbol) IsZero() bool {
  return false
}

func (self SuffixSymbol) GetName() string {
  return self.Base.GetName()
}
