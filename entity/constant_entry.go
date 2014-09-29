package entity

import (
  "bitbucket.org/yyuu/bs/core"
)

type ConstantEntry struct {
  value string
  symbol core.ISymbol
  memref core.IMemoryReference
  address core.IImmediateValue
}

func NewConstantEntry(value string) *ConstantEntry {
  return &ConstantEntry { value, nil, nil, nil }
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

func (self *ConstantEntry) GetAddress() core.IImmediateValue {
  return self.address
}

func (self *ConstantEntry) SetAddress(address core.IImmediateValue) {
  self.address = address
}
