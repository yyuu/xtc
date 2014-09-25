package asm

type ImmediateValue struct {
  ClassName string
  Expr ILiteral
}

func NewImmediateValue(val ILiteral) ImmediateValue {
  return ImmediateValue { "asm.ImmediateValue", val }
}

func (self ImmediateValue) GetExpr() ILiteral {
  return self.Expr
}
