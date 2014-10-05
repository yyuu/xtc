package core

type IEntity interface {
  String() string
  GetName() string
  IsDefined() bool
  IsPrivate() bool
  IsConstant() bool
  IsParameter() bool
  IsVariable() bool
  Refered()
  IsRefered() bool
  GetNumRefered() int
  GetTypeNode() ITypeNode
  GetTypeRef() ITypeRef
  GetType() IType
  GetValue() IExprNode
  SymbolString() string
  GetMemref() IMemoryReference
  SetMemref(IMemoryReference)
  GetAddress() IOperand
  SetAddress(IOperand)
}

type IFunction interface {
  IEntity
  GetReturnType() IType
  IsVoid() bool
  GetCallingSymbol() ISymbol
  SetCallingSymbol(ISymbol)
}

type IVariable interface {
  IEntity
}

type IScope interface {
  IsToplevel() bool
  GetToplevel() IScope
  GetParent() IScope
  GetByName(string) IEntity
  CheckReferences(*ErrorHandler)
  AddChild(IScope)
}
