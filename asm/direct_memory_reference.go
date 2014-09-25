package asm

type DirectMemoryReference struct {
  ClassName string
  Value ILiteral
}

func NewDirectMemoryReference(val ILiteral) DirectMemoryReference {
  return DirectMemoryReference { "asm.DirectMemoryReference", val}
}

func (self DirectMemoryReference) IsRegister() bool {
  return false
}

func (self DirectMemoryReference) IsMemoryReference() bool {
  return true
}

func (self DirectMemoryReference) GetValue() ILiteral {
  return self.Value
}
