package entity

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestEmptyConstantTable(t *testing.T) {
  table := NewConstantTable()
  xt.AssertTrue(t, "empty constant table", table.IsEmpty())
  xt.AssertEquals(t, "empty constant table", len(table.GetEntries()), 0)
}

func TestConstantTableGetEntries(t *testing.T) {
  table := NewConstantTable()
  table.Intern("foo")
  xt.AssertEquals(t, "added a constant entry foo", len(table.GetEntries()), 1)
  table.Intern("bar")
  xt.AssertEquals(t, "added a constant entry bar", len(table.GetEntries()), 2)
}

func TestConstantTableEntityEquality(t *testing.T) {
  table := NewConstantTable()
  foo := table.Intern("foo")
  bar := table.Intern("bar")
  xt.AssertEquals(t, "intern returns same entry", table.Intern("foo"), foo)
  xt.AssertEquals(t, "intern returns same entry", table.Intern("bar"), bar)
}
