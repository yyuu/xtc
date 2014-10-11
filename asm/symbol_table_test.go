package asm

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestSymbolTableNewSymbol(t *testing.T) {
  table1 := NewSymbolTable("L")
  table2 := NewSymbolTable("LC")
  sym1 := table1.NewSymbol()
// table1
  xt.AssertNotEquals(t, "NewSymbol should always create new symbol", sym1, table1.NewSymbol())
  xt.AssertNotEquals(t, "NewSymbol should always create new symbol", sym1, table1.NewSymbol())
  xt.AssertNotEquals(t, "NewSymbol should always create new symbol", sym1, table1.NewSymbol())
// table2
  xt.AssertNotEquals(t, "NewSymbol should always create new symbol", sym1, table2.NewSymbol())
  xt.AssertNotEquals(t, "NewSymbol should always create new symbol", sym1, table2.NewSymbol())
  xt.AssertNotEquals(t, "NewSymbol should always create new symbol", sym1, table2.NewSymbol())
}

func TestSymbolTableSymbolString(t *testing.T) {
  table1 := NewSymbolTable("L")
  table2 := NewSymbolTable("LC")

  sym1 := NewUnnamedSymbol()
  sym2 := NewUnnamedSymbol()
  sym3 := NewUnnamedSymbol()
// table1
  xt.AssertEquals(t, "SymbolString should always returns L0 for sym1", table1.SymbolString(sym1), "L0")
  xt.AssertEquals(t, "SymbolString should always returns L1 for sym2", table1.SymbolString(sym2), "L1")
  xt.AssertEquals(t, "SymbolString should always returns L2 for sym3", table1.SymbolString(sym3), "L2")
  xt.AssertEquals(t, "SymbolString should always returns L0 for sym1", table1.SymbolString(sym1), "L0")
  xt.AssertEquals(t, "SymbolString should always returns L1 for sym2", table1.SymbolString(sym2), "L1")
  xt.AssertEquals(t, "SymbolString should always returns L2 for sym3", table1.SymbolString(sym3), "L2")
// table2
  xt.AssertEquals(t, "SymbolString should always returns LC0 for sym1", table2.SymbolString(sym1), "LC0")
  xt.AssertEquals(t, "SymbolString should always returns LC1 for sym2", table2.SymbolString(sym2), "LC1")
  xt.AssertEquals(t, "SymbolString should always returns LC2 for sym3", table2.SymbolString(sym3), "LC2")
  xt.AssertEquals(t, "SymbolString should always returns LC0 for sym1", table2.SymbolString(sym1), "LC0")
  xt.AssertEquals(t, "SymbolString should always returns LC1 for sym2", table2.SymbolString(sym2), "LC1")
  xt.AssertEquals(t, "SymbolString should always returns LC2 for sym3", table2.SymbolString(sym3), "LC2")
}

func TestSymbolTableNewString(t *testing.T) {
  table1 := NewSymbolTable("L")
  table2 := NewSymbolTable("LC")
// table1
  xt.AssertEquals(t, "newString should return unique string", table1.newString(), "L0")
  xt.AssertEquals(t, "newString should return unique string", table1.newString(), "L1")
  xt.AssertEquals(t, "newString should return unique string", table1.newString(), "L2")
// table2
  xt.AssertEquals(t, "newString should return unique string", table2.newString(), "LC0")
  xt.AssertEquals(t, "newString should return unique string", table2.newString(), "LC1")
  xt.AssertEquals(t, "newString should return unique string", table2.newString(), "LC2")
}
