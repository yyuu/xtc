package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestVoidRef1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewVoidTypeRef(loc)
  xt.AssertEquals(t, "void ref has location", x.GetLocation(), loc)
}

func TestVoidTypeRefToString(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewVoidTypeRef(loc)
  xt.AssertEquals(t, "void", x.String(), "void")
}
