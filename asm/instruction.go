package asm

import (
  "bitbucket.org/yyuu/bs/core"
)

type Instruction struct {
  ClassName string
  Mnemonic string
  Suffix string
  Operands []core.IOperand
  NeedRelocation bool
}

func NewInstruction(mnemonic string, suffix string, operands []core.IOperand, reloc bool) *Instruction {
  return &Instruction { "asm.Instruction", mnemonic, suffix, operands, reloc }
}

func (self *Instruction) AsAssembly() core.IAssembly {
  return self
}

func (self Instruction) IsInstruction() bool {
  return true
}

func (self Instruction) IsLabel() bool {
  return false
}

func (self Instruction) IsDirective() bool {
  return false
}

func (self Instruction) IsComment() bool {
  return false
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

func (self Instruction) Operand1() core.IOperand {
  return self.Operands[0]
}

func (self Instruction) Operand2() core.IOperand {
  return self.Operands[1]
}

func (self Instruction) JumpDestination() core.ISymbol {
  ref := self.Operand1().(*DirectMemoryReference)
  return ref.GetValue().(core.ISymbol)
}

func (self *Instruction) CollectStatistics(stats core.IStatistics) {
  stats.InstructionUsed(self.Mnemonic)
  for i := range self.Operands {
    self.Operands[i].CollectStatistics(stats)
  }
}
