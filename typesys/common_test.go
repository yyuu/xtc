package typesys

import (
  "testing"
)

func assertEquals(t *testing.T, key string, got, expected interface{}) {
  if got != expected {
    t.Errorf("%s: expected %v != got %v", key, expected, got)
    t.Fail()
  }
}

func assertNotEquals(t *testing.T, key string, got, expected interface{}) {
  if got == expected {
    t.Errorf("%s: expected %v == got %v", key, expected, got)
    t.Fail()
  }
}

func assertTrue(t *testing.T, key string, got bool) {
  if ! got {
    t.Errorf("%s: expected true", key)
    t.Fail()
  }
}

func assertFalse(t *testing.T, key string, got bool) {
  if got {
    t.Errorf("%s: expected false", key)
    t.Fail()
  }
}

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