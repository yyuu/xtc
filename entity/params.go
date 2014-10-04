package entity

import (
  "fmt"
  "bitbucket.org/yyuu/xtc/core"
)

type Params struct {
  ClassName string
  Location core.Location
  ParamDescs []*Parameter
  Vararg bool
}

func NewParams(loc core.Location, paramDescs []*Parameter, vararg bool) *Params {
  return &Params { "entity.Params", loc, paramDescs, vararg }
}

func AsParams(x core.IEntity) *Params {
  return x.(*Params)
}

func (self *Params) String() string {
  return fmt.Sprintf("<entity.Params Location=%s ParamDescs=%s>", self.Location, self.ParamDescs)
}

func (self *Params) IsDefined() bool {
  return false
}

func (self *Params) IsConstant() bool {
  return false
}

func (self *Params) IsPrivate() bool {
  return true
}

func (self *Params) IsParameter() bool {
  return false
}

func (self *Params) IsVariable() bool {
  return false
}

func (self *Params) GetNumRefered() int {
  return -1
}

func (self *Params) IsRefered() bool {
  return false
}

func (self *Params) IsVararg() bool {
  return self.Vararg
}

func (self *Params) Refered() {
  // nop
}

func (self *Params) GetLocation() core.Location {
  return self.Location
}

func (self *Params) GetParamDescs() []*Parameter {
  return self.ParamDescs
}

func (self *Params) GetName() string {
  panic("Params#GetName called")
}

func (self *Params) GetTypeNode() core.ITypeNode {
  panic("Params#GetTypeNode called")
}

func (self *Params) GetTypeRef() core.ITypeRef {
  panic("Params#GetTypeRef called")
}

func (self *Params) GetType() core.IType {
  panic("Params#GetType called")
}

func (self *Params) GetValue() core.IExprNode {
  panic("Params#GetValue called")
}

func (self *Params) SymbolString() string {
  panic("Params#SymbolString called")
}

func (self *Params) GetMemref() core.IMemoryReference {
  panic("Params#GetMemref called")
}

func (self *Params) SetMemref(memref core.IMemoryReference) {
  panic("Params#SetMemref called")
}

func (self *Params) GetAddress() core.IOperand {
  panic("Params#GetAddress called")
}

func (self *Params) SetAddress(address core.IOperand) {
  panic("Params#SetAddress called")
}
