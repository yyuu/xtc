package parser

import (
  "testing"
)

func TestParseEmpty(t *testing.T) {
  nodes, err := ParseExpr("")
  if err != nil {
    t.Fail()
  }
  if 0 < len(nodes) {
    t.Fail()
  }
}
