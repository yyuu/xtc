package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestStructTypeRefKeyString(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewStructTypeRef(loc, "foo")
  xt.AssertEquals(t, "struct foo { ... }", x.Key(), "struct foo")
}
