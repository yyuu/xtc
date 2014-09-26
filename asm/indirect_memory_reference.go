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

func NewIndirectMemoryReference(offset core.ILiteral, base core.IRegister) *IndirectMemoryReference {
  return &IndirectMemoryReference { "asm.IndirectMemoryReference", offset, base, true }
}

func (self IndirectMemoryReference) IsRegister() bool {
  return false
}

func (self IndirectMemoryReference) IsMemoryReference() bool {
  return true
}

