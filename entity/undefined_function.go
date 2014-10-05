package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/typesys"
)

type UndefinedFunction struct {
  ClassName string
  TypeNode core.ITypeNode
  Name string
  Params *Params
  numRefered int
  memref core.IMemoryReference
  address core.IOperand
  callingSymbol core.ISymbol
}

func NewUndefinedFunction(t core.ITypeNode, name string, params *Params) *UndefinedFunction {
  return &UndefinedFunction {
    ClassName: "entity.UndefinedFunction",
    TypeNode: t,
    Name: name,
    Params: params,
    numRefered: 0,
    memref: nil,
    address: nil,
  }
}

func NewUndefinedFunctions(xs...*UndefinedFunction) []*UndefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []*UndefinedFunction { }
  }
}

func AsUndefinedFunction(x core.IEntity) *UndefinedFunction {
  return x.(*UndefinedFunction)
}

func (self *UndefinedFunction) String() string {
  return fmt.Sprintf("extern %s %s(%s) { ... } /* ref=%d */", self.TypeNode.GetTypeRef(), self.Name, self.Params, self.numRefered)
}

func (self *UndefinedFunction) IsDefined() bool {
  return false
}

func (self *UndefinedFunction) IsConstant() bool {
  return false
}

func (self *UndefinedFunction) IsPrivate() bool {
  return true
}

func (self *UndefinedFunction) IsParameter() bool {
  return false
}

func (self *UndefinedFunction) IsVariable() bool {
  return false
}

func (self *UndefinedFunction) GetNumRefered() int {
  return self.numRefered
}

func (self *UndefinedFunction) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *UndefinedFunction) Refered() {
  self.numRefered++
}

func (self *UndefinedFunction) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *UndefinedFunction) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *UndefinedFunction) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *UndefinedFunction) GetName() string {
  return self.Name
}

func (self *UndefinedFunction) GetParams() *Params {
  return self.Params
}

func (self *UndefinedFunction) GetParameters() []*Parameter {
  return self.Params.ParamDescs
}

func (self *UndefinedFunction) GetReturnType() core.IType {
  t := self.GetType().(*typesys.FunctionType)
  return t.GetReturnType()
}

func (self *UndefinedFunction) IsVoid() bool {
  return self.GetReturnType().IsVoid()
}

func (self *UndefinedFunction) GetValue() core.IExprNode {
  panic("UndefinedFunction#GetValue called")
}

func (self *UndefinedFunction) SymbolString() string {
  return self.Name
}

func (self *UndefinedFunction) GetMemref() core.IMemoryReference {
  checkAddress(self, self.memref, self.address)
  return self.memref
}

func (self *UndefinedFunction) SetMemref(memref core.IMemoryReference) {
  self.memref = memref
}

func (self *UndefinedFunction) GetAddress() core.IOperand {
  checkAddress(self, self.memref, self.address)
  return self.address
}

func (self *UndefinedFunction) SetAddress(address core.IOperand) {
  self.address = address
}

func (self *UndefinedFunction) GetCallingSymbol() core.ISymbol {
  if self.callingSymbol == nil {
    panic("must not happen: Function#callingSymbol called but nil")
  }
  return self.callingSymbol
}

func (self *UndefinedFunction) SetCallingSymbol(sym core.ISymbol) {
  if self.callingSymbol != nil {
    panic("must not happen: Function#callingSymbol called twice")
  }
  self.callingSymbol = sym
}
