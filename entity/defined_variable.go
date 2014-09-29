package entity

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type DefinedVariable struct {
  ClassName string
  Private bool
  Name string
  TypeNode core.ITypeNode
  Initializer core.IExprNode
  IR core.IExpr
  numRefered int
  sequence int
}

func NewDefinedVariable(isPrivate bool, t core.ITypeNode, name string, init core.IExprNode) *DefinedVariable {
  return &DefinedVariable { "entity.DefinedVariable", isPrivate, name, t, init, nil, 0, -1 }
}

func NewDefinedVariables(xs...*DefinedVariable) []*DefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []*DefinedVariable { }
  }
}

func AsDefinedVariable(x core.IEntity) *DefinedVariable {
  return x.(*DefinedVariable)
}

func AsDefinedVariables(xs []core.IEntity) []*DefinedVariable {
  ys := make([]*DefinedVariable, len(xs))
  for i := range xs {
    ys[i] = xs[i].(*DefinedVariable)
  }
  return ys
}

var tmpSeq int = 0
func temporaryDefinedVariable(t core.ITypeNode) *DefinedVariable {
  tmpSeq++
  return NewDefinedVariable(false, t, fmt.Sprintf("@tmp%d", tmpSeq), nil)
}

func (self *DefinedVariable) String() string {
  return fmt.Sprintf("<entity.DefinedVariable Name=%s Private=%v TypeNode=%s NumRefered=%d Initializer=%s>", self.Name, self.Private, self.TypeNode, self.numRefered, self.Initializer)
}

func (self *DefinedVariable) IsDefined() bool {
  return true
}

func (self *DefinedVariable) HasInitializer() bool {
  return self.Initializer != nil
}

func (self *DefinedVariable) IsPrivate() bool {
  return self.Private
}

func (self *DefinedVariable) IsConstant() bool {
  return false
}

func (self *DefinedVariable) IsParameter() bool {
  return false
}

func (self *DefinedVariable) IsVariable() bool {
  return true
}

func (self *DefinedVariable) GetName() string {
  return self.Name
}

func (self *DefinedVariable) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *DefinedVariable) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *DefinedVariable) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *DefinedVariable) GetNumRefered() int {
  return self.numRefered
}

func (self *DefinedVariable) IsRefered() bool {
  return 0 < self.numRefered
}

func (self *DefinedVariable) Refered() {
  self.numRefered++
}

func (self *DefinedVariable) GetInitializer() core.IExprNode {
  return self.Initializer
}

func (self *DefinedVariable) SetInitializer(init core.IExprNode) {
  self.Initializer = init
}

func (self *DefinedVariable) GetValue() core.IExprNode {
  panic("DefinedVariable#GetValue called")
}

func (self *DefinedVariable) GetIR() core.IExpr {
  return self.IR
}

func (self *DefinedVariable) SetIR(expr core.IExpr) {
  self.IR = expr
}

func (self *DefinedVariable) SetSequence(seq int) {
  self.sequence = seq
}

func (self *DefinedVariable) SymbolString() string {
  if self.sequence < 0 {
    return self.Name
  } else {
    return fmt.Sprintf("%s.%d", self.Name, self.sequence)
  }
}
