package ast

import (
  "encoding/json"
  "fmt"
  "strings"
  "bitbucket.org/yyuu/bs/duck"
)

// IntegerLiteralNode
type IntegerLiteralNode struct {
  location duck.ILocation
  value int
}

func NewIntegerLiteralNode(loc duck.ILocation, literal string) IntegerLiteralNode {
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
  return IntegerLiteralNode { loc, value }
}

func (self IntegerLiteralNode) String() string {
  return fmt.Sprintf("%d", self.value)
}

func (self IntegerLiteralNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Value int
  }
  x.ClassName = "ast.IntegerLiteralNode"
  x.Location = self.location
  x.Value = self.value
  return json.Marshal(x)
}

func (self IntegerLiteralNode) IsExpr() bool {
  return true
}

func (self IntegerLiteralNode) GetLocation() duck.ILocation {
  return self.location
}

// StringLiteralNode
type StringLiteralNode struct {
  location duck.ILocation
  value string
}

func NewStringLiteralNode(loc duck.ILocation, literal string) StringLiteralNode {
  return StringLiteralNode { loc, literal }
}

func (self StringLiteralNode) String() string {
  return self.value
}

func (self StringLiteralNode) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Value string
  }
  x.ClassName = "ast.StringLiteralNode"
  x.Location = self.location
  x.Value = self.value
  return json.Marshal(x)
}

func (self StringLiteralNode) IsExpr() bool {
  return true
}

func (self StringLiteralNode) GetLocation() duck.ILocation {
  return self.location
}
