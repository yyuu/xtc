package core

type IOperand interface {
  AsOperand() IOperand
  IsRegister() bool
  IsMemoryReference() bool
  CollectStatistics(IStatistics)
}

type IRegister interface {
  IOperand
  AsRegister() IRegister
}

type IMemoryReference interface {
  IOperand
  AsMemoryReference() IMemoryReference
  FixOffset(int64)
}

type IImmediateValue interface {
  AsImmediateValue() IImmediateValue
  IOperand
}

type ILiteral interface {
  AsLiteral() ILiteral
  String() string
  IsZero() bool
  CollectStatistics(IStatistics)
}

type ISymbol interface {
  ILiteral
  AsSymbol() ISymbol
  GetName() string
}

type IAssembly interface {
  AsAssembly() IAssembly
  IsInstruction() bool
  IsLabel() bool
  IsDirective() bool
  IsComment() bool
  CollectStatistics(IStatistics)
}

type IStatistics interface {
  AsStatistics() IStatistics
  InstructionUsed(string)
  SymbolUsed(ISymbol)
  RegisterUsed(IRegister)
}
