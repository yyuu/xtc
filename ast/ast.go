package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Location struct {
  sourceName string
  lineNumber int
  lineOffset int
}

func NewLocation(sourceName string, lineNumber int, lineOffset int) duck.ILocation {
  return Location { sourceName, lineNumber, lineOffset }
}

func (self Location) GetSourceName() string {
  return self.sourceName
}

func (self Location) GetLineNumber() int {
  return self.lineNumber
}

func (self Location) GetLineOffset() int {
  return self.lineOffset
}

func (self Location) String() string {
  return fmt.Sprintf("[%s:%d,%d]", self.sourceName, self.lineNumber, self.lineOffset)
}

func (self Location) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.String())
  return []byte(s), nil
}

type AST struct {
  location duck.ILocation
  declarations Declarations
}

func NewAST(loc duck.ILocation, declarations Declarations) AST {
  if loc == nil { panic("location is nil") }
  return AST { loc, declarations }
}

func (self AST) String() string {
  return fmt.Sprintf(";; %s\n%s", self.location, self.declarations)
}

func (self AST) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Location duck.ILocation
    Declarations Declarations
  }
  x.ClassName = "ast.AST"
  x.Location = self.location
  x.Declarations = self.declarations
  return json.Marshal(x)
}

func (self AST) GetLocation() duck.ILocation {
  return self.location
}

func (self AST) GetDeclarations() Declarations {
  return self.declarations
}
