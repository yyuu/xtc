package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestUnionTypeRefKeyString(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewUnionTypeRef(loc, "foo")
  xt.AssertEquals(t, "union foo { ... }", x.Key(), "union foo")
}
