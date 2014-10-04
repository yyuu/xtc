package asm

import (
  "bitbucket.org/yyuu/xtc/core"
)

type DirectMemoryReference struct {
  ClassName string
  Value core.ILiteral
}

func NewDirectMemoryReference(val core.ILiteral) *DirectMemoryReference {
  return &DirectMemoryReference { "asm.DirectMemoryReference", val}
}

func (self *DirectMemoryReference) AsOperand() core.IOperand {
  return self
}

func (self *DirectMemoryReference) AsMemoryReference() core.IMemoryReference {
  return self
}

func (self DirectMemoryReference) IsRegister() bool {
  return false
}

func (self DirectMemoryReference) IsMemoryReference() bool {
  return true
}

func (self DirectMemoryReference) GetValue() core.ILiteral {
  return self.Value
}

func (self *DirectMemoryReference) FixOffset(n int64) {
  panic("#FixOffset called")
}

func (self *DirectMemoryReference) CollectStatistics(stats core.IStatistics) {
  self.Value.CollectStatistics(stats)
}

func (self *DirectMemoryReference) ToSource(table core.ISymbolTable) string {
  return self.Value.ToSource(table)
}

func (self *DirectMemoryReference) String() string {
  return self.ToSource(DummySymbolTable)
}
