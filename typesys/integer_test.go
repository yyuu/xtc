package typesys

import (
  "testing"
)

// signed int 32 : int32
func TestSignedInt32Type(t *testing.T) {
  x := NewIntegerType(32, true, "int32")
  assertEquals(t, "sizeof(int32) == 32", x.Size(), 32)
  assertEquals(t, "sizeof(int32) == 32", x.AllocSize(), 32)
  assertEquals(t, "sizeof(int32) == 32", x.Alignment(), 32)
  assertTrue(t, "int32 is an integer", x.IsInteger())
  assertTrue(t, "int32 is signed", x.IsSigned())
}

// unsigned int 32 : uint32
func TestUnsignedInt32Type(t *testing.T) {
  x := NewIntegerType(32, false, "uint32")
  assertEquals(t, "sizeof(uint32) == 32", x.Size(), 32)
  assertEquals(t, "sizeof(uint32) == 32", x.AllocSize(), 32)
  assertEquals(t, "sizeof(uint32) == 32", x.Alignment(), 32)
  assertTrue(t, "uint32 is an integer", x.IsInteger())
  assertFalse(t, "uint32 is not signed", x.IsSigned())
}

// unsigned int 8 : char
func TestCharType(t *testing.T) {
  x := NewIntegerType(8, false, "char")
  assertEquals(t, "sizeof(char) == 8", x.Size(), 8)
  assertEquals(t, "sizeof(char) == 8", x.AllocSize(), 8)
  assertEquals(t, "sizeof(char) == 8", x.Alignment(), 8)
  assertTrue(t, "char is an integer", x.IsInteger())
  assertFalse(t, "char is not signed", x.IsSigned())
}

func TestSignedInt32TypeRef(t *testing.T) {
  location := location { "", 1, 2 }
  x := NewIntegerTypeRef(location, "int32")
  assertEquals(t, "int32 ref has location", x.GetLocation(), location)
  assertEquals(t, "int32 is int32", x.Name, "int32")
}
