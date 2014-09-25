package asm

type Instruction struct {
  ClassName string
  Mnemonic string
  Suffix string
  Operands []IOperand
  NeedRelocation bool
}

func NewInstruction(mnemonic string, suffix string, operands []IOperand, reloc bool) Instruction {
  return Instruction { "asm.Instruction", mnemonic, suffix, operands, reloc }
}

func (self Instruction) IsInstruction() bool {
  return true
}

func (self Instruction) IsJumpInstruction() bool {
  return self.Mnemonic == "jmp" || self.Mnemonic == "jz" || self.Mnemonic == "jne" || self.Mnemonic == "je" || self.Mnemonic == "jne"
}

func (self Instruction) GetMnemonic() string {
  return self.Mnemonic
}

func (self Instruction) NumOperands() int {
  return len(self.Operands)
}

func (self Instruction) Operand1() IOperand {
  return self.Operands[0]
}

func (self Instruction) Operand2() IOperand {
  return self.Operands[1]
}

func (self Instruction) JumpDestination() ISymbol {
  ref := self.Operand1().(DirectMemoryReference)
  return ref.GetValue().(ISymbol)
}
