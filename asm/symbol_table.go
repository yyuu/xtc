package asm

import (
  "strconv"
)

type SymbolTable struct {
  ClassName string
  Base string
  Map map[ISymbol]string
  Seq int
}

func NewSymbolTable(base string) *SymbolTable {
  return &SymbolTable { "asm.SymbolTable", base, make(map[ISymbol]string), 0 }
}

func (self *SymbolTable) newSymbol() ISymbol {
  return NewNamedSymbol(self.newString())
}

func (self *SymbolTable) newSymbolString(sym *UnnamedSymbol) string {
  s, ok := self.Map[sym]
  if ! ok {
    s = self.newString()
    self.Map[sym] = s
  }
  return s
}

func (self *SymbolTable) newString() string {
  self.Seq++
  return self.Base + strconv.Itoa(self.Seq)
}
