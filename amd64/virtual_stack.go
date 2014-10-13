package amd64

import (
  xtc_asm "bitbucket.org/yyuu/xtc/asm"
)

type VirtualStack struct {
  naturalType int
  offset int64
  max int64
  memrefs []*xtc_asm.IndirectMemoryReference
}

func NewVirtualStack(naturalType int) *VirtualStack {
  memrefs := []*xtc_asm.IndirectMemoryReference { }
  return &VirtualStack { naturalType, 0, 0, memrefs }
}

func (self *VirtualStack) Reset() {
  self.offset = 0
  self.max = 0
  self.memrefs = []*xtc_asm.IndirectMemoryReference { }
}

func (self *VirtualStack) MaxSize() int64 {
  return self.max
}

func (self *VirtualStack) Extend(n int64) {
  self.offset += n
  if self.max < self.offset {
    self.max = self.offset
  }
}

func (self *VirtualStack) Rewind(n int64) {
  self.offset -= n
}

func (self *VirtualStack) Top() *xtc_asm.IndirectMemoryReference {
  mem := self.relocatableMem(-self.offset, self.bp())
  self.memrefs = append(self.memrefs, mem)
  return mem
}

func (self *VirtualStack) relocatableMem(offset int64, base *Register) *xtc_asm.IndirectMemoryReference {
  return xtc_asm.NewIndirectMemoryReference(xtc_asm.NewIntegerLiteral(offset), base, false)
}

func (self *VirtualStack) bp() *Register {
  return NewRegister(BP, self.naturalType)
}

func (self *VirtualStack) FixOffset(diff int64) {
  for i := range self.memrefs {
    self.memrefs[i].FixOffset(diff)
  }
}
