package asm

import (
  "strconv"
  "bitbucket.org/yyuu/bs/core"
)

type SymbolTable struct {
  ClassName string
  Base string
  Seq int
  table map[core.ISymbol]string
}

func NewSymbolTable(base string) *SymbolTable {
  table := make(map[core.ISymbol]string)
  return &SymbolTable { "asm.SymbolTable", base, 0, table }
}

func (self *SymbolTable) NewSymbol() core.ISymbol {
  return NewNamedSymbol(self.newString())
}

func (self *SymbolTable) newSymbolString(sym *UnnamedSymbol) string {
  s, ok := self.table[sym]
  if ! ok {
    s = self.newString()
    self.table[sym] = s
  }
  return s
}

func (self *SymbolTable) newString() string {
  self.Seq++
  return self.Base + strconv.Itoa(self.Seq)
}
