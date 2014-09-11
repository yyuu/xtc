package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestIntArrayTypeRefToString1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewArrayTypeRef(NewIntTypeRef(loc), 255)
  xt.AssertEquals(t, "int[255]", x.String(), "int[255]")
}

func TestIntArrayTypeRefToString2(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewArrayTypeRef(NewIntTypeRef(loc), -1)
  xt.AssertEquals(t, "int[]", x.String(), "int[]")
}

func TestCharPointerArrayTypeRefToString(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewArrayTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)), -1)
  xt.AssertEquals(t, "char*[]", x.String(), "char*[]")
}
