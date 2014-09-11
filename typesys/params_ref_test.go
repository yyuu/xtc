package typesys

import (
  "testing"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/xt"
)

func TestParamTypeRefsToString1(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewParamTypeRefs(loc,
         []core.ITypeRef { },
         false,
       )
  xt.AssertEquals(t, "(empty params)", x.String(), "")
}

func TestParamTypeRefsToString2(t *testing.T) {
  loc := core.NewLocation("", 1, 2)
  x := NewParamTypeRefs(loc,
         []core.ITypeRef {
           NewIntTypeRef(loc),
           NewArrayTypeRef(NewPointerTypeRef(NewCharTypeRef(loc)), -1),
         },
         false,
       )
  xt.AssertEquals(t, "int argc, char*[] argv", x.String(), "int,char*[]")
}
