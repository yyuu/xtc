package entity

import (
  "bitbucket.org/yyuu/xtc/core"
)

type ConstantEntry struct {
  value string
  symbol core.ISymbol
  memref core.IMemoryReference
  address core.IOperand
}

func NewConstantEntry(value string) *ConstantEntry {
  return &ConstantEntry {
    value: value,
    symbol: nil,
    memref: nil,
    address: nil,
  }
}

func (self *ConstantEntry) GetValue() string {
  return self.value
}

func (self *ConstantEntry) GetSymbol() core.ISymbol {
  return self.symbol
}

func (self *ConstantEntry) SetSymbol(symbol core.ISymbol) {
  self.symbol = symbol
}

func (self *ConstantEntry) GetMemref() core.IMemoryReference {
  return self.memref
}

func (self *ConstantEntry) SetMemref(memref core.IMemoryReference) {
  self.memref = memref
}

func (self *ConstantEntry) GetAddress() core.IOperand {
  return self.address
}

func (self *ConstantEntry) SetAddress(address core.IOperand) {
  self.address = address
}
