package parser

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
//"bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/xt"
)

func TestParseEmpty(t *testing.T) {
  _, err := ParseExpr("")
  xt.AssertNil(t, "", err)
}

func defvars(xs...entity.DefinedVariable) []entity.DefinedVariable {
  return xs
}

func vardecls(xs...entity.UndefinedVariable) []entity.UndefinedVariable {
  return xs
}

func defuns(xs...entity.DefinedFunction) []entity.DefinedFunction {
  return xs
}

func funcdecls(xs...entity.UndefinedFunction) []entity.UndefinedFunction {
  return xs
}

func defconsts(xs...entity.Constant) []entity.Constant {
  return xs
}

func defstructs(xs...ast.StructNode) []ast.StructNode {
  return xs
}

func defunions(xs...ast.UnionNode) []ast.UnionNode {
  return xs
}

func typedefs(xs...ast.TypedefNode) []ast.TypedefNode {
  return xs
}

/*
func TestParseFuncallWithoutArguments(t *testing.T) {
  s := `
    int f(void) {
      return getc();
    }
  `
  x := ast.AST {
    loc(1,1),
    ast.Declarations {
      Defvars: defvars(),
      Vardecls: vardecls(),
      Defuns: defuns(
//      []duck.IStmtNode {
//        ast.NewExprStmtNode(loc(1,1),
//          ast.NewFuncallNode(loc(1,1),
//            ast.NewVariableNode(loc(1,1),
//              "gets",
//            ),
//            []duck.IExprNode {
//            },
//          ),
//        ),
//      },
      ),
      Funcdecls: funcdecls(),
      Constants: defconsts(),
      Defstructs: defstructs(),
      Defunions: defunions(),
      Typedefs: typedefs(),
    },
  }
  y, err := ParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "parse funcall w/o arguments", xt.JSON(y), xt.JSON(x))
  xt.AssertDeepEquals(t, "", y, x)
}
 */

/*
func TestParseFuncallWithSingleArgument(t *testing.T) {
  s := `
    void f(int n) {
      println("hello, %d", n);
    }
  `
  x := ast.AST {
//  []duck.IStmtNode {
//    ast.NewExprStmtNode(loc(1,5),
//      ast.NewFuncallNode(loc(1,5),
//        ast.NewVariableNode(loc(1,5),
//          "println",
//        ),
//        []duck.IExprNode {
//          ast.NewStringLiteralNode(loc(1,13),
//            "\"hello, world\"",
//          ),
//        },
//      ),
//    ),
//  },
    loc(1,5),
    ast.Declarations {
      Defvars: []entity.DefinedVariable { },
      Vardecls: []entity.UndefinedVariable { },
      Defuns: []entity.DefinedFunction {
      },
      Funcdecls: []entity.UndefinedFunction { },
      Constants: []entity.Constant { },
      Defstructs: []ast.StructNode { },
      Defunions: []ast.UnionNode { },
      Typedefs: []ast.TypedefNode { },
    },
  }
  y, err := ParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "parse funcall w/ single argument", xt.JSON(y), xt.JSON(x))
  xt.AssertDeepEquals(t, "", y, x)
}
 */

/*
func TestParseFuncallWithMultipleArguments(t *testing.T) {
  s := `

    println(
      "hello, %s",
      "world"
    );
  `
  x := ast.AST {
//  []duck.IStmtNode {
//    ast.NewExprStmtNode(loc(2,5),
//      ast.NewFuncallNode(loc(2,5),
//        ast.NewVariableNode(loc(2,5),
//          "println",
//        ),
//        []duck.IExprNode {
//          ast.NewStringLiteralNode(loc(3,7),
//            "\"hello, %s\"",
//          ),
//          ast.NewStringLiteralNode(loc(4,7),
//            "\"world\"",
//          ),
//        },
//      ),
//    ),
//  },
    loc(2,5),
    ast.Declarations {
      Defvars: []entity.DefinedVariable { },
      Vardecls: []entity.UndefinedVariable { },
      Defuns: []entity.DefinedFunction { },
      Funcdecls: []entity.UndefinedFunction { },
      Constants: []entity.Constant { },
      Defstructs: []ast.StructNode { },
      Defunions: []ast.UnionNode { },
      Typedefs: []ast.TypedefNode { },
    },
  }
  y, err := ParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "funcall w/ multiple arguments", xt.JSON(y), xt.JSON(x))
  xt.AssertDeepEquals(t, "", y, x)
}
 */

/*
func TestFor1(t *testing.T) {
  s := `
    for (i=0; i<100; i++) println(i);
`
  x := ast.AST {
//  []duck.IStmtNode {
//    ast.NewForNode(loc(1,5),
//      ast.NewAssignNode(loc(1,10),
//        ast.NewVariableNode(loc(1,10),
//          "i",
//        ),
//        ast.NewIntegerLiteralNode(loc(1,12),
//          "0",
//        ),
//      ),
//      ast.NewBinaryOpNode(loc(1,15),
//        "<",
//        ast.NewVariableNode(loc(1,15),
//          "i",
//        ),
//        ast.NewIntegerLiteralNode(loc(1,17),
//          "100",
//        ),
//      ),
//      ast.NewSuffixOpNode(loc(1,22),
//        "++",
//        ast.NewVariableNode(loc(1,22),
//          "i",
//        ),
//      ),
//      ast.NewExprStmtNode(loc(1,27),
//        ast.NewFuncallNode(loc(1,27),
//          ast.NewVariableNode(loc(1,27),
//            "println",
//          ),
//          []duck.IExprNode {
//            ast.NewVariableNode(loc(1,35),
//              "i",
//            ),
//          },
//        ),
//      ),
//    ),
//  },
    loc(2,5),
    ast.Declarations {
      Defvars: []entity.DefinedVariable { },
      Vardecls: []entity.UndefinedVariable { },
      Defuns: []entity.DefinedFunction { },
      Funcdecls: []entity.UndefinedFunction { },
      Constants: []entity.Constant { },
      Defstructs: []ast.StructNode { },
      Defunions: []ast.UnionNode { },
      Typedefs: []ast.TypedefNode { },
    },
  }
  y, err := ParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "for1", xt.JSON(y), xt.JSON(x))
  xt.AssertDeepEquals(t, "", y, x)
}
 */
