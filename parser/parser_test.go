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
  x := ast.NewAST(
    loc(1,1),
    ast.NewDeclarations(
      defvars(),
      vardecls(),
      defuns(
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
      funcdecls(),
      defconsts(),
      defstructs(),
      defunions(),
      typedefs(),
    ),
  )
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
  x := ast.NewAST {
    loc(1,5),
    ast.NewDeclarations(
      defvars(),
      vardecls(),
      defuns(
//      []duck.IStmtNode {
//        ast.NewExprStmtNode(loc(1,5),
//          ast.NewFuncallNode(loc(1,5),
//            ast.NewVariableNode(loc(1,5),
//              "println",
//            ),
//            []duck.IExprNode {
//              ast.NewStringLiteralNode(loc(1,13),
//                "\"hello, world\"",
//              ),
//            },
//          ),
//        ),
//      },
      ),
      funcdecls(),
      constants(),
      defstructs(),
      defunions(),
      typedefs(),
    ),
  )
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
  x := ast.NewAST(
    loc(2,5),
    ast.Declarations {
      defvars(),
      vardecls(),
      defuns(
//      []duck.IStmtNode {
//        ast.NewExprStmtNode(loc(2,5),
//          ast.NewFuncallNode(loc(2,5),
//            ast.NewVariableNode(loc(2,5),
//              "println",
//            ),
//            []duck.IExprNode {
//              ast.NewStringLiteralNode(loc(3,7),
//                "\"hello, %s\"",
//              ),
//              ast.NewStringLiteralNode(loc(4,7),
//                "\"world\"",
//              ),
//            },
//          ),
//        ),
//      },
      ),
      funcdecls(),
      constants(),
      defstructs(),
      defunions(),
      typedefs(),
    ),
  )
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
  x := ast.NewAST(
    loc(2,5),
    ast.NewDeclarations(
      defvars(),
      vardecls(),
      defuns(
//      []duck.IStmtNode {
//        ast.NewForNode(loc(1,5),
//          ast.NewAssignNode(loc(1,10),
//            ast.NewVariableNode(loc(1,10),
//              "i",
//            ),
//            ast.NewIntegerLiteralNode(loc(1,12),
//              "0",
//            ),
//          ),
//          ast.NewBinaryOpNode(loc(1,15),
//            "<",
//            ast.NewVariableNode(loc(1,15),
//              "i",
//            ),
//            ast.NewIntegerLiteralNode(loc(1,17),
//              "100",
//            ),
//          ),
//          ast.NewSuffixOpNode(loc(1,22),
//            "++",
//            ast.NewVariableNode(loc(1,22),
//              "i",
//            ),
//          ),
//          ast.NewExprStmtNode(loc(1,27),
//            ast.NewFuncallNode(loc(1,27),
//              ast.NewVariableNode(loc(1,27),
//                "println",
//              ),
//              []duck.IExprNode {
//                ast.NewVariableNode(loc(1,35),
//                  "i",
//                ),
//              },
//            ),
//          ),
//        ),
//      },
      ),
      funcdecls(),
      constants(),
      defstructs(),
      defunions(),
      typedefs(),
    },
  }
  y, err := ParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "for1", xt.JSON(y), xt.JSON(x))
  xt.AssertDeepEquals(t, "", y, x)
}
 */
