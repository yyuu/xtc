package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type DefinedFunction struct {
  ClassName string
  Private bool
  TypeNode core.ITypeNode
  Name string
  Params core.IParams
  Body core.IStmtNode
  scope core.IVariableScope
}

func NewDefinedFunction(priv bool, t core.ITypeNode, name string, params core.IParams, body core.IStmtNode) *DefinedFunction {
  return &DefinedFunction { "entity.DefinedFunction", priv, t, name, params, body, nil }
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

func (self DefinedFunction) GetName() string {
  return self.Name
}

func (self DefinedFunction) IsEntity() bool {
  return true
}

func (self DefinedFunction) IsFunction() bool {
  return true
}

func (self DefinedFunction) IsDefinedFunction() bool {
  return true
}

func (self DefinedFunction) IsDefined() bool {
  return true
}

func (self DefinedFunction) IsConstant() bool {
  return false
}

func (self DefinedFunction) IsRefered() bool {
  return true // FIXME: count up references
}

func (self DefinedFunction) GetParams() core.IParams {
  return self.Params
}

func (self DefinedFunction) ListParameters() []core.IDefinedVariable {
  xs := self.Params.GetParamDescs()
  ys := make([]core.IDefinedVariable, len(xs))
  for i := range xs {
    ys[i] = xs[i].(core.IDefinedVariable)
  }
  return ys
}

func (self DefinedFunction) GetBody() core.IStmtNode {
  return self.Body
}

func (self DefinedFunction) SetBody(body core.IStmtNode) core.IDefinedFunction {
  self.Body = body
  return self
}

func (self DefinedFunction) GetScope() core.IVariableScope {
  return self.scope
}

func (self DefinedFunction) SetScope(scope core.IVariableScope) core.IDefinedFunction {
  self.scope = scope
  return self
}
