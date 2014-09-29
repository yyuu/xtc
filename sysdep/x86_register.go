package sysdep

import (
  "fmt"
)

const (
  x86_ax = iota
  x86_bx
  x86_cx
  x86_dx
  x86_si
  x86_di
  x86_sp
  x86_bp
)

type x86Register struct {
  Class int
  Type int
}

func newX86Register(klass int, t int) *x86Register {
  return &x86Register { klass, t }
}

func (self x86Register) IsRegister() bool {
  return true
}

func (self x86Register) IsMemoryReference() bool {
  return false
}

func (self x86Register) GetBaseName() string {
  switch self.Class {
    case x86_ax: return "ax"
    case x86_bx: return "bx"
    case x86_cx: return "cx"
    case x86_dx: return "dx"
    case x86_si: return "si"
    case x86_di: return "di"
    case x86_sp: return "sp"
    case x86_bp: return "bp"
    default: {
      panic(fmt.Errorf("unknown register class: %d", self.Class))
    }
  }
}

func (self x86Register) GetType() int {
  return self.Type
}
