package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Location struct {
  ClassName string
  SourceName string
  LineNumber int
  LineOffset int
}

func NewLocation(sourceName string, lineNumber int, lineOffset int) duck.ILocation {
  return Location { "ast.Locationn", sourceName, lineNumber, lineOffset }
}

func (self Location) GetSourceName() string {
  return self.SourceName
}

func (self Location) GetLineNumber() int {
  return self.LineNumber
}

func (self Location) GetLineOffset() int {
  return self.LineOffset
}

func (self Location) String() string {
  return fmt.Sprintf("[%s:%d,%d]", self.SourceName, self.LineNumber, self.LineOffset)
}

func (self Location) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.String())
  return []byte(s), nil
}

type AST struct {
  ClassName string
  Location duck.ILocation
  Declarations Declarations
}

func NewAST(loc duck.ILocation, declarations Declarations) AST {
  if loc == nil { panic("location is nil") }
  return AST { "ast.AST", loc, declarations }
}

func (self AST) String() string {
  return fmt.Sprintf(";; %s\n%s", self.Location, self.Declarations)
}

func (self AST) GetLocation() duck.ILocation {
  return self.Location
}

func (self AST) GetDeclarations() Declarations {
  return self.Declarations
}
