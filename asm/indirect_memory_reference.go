package asm

import (
  "fmt"
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

func (self *IndirectMemoryReference) ToSource(table core.ISymbolTable) string {
  if ! self.Fixed {
    panic("must not happen: writing unfixed variable")
  } else {
    if self.Offset.IsZero() {
      return fmt.Sprintf("%s(%s)", self.Offset.ToSource(table), self.Base.ToSource(table))
    } else {
      return fmt.Sprintf("(%s)", self.Base.ToSource(table))
    }
  }
}

func (self *IndirectMemoryReference) String() string {
  return self.ToSource(DummySymbolTable)
}
