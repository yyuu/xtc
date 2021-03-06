package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/typesys"
)

type DefinedFunction struct {
  ClassName string
  Private bool
  TypeNode core.ITypeNode
  Name string
  Params *Params
  Body core.IStmtNode
  IR []core.IStmt
  scope *LocalScope
  numRefered int
  memref core.IMemoryReference
  address core.IOperand
  callingSymbol core.ISymbol
}

func NewDefinedFunction(priv bool, t core.ITypeNode, name string, params *Params, body core.IStmtNode) *DefinedFunction {
  _, ok := t.GetTypeRef().(*typesys.FunctionTypeRef)
  if ! ok {
    panic("not a function type ref: " + t.GetTypeRef().String())
  }
  return &DefinedFunction {
    ClassName: "entity.DefinedFunction",
    Private: priv,
    TypeNode: t,
    Name: name,
    Params: params,
    Body: body,
    IR: []core.IStmt { },
    scope: nil,
    numRefered: 0,
    memref: nil,
    address: nil,
  }
}

func NewDefinedFunctions(xs...*DefinedFunction) []*DefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []*DefinedFunction { }
  }
}

func AsDefinedFunction(x core.IEntity) *DefinedFunction {
  return x.(*DefinedFunction)
}

func (self *DefinedFunction) String() string {
  var storage string
  if self.Private {
    storage = "static "
  }
  return fmt.Sprintf("%s%s %s(%s) { ... } /* ref=%d */", storage, self.TypeNode.GetTypeRef(), self.Name, self.Params, self.numRefered)
}

func (self *DefinedFunction) IsPrivate() bool {
  return self.Private
}

func (self *DefinedFunction) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *DefinedFunction) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *DefinedFunction) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *DefinedFunction) GetName() string {
  return self.Name
}

func (self *DefinedFunction) IsDefined() bool {
  return true
}

func (self *DefinedFunction) IsConstant() bool {
  return false
}

func (self *DefinedFunction) IsParameter() bool {
  return false
}

func (self *DefinedFunction) IsVariable() bool {
  return false
}

func (self *DefinedFunction) GetNumRefered() int {
  return self.numRefered
}

func (self *DefinedFunction) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *DefinedFunction) Refered() {
  self.numRefered++
}

func (self *DefinedFunction) GetParams() *Params {
  return self.Params
}

func (self *DefinedFunction) GetParameters() []*Parameter {
  return self.Params.ParamDescs
}

func (self *DefinedFunction) ListParameters() []*DefinedVariable {
  xs := self.Params.GetParamDescs()
  ys := make([]*DefinedVariable, len(xs))
  for i := range xs {
    ys[i] = xs[i].DefinedVariable
  }
  return ys
}

func (self *DefinedFunction) GetBody() core.IStmtNode {
  return self.Body
}

func (self *DefinedFunction) SetBody(body core.IStmtNode) {
  self.Body = body
}

func (self *DefinedFunction) GetScope() *LocalScope {
  return self.scope
}

func (self *DefinedFunction) SetScope(scope *LocalScope) {
  self.scope = scope
}

func (self *DefinedFunction) LocalVariableScope() *LocalScope {
  return self.Body.GetScope().(*LocalScope)
}

func (self *DefinedFunction) GetLocalVariables() []*DefinedVariable {
  return self.scope.AllLocalVariables()
}

func (self *DefinedFunction) GetReturnType() core.IType {
  t := self.GetType().(*typesys.FunctionType)
  return t.GetReturnType()
}

func (self *DefinedFunction) IsVoid() bool {
  return self.GetReturnType().IsVoid()
}

func (self *DefinedFunction) GetValue() core.IExprNode {
  panic("DefinedFunction#GetValue called")
}

func (self *DefinedFunction) GetIR() []core.IStmt {
  return self.IR
}

func (self *DefinedFunction) SetIR(stmts []core.IStmt) {
  self.IR = stmts
}

func (self *DefinedFunction) SymbolString() string {
  return self.Name
}

func (self *DefinedFunction) GetMemref() core.IMemoryReference {
  checkAddress(self, self.memref, self.address)
  return self.memref
}

func (self *DefinedFunction) SetMemref(memref core.IMemoryReference) {
  self.memref = memref
}

func (self *DefinedFunction) GetAddress() core.IOperand {
  checkAddress(self, self.memref, self.address)
  return self.address
}

func (self *DefinedFunction) SetAddress(address core.IOperand) {
  self.address = address
}

func (self *DefinedFunction) GetCallingSymbol() core.ISymbol {
  if self.callingSymbol == nil {
    panic("must not happen: Function#callingSymbol called but nil")
  }
  return self.callingSymbol
}

func (self *DefinedFunction) SetCallingSymbol(sym core.ISymbol) {
  if self.callingSymbol != nil {
    panic("must not happen: Function#callingSymbol called twice")
  }
  self.callingSymbol = sym
}
