package ast

import (
  "regexp"
  "strings"
  "testing"
)

func assertEquals(t *testing.T, got interface{}, expected interface{}) {
  if got != expected {
    t.Errorf("not equals: expected %q, got %q", expected, got)
  }
}

func trimSpace(s string) string {
  re := regexp.MustCompile("\\s+")
  return re.ReplaceAllString(strings.TrimSpace(s), " ")
}

// dummy
type location struct {
  sourceName string
  lineNumber int
  lineOffset int
}

func (self location) GetSourceName() string {
  return self.sourceName
}

func (self location) GetLineNumber() int {
  return self.lineNumber
}

func (self location) GetLineOffset() int {
  return self.lineOffset
}

var LOC = location { "__test__", 0, 0 }
