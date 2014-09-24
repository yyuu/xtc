package core

import (
  "fmt"
)

type Location struct {
  className string
  sourceName string
  lineNumber int
  lineOffset int
}

func NewLocation(sourceName string, lineNumber int, lineOffset int) Location {
  return Location { "core.Locationn", sourceName, lineNumber, lineOffset }
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
  // display as "1 origin"
  return fmt.Sprintf("[%s:%d,%d]", self.sourceName, self.lineNumber+1, self.lineOffset+1)
}

func (self Location) MarshalJSON() ([]byte, error) {
  s := fmt.Sprintf("%q", self.String())
  return []byte(s), nil
}
