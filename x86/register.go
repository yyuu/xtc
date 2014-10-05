package x86

import (
  "fmt"
  xtc_asm "bitbucket.org/yyuu/xtc/asm"
  xtc_core "bitbucket.org/yyuu/xtc/core"
)

const (
  AX = iota
  BX
  CX
  DX
  SI
  DI
  SP
  BP
)

type Register struct {
  Class int
  TypeId int
}

func NewRegister(klass int, t int) *Register {
  return &Register { klass, t }
}

func (self *Register) AsOperand() xtc_core.IOperand {
  return self
}

func (self *Register) AsRegister() xtc_core.IRegister {
  return self
}

func (self Register) IsRegister() bool {
  return true
}

func (self Register) IsMemoryReference() bool {
  return false
}

func (self Register) GetBaseName() string {
  switch self.Class {
    case AX: return "ax"
    case BX: return "bx"
    case CX: return "cx"
    case DX: return "dx"
    case SI: return "si"
    case DI: return "di"
    case SP: return "sp"
    case BP: return "bp"
    default: {
      panic(fmt.Errorf("unknown register class: %d", self.Class))
    }
  }
}

func (self Register) GetTypeId() int {
  return self.TypeId
}

func (self Register) ForType(t int) *Register {
  return NewRegister(self.Class, t)
}

func (self *Register) CollectStatistics(stats xtc_core.IStatistics) {
  stats.RegisterUsed(self)
}

func (self *Register) ToSource(table xtc_core.ISymbolTable) string {
  return fmt.Sprintf("%%%s", self.GetTypedName())
}

func (self *Register) GetTypedName() string {
  switch self.TypeId {
    case xtc_asm.TYPE_INT8:  return self.lowerByteRegister()
    case xtc_asm.TYPE_INT16: return self.GetBaseName()
    case xtc_asm.TYPE_INT32: return fmt.Sprintf("e%s", self.GetBaseName())
    case xtc_asm.TYPE_INT64: return fmt.Sprintf("r%s", self.GetBaseName())
    default: {
      panic(fmt.Errorf("unknown register type: %d", self.TypeId))
    }
  }
}

func (self *Register) lowerByteRegister() string {
  switch self.Class {
    case AX: return "al"
    case BX: return "bl"
    case CX: return "cl"
    case DX: return "dl"
    default: {
      panic(fmt.Errorf("does not have lower-byte register: %d", self.Class))
    }
  }
}

func (self Register) String() string {
  return fmt.Sprintf("%%%s", self.GetTypedName())
}
