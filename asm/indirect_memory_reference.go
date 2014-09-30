package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type IndirectMemoryReference struct {
  ClassName string
  Offset core.ILiteral
  Base core.IRegister
  Fixed bool
}

func NewIndirectMemoryReference(offset core.ILiteral, base core.IRegister, fixed bool) *IndirectMemoryReference {
  return &IndirectMemoryReference { "asm.IndirectMemoryReference", offset, base, fixed }
}

func (self *IndirectMemoryReference) AsOperand() core.IOperand {
  return self
}

func (self *IndirectMemoryReference) AsMemoryReference() core.IMemoryReference {
  return self
}

func (self IndirectMemoryReference) IsRegister() bool {
  return false
}

func (self IndirectMemoryReference) IsMemoryReference() bool {
  return true
}

func (self *IndirectMemoryReference) FixOffset(diff int64) {
  if self.Fixed {
    panic("must not happedn: fixed = true")
  }
  curr := self.Offset.(*IntegerLiteral).GetValue()
  self.Offset = NewIntegerLiteral(int64(curr + diff))
  self.Fixed = true
}

func (self *IndirectMemoryReference) CollectStatistics(stats core.IStatistics) {
  self.Base.CollectStatistics(stats)
}
