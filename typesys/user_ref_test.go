package typesys

import (
  "testing"
  "bitbucket.org/yyuu/xtc/core"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestUserTypeRefKeyString(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewUserTypeRef(loc, "foo")
  xt.AssertEquals(t, "typedef foo ...", x.Key(), "foo")
}
