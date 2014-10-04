package asm

import (
  "strconv"
  "bitbucket.org/yyuu/xtc/core"
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

var DUMMY_SYMBOL_BASE = "L"
var DummySymbolTable = NewSymbolTable(DUMMY_SYMBOL_BASE)

func (self *SymbolTable) AsSymbolTable() core.ISymbolTable {
  return self
}

func (self *SymbolTable) NewSymbol() core.ISymbol {
  return NewNamedSymbol(self.newString())
}

func (self *SymbolTable) SymbolString(sym core.ISymbol) string {
  unnamed := sym.(*UnnamedSymbol)
  s, ok := self.table[unnamed]
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
