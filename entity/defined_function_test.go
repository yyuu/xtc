package entity

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func TestDefinedFunctionWithFunctionType(t *testing.T) {
/*
  int f() {
    return 0;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  f := NewDefinedFunction(
    false,
    xtc_ast.NewTypeNode(loc,
      xtc_typesys.NewFunctionTypeRef(
        xtc_typesys.NewIntTypeRef(loc),
        xtc_typesys.NewParamTypeRefs(loc, []xtc_core.ITypeRef { }, false),
      ),
    ),
    "f",
    NewParams(loc, NewParameters(), false),
    xtc_ast.NewBlockNode(loc,
      nil,
      []xtc_core.IStmtNode {
        xtc_ast.NewReturnNode(loc, xtc_ast.NewIntegerLiteralNode(loc, "0")),
      },
    ),
  )
  xt.AssertNotNil(t, "f should not be nil", f)
}

func TestDefinedFunctionWithInvalidType(t *testing.T) {
/*
  long g() {
    return 0;
  }
 */
  defer func() {
    r := recover()
    if r == "" {
      t.Errorf("functions must be declared as function type")
    }
  }()
  loc := xtc_core.NewLocation("", 0, 0)
  NewDefinedFunction(
    false,
    xtc_ast.NewTypeNode(loc, xtc_typesys.NewLongTypeRef(loc)),
    "g",
    NewParams(loc, NewParameters(), false),
    xtc_ast.NewBlockNode(loc,
      nil,
      []xtc_core.IStmtNode {
        xtc_ast.NewReturnNode(loc, xtc_ast.NewIntegerLiteralNode(loc, "0")),
      },
    ),
  )
}
