package ast

import (
  "fmt"
  "strconv"
  "strings"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/typesys"
)

// IntegerLiteralNode
type IntegerLiteralNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Value int64
}

func NewIntegerLiteralNode(loc core.Location, literal string) *IntegerLiteralNode {
  var value int64
  _, err := fmt.Sscanf(literal, "%d", &value)
  if err != nil { panic(err) }
  var ref core.ITypeRef
  n := len(literal)
  switch {
    case 2 <= n && strings.LastIndex(literal, "UL") == n-2: ref = typesys.NewUnsignedLongTypeRef(loc)
    case 1 <= n && strings.LastIndex(literal, "L")  == n-1: ref = typesys.NewLongTypeRef(loc)
    case 1 <= n && strings.LastIndex(literal, "U")  == n-1: ref = typesys.NewUnsignedIntTypeRef(loc)
    default:                                                ref = typesys.NewIntTypeRef(loc)
  }
  return &IntegerLiteralNode { "ast.IntegerLiteralNode", loc, NewTypeNode(loc, ref), value }
}

func NewCharacterLiteralNode(loc core.Location, literal string) *IntegerLiteralNode {
  ref := typesys.NewCharTypeRef(loc)
  value, err := strconv.Atoi(literal)
  if err != nil { panic(err) }
  return &IntegerLiteralNode { "ast.IntegerLiteralNode", loc, NewTypeNode(loc, ref), int64(value) }
}

func (self IntegerLiteralNode) String() string {
  return fmt.Sprintf("%d", self.Value)
}

func (self *IntegerLiteralNode) AsExprNode() core.IExprNode {
  return self
}

func (self IntegerLiteralNode) GetLocation() core.Location {
  return self.Location
}

func (self *IntegerLiteralNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *IntegerLiteralNode) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *IntegerLiteralNode) SetType(t core.IType) {
  panic("#SetType called")
}

func (self *IntegerLiteralNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self *IntegerLiteralNode) IsConstant() bool {
  return true
}

func (self *IntegerLiteralNode) IsParameter() bool {
  return false
}

func (self *IntegerLiteralNode) IsLvalue() bool {
  return false
}

func (self *IntegerLiteralNode) IsAssignable() bool {
  return false
}

func (self *IntegerLiteralNode) IsLoadable() bool {
  return false
}

func (self *IntegerLiteralNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *IntegerLiteralNode) IsPointer() bool {
  return self.GetType().IsPointer()
}

func (self *IntegerLiteralNode) GetValue() int64 {
  return self.Value
}
