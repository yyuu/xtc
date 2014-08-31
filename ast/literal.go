package ast

import (
  "fmt"
  "strings"
)

// IntegerLiteralNode
type IntegerLiteralNode struct {
  Location Location
  Value int
}

func NewIntegerLiteralNode(location Location, literal string) IntegerLiteralNode {
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
  return IntegerLiteralNode { location, value }
}

func (self IntegerLiteralNode) String() string {
//return strconv.Itoa(self.Value)
  return fmt.Sprintf("%d", self.Value)
}

func (self IntegerLiteralNode) IsExpr() bool {
  return true
}

func (self IntegerLiteralNode) GetLocation() Location {
  return self.Location
}

// StringLiteralNode
type StringLiteralNode struct {
  Location Location
  Value string
}

func NewStringLiteralNode(location Location, literal string) StringLiteralNode {
  return StringLiteralNode { location, literal }
}

func (self StringLiteralNode) String() string {
  return self.Value
}

func (self StringLiteralNode) IsExpr() bool {
  return true
}

func (self StringLiteralNode) GetLocation() Location {
  return self.Location
}
