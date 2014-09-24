package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

type DefinedFunction struct {
  ClassName string
  Private bool
  TypeNode core.ITypeNode
  Name string
  Params *Params
  Body core.IStmtNode
  scope *LocalScope
  numRefered int
  ir_stmts []core.IStmt
}

func NewDefinedFunction(priv bool, t core.ITypeNode, name string, params *Params, body core.IStmtNode) *DefinedFunction {
  return &DefinedFunction { "entity.DefinedFunction", priv, t, name, params, body, nil, 0, []core.IStmt { } }
}

func NewDefinedFunctions(xs...*DefinedFunction) []*DefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []*DefinedFunction { }
  }
}

func (self DefinedFunction) String() string {
  return fmt.Sprintf("<entity.DefinedFunction Name=%s Private=%v TypeNode=%s Params=%s Body=%s>", self.Name, self.Private, self.TypeNode, self.Params, self.Body)
}

func (self DefinedFunction) IsPrivate() bool {
  return self.Private
}

func (self DefinedFunction) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self DefinedFunction) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self DefinedFunction) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self DefinedFunction) GetName() string {
  return self.Name
}

func (self DefinedFunction) IsDefined() bool {
  return true
}

func (self DefinedFunction) IsConstant() bool {
  return false
}

func (self DefinedFunction) IsParameter() bool {
  return false
}

func (self DefinedFunction) GetNumRefered() int {
  return self.numRefered
}

func (self DefinedFunction) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *DefinedFunction) Refered() {
  self.numRefered++
}

func (self DefinedFunction) GetParams() *Params {
  return self.Params
}

func (self DefinedFunction) GetParameters() []*Parameter {
  return self.Params.ParamDescs
}

func (self DefinedFunction) ListParameters() []*DefinedVariable {
  xs := self.Params.GetParamDescs()
  ys := make([]*DefinedVariable, len(xs))
  for i := range xs {
    ys[i] = xs[i].DefinedVariable
  }
  return ys
}

func (self DefinedFunction) GetBody() core.IStmtNode {
  return self.Body
}

func (self *DefinedFunction) SetBody(body core.IStmtNode) {
  self.Body = body
}

func (self DefinedFunction) GetScope() *LocalScope {
  return self.scope
}

func (self *DefinedFunction) SetScope(scope *LocalScope) {
  self.scope = scope
}

func (self DefinedFunction) GetReturnType() core.IType {
  t := self.GetType().(*typesys.FunctionType)
  return t.GetReturnType()
}

func (self DefinedFunction) IsVoid() bool {
  return self.GetReturnType().IsVoid()
}

func (self DefinedFunction) GetValue() core.IExprNode {
  panic("DefinedFunction#GetValue called")
}

func (self DefinedFunction) GetIR() []core.IStmt {
  return self.ir_stmts
}

func (self *DefinedFunction) SetIR(stmts []core.IStmt) {
  self.ir_stmts = stmts
}
