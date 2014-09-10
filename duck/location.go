package duck

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
  return Location { "duck.Locationn", sourceName, lineNumber, lineOffset }
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
