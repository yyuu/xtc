package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestFunctionTypeRefToString1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewFunctionTypeRef(
    NewIntTypeRef(loc),
    NewParamTypeRefs(loc,
      []core.ITypeRef {
        NewIntTypeRef(loc),
      },
      false,
    ),
  )
  xt.AssertEquals(t, "int f(int x)", x.String(), "int(int)")
}

func TestFunctionTypeRefToString2(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewFunctionTypeRef(
    NewIntTypeRef(loc),
    NewParamTypeRefs(loc,
      []core.ITypeRef {
        NewIntTypeRef(loc),
        NewArrayTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)), -1),
      },
      false,
    ),
  )
  xt.AssertEquals(t, "int main(int argc, char*[] argv)", x.String(), "int(int,char*[])")
}
