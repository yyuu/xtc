package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
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
  memref core.IMemoryReference
  address core.IOperand
}

func NewDefinedVariable(isPrivate bool, t core.ITypeNode, name string, init core.IExprNode) *DefinedVariable {
  return &DefinedVariable {
    ClassName: "entity.DefinedVariable",
    Private: isPrivate,
    Name: name,
    TypeNode: t,
    Initializer: init,
    IR: nil,
    numRefered: 0,
    sequence: -1,
    memref: nil,
    address: nil,
  }
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
  var storage string
  if self.Private {
    storage = "static "
  }
  var init string
  if self.HasInitializer() {
    init = fmt.Sprintf(" = %s", self.Initializer)
  }
  return fmt.Sprintf("%s%s %s%s; /* ref=%d */", storage, self.TypeNode.GetTypeRef(), self.Name, init, self.numRefered)
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

func (self *DefinedVariable) GetMemref() core.IMemoryReference {
  checkAddress(self, self.memref, self.address)
  return self.memref
}

func (self *DefinedVariable) SetMemref(memref core.IMemoryReference) {
  self.memref = memref
}

func (self *DefinedVariable) GetAddress() core.IOperand {
  checkAddress(self, self.memref, self.address)
  return self.address
}

func (self *DefinedVariable) SetAddress(address core.IOperand) {
  self.address = address
}
