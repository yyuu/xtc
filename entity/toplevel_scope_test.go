package entity

import (
  "testing"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestToplevelScope(t *testing.T) {
  toplevel := NewToplevelScope()
  xt.AssertTrue(t, "toplevel is toplevel", toplevel.IsToplevel())
  xt.AssertEquals(t, "toplevel's toplevel is toplevel", toplevel.GetToplevel(), toplevel)
  xt.AssertTrue(t, "toplevel doesn't have parent", toplevel.GetParent() == nil)
}
