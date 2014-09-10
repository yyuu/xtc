package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/xt"
)

// signed int 32 : int32
func TestSignedInt32Type(t *testing.T) {
  x := NewIntegerType(32, true, "int32")
  xt.AssertEquals(t, "sizeof(int32) == 32", x.Size(), 32)
  xt.AssertEquals(t, "sizeof(int32) == 32", x.AllocSize(), 32)
  xt.AssertEquals(t, "sizeof(int32) == 32", x.Alignment(), 32)
  xt.AssertTrue(t, "int32 is an integer", x.IsInteger())
  xt.AssertTrue(t, "int32 is signed", x.IsSigned())
}

// unsigned int 32 : uint32
func TestUnsignedInt32Type(t *testing.T) {
  x := NewIntegerType(32, false, "uint32")
  xt.AssertEquals(t, "sizeof(uint32) == 32", x.Size(), 32)
  xt.AssertEquals(t, "sizeof(uint32) == 32", x.AllocSize(), 32)
  xt.AssertEquals(t, "sizeof(uint32) == 32", x.Alignment(), 32)
  xt.AssertTrue(t, "uint32 is an integer", x.IsInteger())
  xt.AssertFalse(t, "uint32 is not signed", x.IsSigned())
}

// unsigned int 8 : char
func TestCharType(t *testing.T) {
  x := NewIntegerType(8, false, "char")
  xt.AssertEquals(t, "sizeof(char) == 8", x.Size(), 8)
  xt.AssertEquals(t, "sizeof(char) == 8", x.AllocSize(), 8)
  xt.AssertEquals(t, "sizeof(char) == 8", x.Alignment(), 8)
  xt.AssertTrue(t, "char is an integer", x.IsInteger())
  xt.AssertFalse(t, "char is not signed", x.IsSigned())
}

func TestSignedInt32TypeRef(t *testing.T) {
  loc := duck.NewLocation("", 1, 2)
  x := NewIntegerTypeRef(loc, "int32")
  xt.AssertEquals(t, "int32 ref has location", x.GetLocation(), loc)
  xt.AssertEquals(t, "int32 is int32", x.Name, "int32")
}
