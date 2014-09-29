package asm

type UnnamedSymbol struct {
  ClassName string
}

func NewUnnamedSymbol() *UnnamedSymbol {
  return &UnnamedSymbol { "asm.UnnamedSymbol" }
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
