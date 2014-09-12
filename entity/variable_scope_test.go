package entity

import (
  "testing"
  "bitbucket.org/yyuu/bs/xt"
)

func TestToplevelScope(t *testing.T) {
  toplevel := NewToplevelScope()
  xt.AssertTrue(t, "toplevel is toplevel", toplevel.IsToplevel())
  xt.AssertEquals(t, "toplevel's toplevel is toplevel", toplevel.GetToplevel(), toplevel)
  xt.AssertTrue(t, "toplevel doesn't have parent", toplevel.GetParent() == nil)
}

func TestLocalScope(t *testing.T) {
  toplevel := NewToplevelScope()
  scope1 := NewLocalScope(toplevel)
  scope2 := NewLocalScope(scope1)
  xt.AssertFalse(t, "scope1 is not toplevel", scope1.IsToplevel())
  xt.AssertFalse(t, "scope2 is not toplevel", scope2.IsToplevel())
  xt.AssertEquals(t, "scope1's parent is toplevel", scope1.GetParent(), toplevel)
  xt.AssertEquals(t, "scope2's parent is scope1", scope2.GetParent(), scope1)
  xt.AssertEquals(t, "scope1's toplevel is toplevel", scope2.GetToplevel(), toplevel)
  xt.AssertEquals(t, "scope2's toplevel is toplevel", scope2.GetToplevel(), toplevel)
}
