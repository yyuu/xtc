package location

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
  return Location { "location.Locationn", sourceName, lineNumber, lineOffset }
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
