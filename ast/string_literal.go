package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

// StringLiteralNode
type StringLiteralNode struct {
  ClassName string
  Location core.Location
  TypeNode core.ITypeNode
  Value string
  entry *entity.ConstantEntry
}

func NewStringLiteralNode(loc core.Location, literal string) *StringLiteralNode {
  ref := typesys.NewPointerTypeRef(typesys.NewCharTypeRef(loc))
  t := NewTypeNode(loc, ref)
  s, err := parseStringLiteral(literal)
  if err != nil {
    panic(err)
  }
  return &StringLiteralNode { "ast.StringLiteralNode", loc, t, s, nil }
}

func parseStringLiteral(src string) (string, error) {
  var buf []rune
  for _, r := range src {
    buf = append(buf, r)
  }

  if 2 < len(buf) && buf[0] == '"' && buf[len(buf)-1] == '"' {
    buf = buf[1:len(buf)-1] // remove quotations
  } else {
    return "", fmt.Errorf("not a quoted string literal: %v", src)
  }

  var res string
  for {
    var r rune
    if len(buf) < 1 {
      return res, nil
    }
    r, buf = buf[0], buf[1:]
    switch r {
      case '\\': {
        if len(buf) < 1 {
          return res, fmt.Errorf("no string given after escape sequence")
        }
        r, buf = buf[0], buf[1:]
        switch r {
          case '\\': res += "\\"
          case 'n':  res += "\n"
          case 'r':  res += "\r"
          case 't':  res += "\t"
          default: {
            return res, fmt.Errorf("unknown escape sequence: %q" + string(r))
          }
        }
      }
      default: {
        res += string(r)
      }
    }
  }
}

func (self StringLiteralNode) String() string {
  return self.Value
}

func (self *StringLiteralNode) AsExprNode() core.IExprNode {
  return self
}

func (self StringLiteralNode) GetLocation() core.Location {
  return self.Location
}

func (self *StringLiteralNode) GetValue() string {
  return self.Value
}

func (self *StringLiteralNode) GetEntry() *entity.ConstantEntry {
  return self.entry
}

func (self *StringLiteralNode) SetEntry(e *entity.ConstantEntry) {
  self.entry = e
}

func (self *StringLiteralNode) GetTypeNode() core.ITypeNode {
  return self.TypeNode
}

func (self *StringLiteralNode) GetTypeRef() core.ITypeRef {
  return self.TypeNode.GetTypeRef()
}

func (self *StringLiteralNode) GetType() core.IType {
  return self.TypeNode.GetType()
}

func (self *StringLiteralNode) SetType(t core.IType) {
  panic("#SetType called")
}

func (self *StringLiteralNode) GetOrigType() core.IType {
  return self.GetType()
}

func (self *StringLiteralNode) IsConstant() bool {
  return true
}

func (self *StringLiteralNode) IsParameter() bool {
  return false
}

func (self *StringLiteralNode) IsLvalue() bool {
  return false
}

func (self *StringLiteralNode) IsAssignable() bool {
  return false
}

func (self *StringLiteralNode) IsLoadable() bool {
  return false
}

func (self *StringLiteralNode) IsCallable() bool {
  return self.GetType().IsCallable()
}

func (self *StringLiteralNode) IsPointer() bool {
  return self.GetType().IsPointer()
}
