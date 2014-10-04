package typesys

import (
  "testing"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestVoidRef1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewVoidTypeRef(loc)
  xt.AssertEquals(t, "void ref has location", x.GetLocation(), loc)
}

func TestVoidTypeRefKeyString(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewVoidTypeRef(loc)
  xt.AssertEquals(t, "void", x.Key(), "void")
}
