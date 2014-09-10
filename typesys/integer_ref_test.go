package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestSignedInt32TypeRef(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewIntegerTypeRef(loc, "int32")
  xt.AssertEquals(t, "int32 ref has location", x.GetLocation(), loc)
  xt.AssertEquals(t, "int32 is int32", x.Name, "int32")
}
