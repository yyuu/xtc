package parser

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
  "bitbucket.org/yyuu/bs/xt"
)

func testParseExpr(source string) (*ast.AST, error) {
  return ParseExpr(source, core.NewErrorHandler(core.LOG_DEBUG), core.NewOptions("parser_test.go"))
}

func TestParseEmpty(t *testing.T) {
  _, err := testParseExpr("")
  xt.AssertNil(t, "", err)
}

func TestParseFuncallWithoutArguments(t *testing.T) {
  s := `
    int f() {
      return getc();
    }
  `
  x := ast.NewAST(loc(1,1),
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(2,5),
            typesys.NewFunctionTypeRef(
              typesys.NewIntTypeRef(loc(2,5)),
              typesys.NewParamTypeRefs(loc(2,10),
                []core.ITypeRef { },
                false,
              ),
            ),
          ),
          "f",
          entity.NewParams(loc(2,10),
            entity.NewParameters(),
            false,
          ),
          ast.NewBlockNode(loc(2,13),
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewReturnNode(loc(3,7),
                ast.NewFuncallNode(loc(3,14),
                  ast.NewVariableNode(loc(3,14),
                    "getc",
                  ),
                  []core.IExprNode { },
                ),
              ),
            },
          ),
        ),
      ),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "parse funcall w/o arguments", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestParseFuncallWithSingleArgument(t *testing.T) {
  s := `
    void f(int n) {
      println("hello, %d", n);
    }
  `
  x := ast.NewAST(loc(1,1),
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(2,5),
            typesys.NewFunctionTypeRef(
              typesys.NewVoidTypeRef(loc(2,5)),
              typesys.NewParamTypeRefs(loc(2,12),
                []core.ITypeRef {
                  typesys.NewIntTypeRef(loc(2,12)),
                },
                false,
              ),
            ),
          ),
          "f",
          entity.NewParams(loc(2,12),
            entity.NewParameters(
              entity.NewParameter(
                ast.NewTypeNode(loc(2,12),
                  typesys.NewIntTypeRef(loc(2,12)),
                ),
                "n",
              ),
            ),
            false,
          ),
          ast.NewBlockNode(loc(2,19),
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewExprStmtNode(loc(3,7),
                ast.NewFuncallNode(loc(3,7),
                  ast.NewVariableNode(loc(3,7), "println"),
                  []core.IExprNode {
                    ast.NewStringLiteralNode(loc(3,15), "\"hello, %d\""),
                    ast.NewVariableNode(loc(3,28), "n"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "parse funcall w/ single argument", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestParseDefunWithMultipleArguments(t *testing.T) {
  s := `

    int g(int x, int y) {
      int n = x * y;
      return n;
    }
  `
  x := ast.NewAST(loc(1,1),
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(3,5),
            typesys.NewFunctionTypeRef(
              typesys.NewIntTypeRef(loc(3,5)),
              typesys.NewParamTypeRefs(loc(3,11),
                []core.ITypeRef {
                  typesys.NewIntTypeRef(loc(3,11)),
                  typesys.NewIntTypeRef(loc(3,18)),
                },
                false,
              ),
            ),
          ),
          "g",
          entity.NewParams(loc(3,11),
            entity.NewParameters(
              entity.NewParameter(
                ast.NewTypeNode(loc(3,11),
                  typesys.NewIntTypeRef(loc(3,11)),
                ),
                "x",
              ),
              entity.NewParameter(
                ast.NewTypeNode(loc(3,18),
                  typesys.NewIntTypeRef(loc(3,18)),
                ),
                "y",
              ),
            ),
            false,
          ),
          ast.NewBlockNode(loc(3,25),
            entity.NewDefinedVariables(
              entity.NewDefinedVariable(
                true,
                ast.NewTypeNode(loc(4,7),
                  typesys.NewIntTypeRef(loc(4,7)),
                ),
                "n",
                ast.NewBinaryOpNode(loc(4,15),
                  "*",
                  ast.NewVariableNode(loc(4,15), "x"),
                  ast.NewVariableNode(loc(4,19), "y"),
                ),
              ),
            ),
            []core.IStmtNode {
              ast.NewReturnNode(loc(5,7),
                ast.NewVariableNode(loc(5,14), "n"),
              ),
            },
          ),
        ),
      ),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "defun w/ multiple arguments", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestParseDefunWithVariableArguments(t *testing.T) {
  s := `
    void myPrintf(char* fmt, ...) {
      _printf(fmt);
    }
  `
  x := ast.NewAST(loc(1,1),
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(2,5),
            typesys.NewFunctionTypeRef(
              typesys.NewVoidTypeRef(loc(2,5)),
              typesys.NewParamTypeRefs(loc(2,19),
                []core.ITypeRef {
                  typesys.NewPointerTypeRef(typesys.NewCharTypeRef(loc(2,19))),
                },
                true,
              ),
            ),
          ),
          "myPrintf",
          entity.NewParams(loc(2,19),
            entity.NewParameters(
              entity.NewParameter(
                ast.NewTypeNode(loc(2,19),
                  typesys.NewPointerTypeRef(typesys.NewCharTypeRef(loc(2,19))),
                ),
                "fmt",
              ),
            ),
            true,
          ),
          ast.NewBlockNode(loc(2,35),
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewExprStmtNode(loc(3,7),
                ast.NewFuncallNode(loc(3,7),
                  ast.NewVariableNode(loc(3,7), "_printf"),
                  []core.IExprNode {
                    ast.NewVariableNode(loc(3,15), "fmt"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "defun w/ variable arguments", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestFor1(t *testing.T) {
  s := `
    void f(int n) {
      for (i=0; i<n; i++) {
        s = sprintf("%d", i);
        println(s);
      }
    }
`
  x := ast.NewAST(loc(1,1),
    ast.NewDeclaration(
      entity.NewDefinedVariables(),
      entity.NewUndefinedVariables(),
      entity.NewDefinedFunctions(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(2,5),
            typesys.NewFunctionTypeRef(
              typesys.NewVoidTypeRef(loc(2,5)),
              typesys.NewParamTypeRefs(loc(2,12),
                []core.ITypeRef {
                  typesys.NewIntTypeRef(loc(2,12)),
                },
                false,
              ),
            ),
          ),
          "f",
          entity.NewParams(loc(2,12),
            entity.NewParameters(
              entity.NewParameter(
                ast.NewTypeNode(loc(2,12),
                  typesys.NewIntTypeRef(loc(2,12)),
                ),
                "n",
              ),
            ),
            false,
          ),
          ast.NewBlockNode(loc(2,19),
            entity.NewDefinedVariables(),
            []core.IStmtNode {
              ast.NewForNode(loc(3,7),
                ast.NewAssignNode(loc(3,12),
                  ast.NewVariableNode(loc(3,12), "i"),
                  ast.NewIntegerLiteralNode(loc(3,14), "0"),
                ),
                ast.NewBinaryOpNode(loc(3,17),
                  "<",
                  ast.NewVariableNode(loc(3,17), "i"),
                  ast.NewVariableNode(loc(3,19), "n"),
                ),
                ast.NewSuffixOpNode(loc(3,22),
                  "++",
                  ast.NewVariableNode(loc(3,22), "i"),
                ),
                ast.NewBlockNode(loc(3,27),
                  entity.NewDefinedVariables(),
                  []core.IStmtNode {
                    ast.NewExprStmtNode(loc(4,9),
                      ast.NewAssignNode(loc(4,9),
                        ast.NewVariableNode(loc(4,9), "s"),
                        ast.NewFuncallNode(loc(4,13),
                          ast.NewVariableNode(loc(4,13), "sprintf"),
                          []core.IExprNode {
                            ast.NewStringLiteralNode(loc(4,21), "\"%d\""),
                            ast.NewVariableNode(loc(4,27), "i"),
                          },
                        ),
                      ),
                    ),
                    ast.NewExprStmtNode(loc(5,9),
                      ast.NewFuncallNode(loc(5,9),
                        ast.NewVariableNode(loc(5,9), "println"),
                        []core.IExprNode {
                          ast.NewVariableNode(loc(5,17), "s"),
                        },
                      ),
                    ),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      entity.NewUndefinedFunctions(),
      entity.NewConstants(),
      ast.NewStructNodes(),
      ast.NewUnionNodes(),
      ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "for1", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}
