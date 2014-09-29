package sysdep

import (
  "fmt"
)

const (
  X86_AX = iota
  X86_BX
  X86_CX
  X86_DX
  X86_SI
  X86_DI
  X86_SP
  X86_BP
)

type X86Register struct {
  Class int
  Type int
}

func NewX86Register(klass int, t int) *X86Register {
  return &X86Register { klass, t }
}

func (self X86Register) IsRegister() bool {
  return true
}

func (self X86Register) IsMemoryReference() bool {
  return false
}

func (self X86Register) GetBaseName() string {
  switch self.Class {
    case X86_AX: return "ax"
    case X86_BX: return "bx"
    case X86_CX: return "cx"
    case X86_DX: return "dx"
    case X86_SI: return "si"
    case X86_DI: return "di"
    case X86_SP: return "sp"
    case X86_BP: return "bp"
    default: {
      panic(fmt.Errorf("unknown register class: %d", self.Class))
    }
  }
}
