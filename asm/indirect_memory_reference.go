package asm

type IndirectMemoryReference struct {
  ClassName string
  Offset ILiteral
  Base IRegister
  Fixed bool
}

func NewIndirectMemoryReference(offset ILiteral, base IRegister) IndirectMemoryReference {
  return IndirectMemoryReference { "asm.IndirectMemoryReference", offset, base, true }
}

func (self IndirectMemoryReference) IsRegister() bool {
  return false
}

func (self IndirectMemoryReference) IsMemoryReference() bool {
  return true
}

