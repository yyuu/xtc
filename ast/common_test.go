package ast

import (
  "encoding/json"
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

func loc(lineNumber int, lineOffset int) Location {
  return Location { "", lineNumber, lineOffset }
}

func jsonString(x interface{}) string {
  cs, err := json.Marshal(x)
  if err != nil {
    panic(err)
  }
  return string(cs)
}
