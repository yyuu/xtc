package asm

type AbsoluteAddress struct {
  ClassName string
  Register IRegister
}

func NewAbsoluteAddress(reg IRegister) AbsoluteAddress {
  return AbsoluteAddress { "asm.AbsoluteAddress", reg }
}

func (self AbsoluteAddress) GetRegister() IOperand {
  return self.Register
}
