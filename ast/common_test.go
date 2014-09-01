package ast

import (
  "encoding/json"
  "testing"
)

func assertJsonEquals(t *testing.T, got INode, expected string) {
  s := jsonString(got)
  if s != expected {
    t.Errorf("\n// expected\n%s\n// got\n%s\n", expected, s)
    t.Fail()
  }
}

func loc(lineNumber int, lineOffset int) Location {
  return Location { "", lineNumber, lineOffset }
}

func jsonString(x interface{}) string {
  cs, err := json.MarshalIndent(x, "", "  ")
  if err != nil {
    panic(err)
  }
  return string(cs)
}
