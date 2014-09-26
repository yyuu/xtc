package asm

type IntegerLiteral struct {
  ClassName string
  Value int64
}

func NewIntegerLiteral(n int64) *IntegerLiteral {
  return &IntegerLiteral { "asm.IntegerLiteral", n }
}

func (self IntegerLiteral) IsZero() bool {
  return self.Value == 0
}

func (self IntegerLiteral) GetValue() int64 {
  return self.Value
}
