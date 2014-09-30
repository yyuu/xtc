package sysdep

import (
  bs_asm "bitbucket.org/yyuu/bs/asm"
)

type x86VirtualStack struct {
  naturalType int
  offset int64
  max int64
  memrefs []*bs_asm.IndirectMemoryReference
}

func newX86VirtualStack(naturalType int) *x86VirtualStack {
  memrefs := []*bs_asm.IndirectMemoryReference { }
  return &x86VirtualStack { naturalType, 0, 0, memrefs }
}

func (self *x86VirtualStack) reset() {
  self.offset = 0
  self.max = 0
  self.memrefs = []*bs_asm.IndirectMemoryReference { }
}

func (self *x86VirtualStack) maxSize() int64 {
  return self.max
}

func (self *x86VirtualStack) extend(n int64) {
  self.offset += n
  if self.max < self.offset {
    self.max = self.offset
  }
}

func (self *x86VirtualStack) rewind(n int64) {
  self.offset -= n
}

func (self *x86VirtualStack) top() *bs_asm.IndirectMemoryReference {
  mem := self.relocatableMem(-self.offset, self.bp())
  self.memrefs = append(self.memrefs, mem)
  return mem
}

func (self *x86VirtualStack) relocatableMem(offset int64, base *x86Register) *bs_asm.IndirectMemoryReference {
  return bs_asm.NewIndirectMemoryReference(bs_asm.NewIntegerLiteral(offset), base, false)
}

func (self *x86VirtualStack) bp() *x86Register {
  return newX86Register(x86_bp, self.naturalType)
}

func (self *x86VirtualStack) fixOffset(diff int64) {
  for i := range self.memrefs {
    self.memrefs[i].FixOffset(diff)
  }
}
