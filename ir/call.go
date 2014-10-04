package ir

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/xtc/core"
)

type Call struct {
  ClassName string
  TypeId int
  Expr core.IExpr
  Args []core.IExpr
}

func NewCall(t int, expr core.IExpr, args []core.IExpr) *Call {
  return &Call { "ir.Call", t, expr, args }
}

func (self *Call) AsExpr() core.IExpr {
  return self
}

func (self Call) GetTypeId() int {
  return self.TypeId
}

func (self Call) GetExpr() core.IExpr {
  return self.Expr
}

func (self Call) GetArgs() []core.IExpr {
  return self.Args
}

func (self Call) NumArgs() int {
  return len(self.Args)
}

func (self Call) IsAddr() bool {
  return false
}

func (self Call) IsConstant() bool {
  return false
}

func (self Call) IsVar() bool {
  return false
}

func (self *Call) GetAddress() core.IOperand {
  panic("#GetAddress called")
}

func (self *Call) GetAsmValue() core.IImmediateValue {
  panic("#GetAsmValue called")
}

func (self *Call) GetMemref() core.IMemoryReference {
  panic("#GetMemref called")
}

func (self Call) GetAddressNode(t int) core.IExpr {
  panic("unexpected node for LHS")
}

func (self Call) GetEntityForce() core.IEntity {
  return nil
}

func (self Call) IsStaticCall() bool {
  ent := self.Expr.GetEntityForce()
  if ent == nil {
    return false
  } else {
    _, ok := ent.(core.IFunction)
    return ok
  }
}

func (self Call) GetFunction() core.IFunction {
  ent := self.Expr.GetEntityForce()
  if ent == nil {
    panic("not a static funcall")
  }
  return ent.(core.IFunction)
}

func (self Call) String() string {
  args := make([]string, len(self.Args))
  for i := range self.Args {
    args[i] = fmt.Sprint(self.Args[i])
  }
  return fmt.Sprintf("Call(%s,%s)", self.Expr, strings.Join(args, ","))
}
