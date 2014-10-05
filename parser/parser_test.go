package parser

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func testParseExpr(source string) (*xtc_ast.AST, error) {
  return ParseExpr(source, xtc_core.NewErrorHandler(xtc_core.LOG_WARN), xtc_core.NewOptions("parser_test.go"))
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
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc(2,5),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc(2,5)),
              xtc_typesys.NewParamTypeRefs(loc(2,10),
                []xtc_core.ITypeRef { },
                false,
              ),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc(2,10),
            xtc_entity.NewParameters(),
            false,
          ),
          xtc_ast.NewBlockNode(loc(2,13),
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc(3,7),
                xtc_ast.NewFuncallNode(loc(3,14),
                  xtc_ast.NewVariableNode(loc(3,14),
                    "getc",
                  ),
                  []xtc_core.IExprNode { },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
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
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc(2,5),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc(2,5)),
              xtc_typesys.NewParamTypeRefs(loc(2,12),
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc(2,12)),
                },
                false,
              ),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc(2,12),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(2,12),
                  xtc_typesys.NewIntTypeRef(loc(2,12)),
                ),
                "n",
              ),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc(2,19),
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewExprStmtNode(loc(3,7),
                xtc_ast.NewFuncallNode(loc(3,7),
                  xtc_ast.NewVariableNode(loc(3,7), "println"),
                  []xtc_core.IExprNode {
                    xtc_ast.NewStringLiteralNode(loc(3,15), "hello, %d"),
                    xtc_ast.NewVariableNode(loc(3,28), "n"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
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
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc(3,5),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc(3,5)),
              xtc_typesys.NewParamTypeRefs(loc(3,11),
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc(3,11)),
                  xtc_typesys.NewIntTypeRef(loc(3,18)),
                },
                false,
              ),
            ),
          ),
          "g",
          xtc_entity.NewParams(loc(3,11),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(3,11),
                  xtc_typesys.NewIntTypeRef(loc(3,11)),
                ),
                "x",
              ),
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(3,18),
                  xtc_typesys.NewIntTypeRef(loc(3,18)),
                ),
                "y",
              ),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc(3,25),
            xtc_entity.NewDefinedVariables(
              xtc_entity.NewDefinedVariable(
                false,
                xtc_ast.NewTypeNode(loc(4,7),
                  xtc_typesys.NewIntTypeRef(loc(4,7)),
                ),
                "n",
                xtc_ast.NewBinaryOpNode(loc(4,15),
                  "*",
                  xtc_ast.NewVariableNode(loc(4,15), "x"),
                  xtc_ast.NewVariableNode(loc(4,19), "y"),
                ),
              ),
            ),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc(5,7),
                xtc_ast.NewVariableNode(loc(5,14), "n"),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
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
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc(2,5),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc(2,5)),
              xtc_typesys.NewParamTypeRefs(loc(2,19),
                []xtc_core.ITypeRef {
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc(2,19))),
                },
                true,
              ),
            ),
          ),
          "myPrintf",
          xtc_entity.NewParams(loc(2,19),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(2,19),
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc(2,19))),
                ),
                "fmt",
              ),
            ),
            true,
          ),
          xtc_ast.NewBlockNode(loc(2,35),
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewExprStmtNode(loc(3,7),
                xtc_ast.NewFuncallNode(loc(3,7),
                  xtc_ast.NewVariableNode(loc(3,7), "_printf"),
                  []xtc_core.IExprNode {
                    xtc_ast.NewVariableNode(loc(3,15), "fmt"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
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
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc(2,5),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc(2,5)),
              xtc_typesys.NewParamTypeRefs(loc(2,12),
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc(2,12)),
                },
                false,
              ),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc(2,12),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(2,12),
                  xtc_typesys.NewIntTypeRef(loc(2,12)),
                ),
                "n",
              ),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc(2,19),
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewForNode(loc(3,7),
                xtc_ast.NewAssignNode(loc(3,12),
                  xtc_ast.NewVariableNode(loc(3,12), "i"),
                  xtc_ast.NewIntegerLiteralNode(loc(3,14), "0"),
                ),
                xtc_ast.NewBinaryOpNode(loc(3,17),
                  "<",
                  xtc_ast.NewVariableNode(loc(3,17), "i"),
                  xtc_ast.NewVariableNode(loc(3,19), "n"),
                ),
                xtc_ast.NewSuffixOpNode(loc(3,22),
                  "++",
                  xtc_ast.NewVariableNode(loc(3,22), "i"),
                ),
                xtc_ast.NewBlockNode(loc(3,27),
                  xtc_entity.NewDefinedVariables(),
                  []xtc_core.IStmtNode {
                    xtc_ast.NewExprStmtNode(loc(4,9),
                      xtc_ast.NewAssignNode(loc(4,9),
                        xtc_ast.NewVariableNode(loc(4,9), "s"),
                        xtc_ast.NewFuncallNode(loc(4,13),
                          xtc_ast.NewVariableNode(loc(4,13), "sprintf"),
                          []xtc_core.IExprNode {
                            xtc_ast.NewStringLiteralNode(loc(4,21), "%d"),
                            xtc_ast.NewVariableNode(loc(4,27), "i"),
                          },
                        ),
                      ),
                    ),
                    xtc_ast.NewExprStmtNode(loc(5,9),
                      xtc_ast.NewFuncallNode(loc(5,9),
                        xtc_ast.NewVariableNode(loc(5,9), "println"),
                        []xtc_core.IExprNode {
                          xtc_ast.NewVariableNode(loc(5,17), "s"),
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
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "for1", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestIfWithElse(t *testing.T) {
  s := `
    int even_p(int n) {
      if (n % 2 == 0) {
        return 1;
      } else {
        return 0;
      }
    }
`
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc(2,5),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc(2,5)),
              xtc_typesys.NewParamTypeRefs(loc(2,12),
                xtc_typesys.NewTypeRefs(
                  xtc_typesys.NewIntTypeRef(loc(2,12)),
                ),
                false,
              ),
            ),
          ),
          "even_p",
          xtc_entity.NewParams(loc(2,16),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(2,16),
                  xtc_typesys.NewIntTypeRef(loc(2,16)),
                ),
                "n",
              ),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc(2,23),
            xtc_entity.NewDefinedVariables(),
            xtc_ast.NewStmtNodes(
              xtc_ast.NewIfNode(loc(3,7),
                xtc_ast.NewBinaryOpNode(loc(3,11),
                  "==",
                  xtc_ast.NewBinaryOpNode(loc(3,11),
                    "%",
                    xtc_ast.NewVariableNode(loc(3,11), "n"),
                    xtc_ast.NewIntegerLiteralNode(loc(3,15), "2"),
                  ),
                  xtc_ast.NewIntegerLiteralNode(loc(3,20), "0"),
                ),
                xtc_ast.NewBlockNode(loc(3,23),
                  xtc_entity.NewDefinedVariables(),
                  xtc_ast.NewStmtNodes(
                    xtc_ast.NewReturnNode(loc(4,9),
                      xtc_ast.NewIntegerLiteralNode(loc(4,16), "1"),
                    ),
                  ),
                ),
                xtc_ast.NewBlockNode(loc(5,14),
                  xtc_entity.NewDefinedVariables(),
                  xtc_ast.NewStmtNodes(
                    xtc_ast.NewReturnNode(loc(6,9),
                      xtc_ast.NewIntegerLiteralNode(loc(6,16), "0"),
                    ),
                  ),
                ),
              ),
            ),
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "if w/ else", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestIfWithoutElse(t *testing.T) {
  s := `
    void onEven(int n) {
      if (n % 2 == 0) {
        println("even");
      }
    }
`
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc(2,5),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc(2,5)),
              xtc_typesys.NewParamTypeRefs(loc(2,12),
                xtc_typesys.NewTypeRefs(
                  xtc_typesys.NewIntTypeRef(loc(2,12)),
                ),
                false,
              ),
            ),
          ),
          "onEven",
          xtc_entity.NewParams(loc(2,17),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(2,17),
                  xtc_typesys.NewIntTypeRef(loc(2,17)),
                ),
                "n",
              ),
            ),
            false,
          ),
          xtc_ast.NewBlockNode(loc(2,24),
            xtc_entity.NewDefinedVariables(),
            xtc_ast.NewStmtNodes(
              xtc_ast.NewIfNode(loc(3,7),
                xtc_ast.NewBinaryOpNode(loc(3,11),
                  "==",
                  xtc_ast.NewBinaryOpNode(loc(3,11),
                    "%",
                    xtc_ast.NewVariableNode(loc(3,11), "n"),
                    xtc_ast.NewIntegerLiteralNode(loc(3,15), "2"),
                  ),
                  xtc_ast.NewIntegerLiteralNode(loc(3,20), "0"),
                ),
                xtc_ast.NewBlockNode(loc(3,23),
                  xtc_entity.NewDefinedVariables(),
                  xtc_ast.NewStmtNodes(
                    xtc_ast.NewExprStmtNode(loc(4,9),
                      xtc_ast.NewFuncallNode(loc(4,9),
                        xtc_ast.NewVariableNode(loc(4,9), "println"),
                        xtc_ast.NewExprNodes(
                          xtc_ast.NewStringLiteralNode(loc(4,17), "even"),
                        ),
                      ),
                    ),
                  ),
                ),
                nil,
              ),
            ),
          ),
        ),
      ),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "if w/o else", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestUndefinedFunctions(t *testing.T) {
  s := `
    extern int printf(char* format, ...);
    extern int scanf(char* format, ...);
`
  x := xtc_ast.NewAST(loc(1,1),
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(
        xtc_entity.NewUndefinedFunction(
          xtc_ast.NewTypeNode(loc(2,12),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc(2,23)),
              xtc_typesys.NewParamTypeRefs(loc(2,23),
                xtc_typesys.NewTypeRefs(
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc(1,1))),
                ),
                true,
              ),
            ),
          ),
          "printf",
          xtc_entity.NewParams(loc(2,23),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(2,23),
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc(1,1))),
                ),
                "format",
              ),
            ),
            true,
          ),
        ),
        xtc_entity.NewUndefinedFunction(
          xtc_ast.NewTypeNode(loc(3,12),
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc(3,22)),
              xtc_typesys.NewParamTypeRefs(loc(3,22),
                xtc_typesys.NewTypeRefs(
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc(1,1))),
                ),
                true,
              ),
            ),
          ),
          "scanf",
          xtc_entity.NewParams(loc(3,22),
            xtc_entity.NewParameters(
              xtc_entity.NewParameter(
                xtc_ast.NewTypeNode(loc(3,22),
                  xtc_typesys.NewPointerTypeRef(xtc_typesys.NewCharTypeRef(loc(1,1))),
                ),
                "format",
              ),
            ),
            true,
          ),
        ),
      ),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "undefined functions", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}
