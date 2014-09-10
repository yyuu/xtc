package ast

import (
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/core"
)

// IntegerLiteralNode
type IntegerLiteralNode struct {
  ClassName string
  Location core.Location
  Value int
}

func NewIntegerLiteralNode(loc core.Location, literal string) IntegerLiteralNode {
  var value int
  var err error
  if ( strings.Index(literal, "'") == 0 && strings.LastIndex(literal, "'") == len(literal)-1 ) && 2 < len(literal) {
    _, err = fmt.Sscanf(literal[1:len(literal)-1], "%c", &value)
  } else {
    if ( strings.Index(literal, "0X") == 0 || strings.Index(literal, "0x") == 0 ) && 2 < len(literal) {
      // hexadecimal
      _, err = fmt.Sscanf(literal[2:], "%x", &value)
    } else {
      if ( strings.Index(literal, "0") == 0 ) && 1 < len(literal) {
        // octal
        _, err = fmt.Sscanf(literal[1:], "%o", &value)
      } else {
        // decimal
        _, err = fmt.Sscanf(literal, "%d", &value)
      }
    }
  }
  if err != nil {
    panic(err)
  }
  return IntegerLiteralNode { "ast.IntegerLiteralNode", loc, value }
}

func (self IntegerLiteralNode) String() string {
  return fmt.Sprintf("%d", self.Value)
}

func (self IntegerLiteralNode) IsExprNode() bool {
  return true
}

func (self IntegerLiteralNode) GetLocation() core.Location {
  return self.Location
}
