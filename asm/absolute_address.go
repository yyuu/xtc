package asm

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type AbsoluteAddress struct {
  ClassName string
  Register core.IRegister
}

func NewAbsoluteAddress(reg core.IRegister) *AbsoluteAddress {
  return &AbsoluteAddress { "asm.AbsoluteAddress", reg }
}

func (self *AbsoluteAddress) AsOperand() core.IOperand {
  return self
}

func (self AbsoluteAddress) IsRegister() bool {
  return false
}

func (self AbsoluteAddress) IsMemoryReference() bool {
  return false
}

func (self AbsoluteAddress) GetRegister() core.IOperand {
  return self.Register
}

func (self *AbsoluteAddress) CollectStatistics(stats core.IStatistics) {
  self.Register.CollectStatistics(stats)
}

func (self *AbsoluteAddress) ToSource(table core.ISymbolTable) string {
  return fmt.Sprintf("*%s", self.Register.ToSource(table))
}
