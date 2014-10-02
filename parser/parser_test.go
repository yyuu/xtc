package parser

import (
  "testing"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
  "bitbucket.org/yyuu/bs/xt"
)

func testParseExpr(source string) (*bs_ast.AST, error) {
  return ParseExpr(source, bs_core.NewErrorHandler(bs_core.LOG_WARN), bs_core.NewOptions("parser_test.go"))
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
  x := bs_ast.NewAST(loc(1,1),
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc(2,5),
            bs_typesys.NewFunctionTypeRef(
              bs_typesys.NewIntTypeRef(loc(2,5)),
              bs_typesys.NewParamTypeRefs(loc(2,10),
                []bs_core.ITypeRef { },
                false,
              ),
            ),
          ),
          "f",
          bs_entity.NewParams(loc(2,10),
            bs_entity.NewParameters(),
            false,
          ),
          bs_ast.NewBlockNode(loc(2,13),
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewReturnNode(loc(3,7),
                bs_ast.NewFuncallNode(loc(3,14),
                  bs_ast.NewVariableNode(loc(3,14),
                    "getc",
                  ),
                  []bs_core.IExprNode { },
                ),
              ),
            },
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
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
  x := bs_ast.NewAST(loc(1,1),
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc(2,5),
            bs_typesys.NewFunctionTypeRef(
              bs_typesys.NewVoidTypeRef(loc(2,5)),
              bs_typesys.NewParamTypeRefs(loc(2,12),
                []bs_core.ITypeRef {
                  bs_typesys.NewIntTypeRef(loc(2,12)),
                },
                false,
              ),
            ),
          ),
          "f",
          bs_entity.NewParams(loc(2,12),
            bs_entity.NewParameters(
              bs_entity.NewParameter(
                bs_ast.NewTypeNode(loc(2,12),
                  bs_typesys.NewIntTypeRef(loc(2,12)),
                ),
                "n",
              ),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc(2,19),
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewExprStmtNode(loc(3,7),
                bs_ast.NewFuncallNode(loc(3,7),
                  bs_ast.NewVariableNode(loc(3,7), "println"),
                  []bs_core.IExprNode {
                    bs_ast.NewStringLiteralNode(loc(3,15), "hello, %d"),
                    bs_ast.NewVariableNode(loc(3,28), "n"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
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
  x := bs_ast.NewAST(loc(1,1),
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc(3,5),
            bs_typesys.NewFunctionTypeRef(
              bs_typesys.NewIntTypeRef(loc(3,5)),
              bs_typesys.NewParamTypeRefs(loc(3,11),
                []bs_core.ITypeRef {
                  bs_typesys.NewIntTypeRef(loc(3,11)),
                  bs_typesys.NewIntTypeRef(loc(3,18)),
                },
                false,
              ),
            ),
          ),
          "g",
          bs_entity.NewParams(loc(3,11),
            bs_entity.NewParameters(
              bs_entity.NewParameter(
                bs_ast.NewTypeNode(loc(3,11),
                  bs_typesys.NewIntTypeRef(loc(3,11)),
                ),
                "x",
              ),
              bs_entity.NewParameter(
                bs_ast.NewTypeNode(loc(3,18),
                  bs_typesys.NewIntTypeRef(loc(3,18)),
                ),
                "y",
              ),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc(3,25),
            bs_entity.NewDefinedVariables(
              bs_entity.NewDefinedVariable(
                false,
                bs_ast.NewTypeNode(loc(4,7),
                  bs_typesys.NewIntTypeRef(loc(4,7)),
                ),
                "n",
                bs_ast.NewBinaryOpNode(loc(4,15),
                  "*",
                  bs_ast.NewVariableNode(loc(4,15), "x"),
                  bs_ast.NewVariableNode(loc(4,19), "y"),
                ),
              ),
            ),
            []bs_core.IStmtNode {
              bs_ast.NewReturnNode(loc(5,7),
                bs_ast.NewVariableNode(loc(5,14), "n"),
              ),
            },
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
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
  x := bs_ast.NewAST(loc(1,1),
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc(2,5),
            bs_typesys.NewFunctionTypeRef(
              bs_typesys.NewVoidTypeRef(loc(2,5)),
              bs_typesys.NewParamTypeRefs(loc(2,19),
                []bs_core.ITypeRef {
                  bs_typesys.NewPointerTypeRef(bs_typesys.NewCharTypeRef(loc(2,19))),
                },
                true,
              ),
            ),
          ),
          "myPrintf",
          bs_entity.NewParams(loc(2,19),
            bs_entity.NewParameters(
              bs_entity.NewParameter(
                bs_ast.NewTypeNode(loc(2,19),
                  bs_typesys.NewPointerTypeRef(bs_typesys.NewCharTypeRef(loc(2,19))),
                ),
                "fmt",
              ),
            ),
            true,
          ),
          bs_ast.NewBlockNode(loc(2,35),
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewExprStmtNode(loc(3,7),
                bs_ast.NewFuncallNode(loc(3,7),
                  bs_ast.NewVariableNode(loc(3,7), "_printf"),
                  []bs_core.IExprNode {
                    bs_ast.NewVariableNode(loc(3,15), "fmt"),
                  },
                ),
              ),
            },
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
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
  x := bs_ast.NewAST(loc(1,1),
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc(2,5),
            bs_typesys.NewFunctionTypeRef(
              bs_typesys.NewVoidTypeRef(loc(2,5)),
              bs_typesys.NewParamTypeRefs(loc(2,12),
                []bs_core.ITypeRef {
                  bs_typesys.NewIntTypeRef(loc(2,12)),
                },
                false,
              ),
            ),
          ),
          "f",
          bs_entity.NewParams(loc(2,12),
            bs_entity.NewParameters(
              bs_entity.NewParameter(
                bs_ast.NewTypeNode(loc(2,12),
                  bs_typesys.NewIntTypeRef(loc(2,12)),
                ),
                "n",
              ),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc(2,19),
            bs_entity.NewDefinedVariables(),
            []bs_core.IStmtNode {
              bs_ast.NewForNode(loc(3,7),
                bs_ast.NewAssignNode(loc(3,12),
                  bs_ast.NewVariableNode(loc(3,12), "i"),
                  bs_ast.NewIntegerLiteralNode(loc(3,14), "0"),
                ),
                bs_ast.NewBinaryOpNode(loc(3,17),
                  "<",
                  bs_ast.NewVariableNode(loc(3,17), "i"),
                  bs_ast.NewVariableNode(loc(3,19), "n"),
                ),
                bs_ast.NewSuffixOpNode(loc(3,22),
                  "++",
                  bs_ast.NewVariableNode(loc(3,22), "i"),
                ),
                bs_ast.NewBlockNode(loc(3,27),
                  bs_entity.NewDefinedVariables(),
                  []bs_core.IStmtNode {
                    bs_ast.NewExprStmtNode(loc(4,9),
                      bs_ast.NewAssignNode(loc(4,9),
                        bs_ast.NewVariableNode(loc(4,9), "s"),
                        bs_ast.NewFuncallNode(loc(4,13),
                          bs_ast.NewVariableNode(loc(4,13), "sprintf"),
                          []bs_core.IExprNode {
                            bs_ast.NewStringLiteralNode(loc(4,21), "%d"),
                            bs_ast.NewVariableNode(loc(4,27), "i"),
                          },
                        ),
                      ),
                    ),
                    bs_ast.NewExprStmtNode(loc(5,9),
                      bs_ast.NewFuncallNode(loc(5,9),
                        bs_ast.NewVariableNode(loc(5,9), "println"),
                        []bs_core.IExprNode {
                          bs_ast.NewVariableNode(loc(5,17), "s"),
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
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
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
  x := bs_ast.NewAST(loc(1,1),
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc(2,5),
            bs_typesys.NewFunctionTypeRef(
              bs_typesys.NewIntTypeRef(loc(2,5)),
              bs_typesys.NewParamTypeRefs(loc(2,12),
                bs_typesys.NewTypeRefs(
                  bs_typesys.NewIntTypeRef(loc(2,12)),
                ),
                false,
              ),
            ),
          ),
          "even_p",
          bs_entity.NewParams(loc(2,16),
            bs_entity.NewParameters(
              bs_entity.NewParameter(
                bs_ast.NewTypeNode(loc(2,16),
                  bs_typesys.NewIntTypeRef(loc(2,16)),
                ),
                "n",
              ),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc(2,23),
            bs_entity.NewDefinedVariables(),
            bs_ast.NewStmtNodes(
              bs_ast.NewIfNode(loc(3,7),
                bs_ast.NewBinaryOpNode(loc(3,11),
                  "==",
                  bs_ast.NewBinaryOpNode(loc(3,11),
                    "%",
                    bs_ast.NewVariableNode(loc(3,11), "n"),
                    bs_ast.NewIntegerLiteralNode(loc(3,15), "2"),
                  ),
                  bs_ast.NewIntegerLiteralNode(loc(3,20), "0"),
                ),
                bs_ast.NewBlockNode(loc(3,23),
                  bs_entity.NewDefinedVariables(),
                  bs_ast.NewStmtNodes(
                    bs_ast.NewReturnNode(loc(4,9),
                      bs_ast.NewIntegerLiteralNode(loc(4,16), "1"),
                    ),
                  ),
                ),
                bs_ast.NewBlockNode(loc(5,14),
                  bs_entity.NewDefinedVariables(),
                  bs_ast.NewStmtNodes(
                    bs_ast.NewReturnNode(loc(6,9),
                      bs_ast.NewIntegerLiteralNode(loc(6,16), "0"),
                    ),
                  ),
                ),
              ),
            ),
          ),
        ),
      ),
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
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
  x := bs_ast.NewAST(loc(1,1),
    bs_ast.NewDeclaration(
      bs_entity.NewDefinedVariables(),
      bs_entity.NewUndefinedVariables(),
      bs_entity.NewDefinedFunctions(
        bs_entity.NewDefinedFunction(
          false,
          bs_ast.NewTypeNode(loc(2,5),
            bs_typesys.NewFunctionTypeRef(
              bs_typesys.NewVoidTypeRef(loc(2,5)),
              bs_typesys.NewParamTypeRefs(loc(2,12),
                bs_typesys.NewTypeRefs(
                  bs_typesys.NewIntTypeRef(loc(2,12)),
                ),
                false,
              ),
            ),
          ),
          "onEven",
          bs_entity.NewParams(loc(2,17),
            bs_entity.NewParameters(
              bs_entity.NewParameter(
                bs_ast.NewTypeNode(loc(2,17),
                  bs_typesys.NewIntTypeRef(loc(2,17)),
                ),
                "n",
              ),
            ),
            false,
          ),
          bs_ast.NewBlockNode(loc(2,24),
            bs_entity.NewDefinedVariables(),
            bs_ast.NewStmtNodes(
              bs_ast.NewIfNode(loc(3,7),
                bs_ast.NewBinaryOpNode(loc(3,11),
                  "==",
                  bs_ast.NewBinaryOpNode(loc(3,11),
                    "%",
                    bs_ast.NewVariableNode(loc(3,11), "n"),
                    bs_ast.NewIntegerLiteralNode(loc(3,15), "2"),
                  ),
                  bs_ast.NewIntegerLiteralNode(loc(3,20), "0"),
                ),
                bs_ast.NewBlockNode(loc(3,23),
                  bs_entity.NewDefinedVariables(),
                  bs_ast.NewStmtNodes(
                    bs_ast.NewExprStmtNode(loc(4,9),
                      bs_ast.NewFuncallNode(loc(4,9),
                        bs_ast.NewVariableNode(loc(4,9), "println"),
                        bs_ast.NewExprNodes(
                          bs_ast.NewStringLiteralNode(loc(4,17), "even"),
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
      bs_entity.NewUndefinedFunctions(),
      bs_entity.NewConstants(),
      bs_ast.NewStructNodes(),
      bs_ast.NewUnionNodes(),
      bs_ast.NewTypedefNodes(),
    ),
  )
  y, err := testParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "if w/o else", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}
