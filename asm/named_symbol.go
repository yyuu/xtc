package asm

type NamedSymbol struct {
  ClassName string
  Name string
}

func NewNamedSymbol(name string) NamedSymbol {
  return NamedSymbol { "asm.NamedSymbol", name }
}

func (self NamedSymbol) IsZero() bool {
  return false
}

func (self NamedSymbol) GetName() string {
  return self.Name
}
