package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// FuncallNode
type FuncallNode struct {
  ClassName string
  Location core.Location
  Expr core.IExprNode
  Args []core.IExprNode
}

func NewFuncallNode(loc core.Location, expr core.IExprNode, args []core.IExprNode) *FuncallNode {
  if expr == nil { panic("expr is nil") }
  return &FuncallNode { "ast.FuncallNode", loc, expr, args }
}

func (self FuncallNode) String() string {
  sArgs := make([]string, len(self.Args))
  for i := range self.Args {
    sArgs[i] = fmt.Sprintf("%s", self.Args[i])
  }
  if len(sArgs) == 0 {
    return fmt.Sprintf("(%s)", self.Expr)
  } else {
    return fmt.Sprintf("(%s %s)", self.Expr, strings.Join(sArgs, " "))
  }
}

func (self *FuncallNode) AsExprNode() core.IExprNode {
  return self
}

func (self FuncallNode) GetLocation() core.Location {
  return self.Location
}

func (self FuncallNode) GetExpr() core.IExprNode {
  return self.Expr
}

func (self FuncallNode) GetArgs() []core.IExprNode {
  return self.Args
}

func (self FuncallNode) NumArgs() int {
  return len(self.Args)
}

func (self FuncallNode) GetFunctionType() *typesys.FunctionType {
//pt := self.Expr.GetType().(*typesys.PointerType)
//return pt.GetBaseType().(*typesys.FunctionType)
  return self.Expr.GetType().(*typesys.FunctionType)
}

func (self FuncallNode) GetType() core.IType {
  return self.GetFunctionType().GetReturnType()
}

func (self *FuncallNode) SetType(t core.IType) {
  panic("FuncallNode#SetType called")
}

func (self FuncallNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self FuncallNode) IsConstant() bool {
  return false
}

func (self FuncallNode) IsParameter() bool {
  return false
}

func (self FuncallNode) IsLvalue() bool {
  return false
}

func (self FuncallNode) IsAssignable() bool {
  return false
}

func (self FuncallNode) IsLoadable() bool {
  return false
}

func (self FuncallNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self FuncallNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
