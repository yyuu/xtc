package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type AbsoluteAddress struct {
  ClassName string
  Register core.IRegister
}

func NewAbsoluteAddress(reg core.IRegister) *AbsoluteAddress {
  return &AbsoluteAddress { "asm.AbsoluteAddress", reg }
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
