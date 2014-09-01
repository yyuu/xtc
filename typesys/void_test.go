package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
)

func TestVoid1(t *testing.T) {
  x := NewVoidType()
  assertEquals(t, "sizeof(void) == 1", x.Size(), 1)
  assertEquals(t, "sizeof(void) == 1", x.AllocSize(), 1)
  assertEquals(t, "sizeof(void) == 1", x.Alignment(), 1)
  assertTrue(t, "void is void", x.IsVoid())
}

func TestVoidRef1(t *testing.T) {
  location := ast.Location { "", 1, 2 }
  x := NewVoidTypeRef(location)
  assertEquals(t, "void ref has location", x.GetLocation(), location)
}
