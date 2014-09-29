package core

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

type IImmediateValue interface {
  IOperand
}

type ILiteral interface {
  String() string
  IsZero() bool
}

type ISymbol interface {
  ILiteral
  GetName() string
}

type IAssembly interface {
  IsInstruction() bool
  IsLabel() bool
  IsDirective() bool
  IsComment() bool
}
