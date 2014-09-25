package asm

import (
  "strconv"
)

type SymbolTable struct {
  ClassName string
  Base string
  Seq int
}

func NewSymbolTable(base string) *SymbolTable {
  return &SymbolTable { "asm.SymbolTable", base, 0 }
}

func (self *SymbolTable) newSymbol() ISymbol {
  return NewNamedSymbol(self.newString())
}

//func (self *SymbolTable) newSymbolString(sym UnnamedSymbol) string {
//  s := self.Map[sym.Key()]
//  if s != nil {
//    return s
//  } else {
//    self.Map[sym.Key()] = sym
//  }
//}

func (self *SymbolTable) newString() string {
  self.Seq++
  return self.Base + strconv.Itoa(self.Seq)
}
