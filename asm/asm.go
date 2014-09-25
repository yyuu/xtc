package asm

type IOperand interface {
  IsRegister() bool
  IsMemoryReference() bool
}

type IRegister interface {
  IOperand
}

type IMemoryReference interface {
  IOperand
}

type ILiteral interface {
  IsZero() bool
}

type ISymbol interface {
  ILiteral
  GetName() string
}

type IBaseSymbol interface {
  ISymbol
}

type IAssembly interface {
  IsInstruction() bool
  IsLabel() bool
  IsDirective() bool
  IsComment() bool
}
