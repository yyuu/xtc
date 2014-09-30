package sysdep

import (
  "fmt"
  bs_core "bitbucket.org/yyuu/bs/core"
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
  TypeId int
}

func newX86Register(klass int, t int) *x86Register {
  return &x86Register { klass, t }
}

func (self *x86Register) AsOperand() bs_core.IOperand {
  return self
}

func (self *x86Register) AsRegister() bs_core.IRegister {
  return self
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

func (self x86Register) GetTypeId() int {
  return self.TypeId
}

func (self x86Register) ForType(t int) *x86Register {
  return newX86Register(self.Class, t)
}

func (self *x86Register) CollectStatistics(stats bs_core.IStatistics) {
  stats.RegisterUsed(self)
}
