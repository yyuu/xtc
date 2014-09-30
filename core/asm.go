package core

type IOperand interface {
  AsOperand() IOperand
  IsRegister() bool
  IsMemoryReference() bool
  CollectStatistics(IStatistics)
  ToSource(ISymbolTable) string
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
  ToSource(ISymbolTable) string
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
  ToSource(ISymbolTable) string
}

type ISymbolTable interface {
  AsSymbolTable() ISymbolTable
  SymbolString(ISymbol) string
}

type IStatistics interface {
  AsStatistics() IStatistics
  InstructionUsed(string)
  SymbolUsed(ISymbol)
  RegisterUsed(IRegister)
}
