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
