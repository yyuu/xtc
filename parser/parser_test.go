package parser

import (
  "testing"
)

func TestParseEmpty(t *testing.T) {
  _, err := ParseExpr("")
  if err != nil {
    t.Fail()
  }
}
