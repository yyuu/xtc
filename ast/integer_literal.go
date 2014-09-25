package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

// IntegerLiteralNode
type IntegerLiteralNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Value int64
}

func NewIntegerLiteralNode(loc core.Location, literal string) *IntegerLiteralNode {
  var ref core.ITypeRef
  var value int64
  var err error
  if ( startsWith(literal, "'") && endsWith(literal, "'") ) && 2 < len(literal) {
    _, err = fmt.Sscanf(literal[1:len(literal)-1], "%c", &value)
    ref = typesys.NewCharTypeRef(loc)
  } else {
    if ( startsWith(literal, "0X") || startsWith(literal, "0x") ) && 2 < len(literal) {
      // hexadecimal
      _, err = fmt.Sscanf(literal[2:], "%x", &value)
    } else {
      if startsWith(literal, "0") && 1 < len(literal) {
        // octal
        _, err = fmt.Sscanf(literal[1:], "%o", &value)
      } else {
        // decimal
        _, err = fmt.Sscanf(literal, "%d", &value)
      }
    }
    if endsWith(literal, "UL") {
      ref = typesys.NewUnsignedLongTypeRef(loc)
    } else if endsWith(literal, "L") {
      ref = typesys.NewLongTypeRef(loc)
    } else if endsWith(literal, "U") {
      ref = typesys.NewUnsignedIntTypeRef(loc)
    } else {
      ref = typesys.NewIntTypeRef(loc)
    }
  }
  if err != nil {
    panic(err)
  }
  return &IntegerLiteralNode { "ast.IntegerLiteralNode", loc, NewTypeNode(loc, ref), value }
}

func startsWith(s, prefix string) bool {
  return len(prefix) <= len(s) && strings.Index(s, prefix) == 0
}

func endsWith(s, suffix string) bool {
  return len(suffix) <= len(s) && strings.LastIndex(s, suffix) == len(s)-len(suffix)
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

func (self IntegerLiteralNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self IntegerLiteralNode) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *IntegerLiteralNode) SetType(t core.IType) {
  panic("IntegerLiteralNode#SetType called")
}

func (self IntegerLiteralNode) IsConstant() bool {
  return true
}

func (self IntegerLiteralNode) IsParameter() bool {
  return false
}

func (self IntegerLiteralNode) IsLvalue() bool {
  return false
}

func (self IntegerLiteralNode) IsAssignable() bool {
  return false
}

func (self IntegerLiteralNode) IsLoadable() bool {
  return false
}

func (self IntegerLiteralNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self IntegerLiteralNode) IsPointer() bool {
  return self.GetType().IsPointer()
}

func (self IntegerLiteralNode) GetValue() int64 {
  return self.Value
}
