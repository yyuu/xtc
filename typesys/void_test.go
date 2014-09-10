package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/xt"
)

func TestVoid1(t *testing.T) {
  x := NewVoidType()
  xt.AssertEquals(t, "sizeof(void) == 1", x.Size(), 1)
  xt.AssertEquals(t, "sizeof(void) == 1", x.AllocSize(), 1)
  xt.AssertEquals(t, "sizeof(void) == 1", x.Alignment(), 1)
  xt.AssertTrue(t, "void is void", x.IsVoid())
}

func TestVoidRef1(t *testing.T) {
  loc := duck.NewLocation("", 1, 2)
  x := NewVoidTypeRef(loc)
  xt.AssertEquals(t, "void ref has location", x.GetLocation(), loc)
}
