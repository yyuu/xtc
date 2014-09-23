package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestUserTypeRefKeyString(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewUserTypeRef(loc, "foo")
  xt.AssertEquals(t, "typedef foo ...", x.Key(), "foo")
}
