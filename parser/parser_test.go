package parser

import (
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
  "bitbucket.org/yyuu/bs/xt"
)

func TestParseEmpty(t *testing.T) {
  _, err := ParseExpr("")
  xt.AssertNil(t, "", err)
}

func TestParseFuncallWithoutArguments(t *testing.T) {
  s := `
    int f(void) {
      return getc();
    }
  `
  x := ast.NewAST(loc(0,0),
    ast.NewDeclarations(
      defvars(),
      vardecls(),
      defuns(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(1,5),
            typesys.NewFunctionTypeRef(
              typesys.NewIntegerTypeRef(loc(1,5), "int"),
              typesys.NewParamTypeRefs(loc(1,11),
                []duck.ITypeRef { },
                false,
              ),
            ),
          ),
          "f",
          entity.NewParams(loc(1,11),
            []duck.IParameter { },
          ),
          ast.NewBlockNode(loc(1,17),
            []duck.IDefinedVariable { },
            []duck.IStmtNode {
              ast.NewReturnNode(loc(2,7),
                ast.NewFuncallNode(loc(2,14),
                  ast.NewVariableNode(loc(2,14),
                    "getc",
                  ),
                  []duck.IExprNode { },
                ),
              ),
            },
          ),
        ),
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
//xt.AssertDeepEquals(t, "", y, x)
}

func TestParseFuncallWithSingleArgument(t *testing.T) {
  s := `
    void f(int n) {
      println("hello, %d", n);
    }
  `
  x := ast.NewAST(loc(0,0),
    ast.NewDeclarations(
      defvars(),
      vardecls(),
      defuns(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(1,5),
            typesys.NewFunctionTypeRef(
              typesys.NewVoidTypeRef(loc(1,5)),
              typesys.NewParamTypeRefs(loc(1,12),
                []duck.ITypeRef {
                  typesys.NewIntegerTypeRef(loc(1,12), "int"),
                },
                false,
              ),
            ),
          ),
          "f",
          entity.NewParams(loc(1,12),
            []duck.IParameter {
              entity.NewParameter(
                ast.NewTypeNode(loc(1,12),
                  typesys.NewIntegerTypeRef(loc(1,12), "int"),
                ),
                "n",
              ),
            },
          ),
          ast.NewBlockNode(loc(1,19),
            []duck.IDefinedVariable { },
            []duck.IStmtNode {
              ast.NewExprStmtNode(loc(2,7),
                ast.NewFuncallNode(loc(2,7),
                  ast.NewVariableNode(loc(2,7), "println"),
                  []duck.IExprNode {
                    ast.NewStringLiteralNode(loc(2,15), "\"hello, %d\""),
                    ast.NewVariableNode(loc(2,28), "n"),
                  },
                ),
              ),
            },
          ),
        ),
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
  xt.AssertStringEqualsDiff(t, "parse funcall w/ single argument", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}

func TestParseFuncallWithMultipleArguments(t *testing.T) {
  s := `

    int g(int x, int y) {
      int n = x * y;
      return n;
    }
  `
  x := ast.NewAST(loc(0,0),
    ast.NewDeclarations(
      defvars(),
      vardecls(),
      defuns(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(2,5),
            typesys.NewFunctionTypeRef(
              typesys.NewIntegerTypeRef(loc(2,5), "int"),
              typesys.NewParamTypeRefs(loc(2,11),
                []duck.ITypeRef {
                  typesys.NewIntegerTypeRef(loc(2,11), "int"),
                  typesys.NewIntegerTypeRef(loc(2,18), "int"),
                },
                false,
              ),
            ),
          ),
          "g",
          entity.NewParams(loc(2,11),
            []duck.IParameter {
              entity.NewParameter(
                ast.NewTypeNode(loc(2,11),
                  typesys.NewIntegerTypeRef(loc(2,11), "int"),
                ),
                "x",
              ),
              entity.NewParameter(
                ast.NewTypeNode(loc(2,18),
                  typesys.NewIntegerTypeRef(loc(2,18), "int"),
                ),
                "y",
              ),
            },
          ),
          ast.NewBlockNode(loc(2,25),
            []duck.IDefinedVariable {
              entity.NewDefinedVariable(
                true,
                ast.NewTypeNode(loc(3,7),
                  typesys.NewIntegerTypeRef(loc(3,7), "int"),
                ),
                "n",
                ast.NewBinaryOpNode(loc(3,15),
                  "*",
                  ast.NewVariableNode(loc(3,15), "x"),
                  ast.NewVariableNode(loc(3,19), "y"),
                ),
              ),
            },
            []duck.IStmtNode {
              ast.NewReturnNode(loc(4,7),
                ast.NewVariableNode(loc(4,14), "n"),
              ),
            },
          ),
        ),
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
  xt.AssertStringEqualsDiff(t, "funcall w/ multiple arguments", xt.JSON(y), xt.JSON(x))
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
  x := ast.NewAST(loc(0,0),
    ast.NewDeclarations(
      defvars(),
      vardecls(),
      defuns(
        entity.NewDefinedFunction(
          true,
          ast.NewTypeNode(loc(1,5),
            typesys.NewFunctionTypeRef(
              typesys.NewVoidTypeRef(loc(1,5)),
              typesys.NewParamTypeRefs(loc(1,12),
                []duck.ITypeRef {
                  typesys.NewIntegerTypeRef(loc(1,12), "int"),
                },
                false,
              ),
            ),
          ),
          "f",
          entity.NewParams(loc(1,12),
            []duck.IParameter {
              entity.NewParameter(
                ast.NewTypeNode(loc(1,12),
                  typesys.NewIntegerTypeRef(loc(1,12), "int"),
                ),
                "n",
              ),
            },
          ),
          ast.NewBlockNode(loc(1,19),
            []duck.IDefinedVariable { },
            []duck.IStmtNode {
              ast.NewForNode(loc(2,7),
                ast.NewAssignNode(loc(2,12),
                  ast.NewVariableNode(loc(2,12), "i"),
                  ast.NewIntegerLiteralNode(loc(2,14), "0"),
                ),
                ast.NewBinaryOpNode(loc(2,17),
                  "<",
                  ast.NewVariableNode(loc(2,17), "i"),
                  ast.NewVariableNode(loc(2,19), "n"),
                ),
                ast.NewSuffixOpNode(loc(2,22),
                  "++",
                  ast.NewVariableNode(loc(2,22), "i"),
                ),
                ast.NewBlockNode(loc(2,27),
                  []duck.IDefinedVariable { },
                  []duck.IStmtNode {
                    ast.NewExprStmtNode(loc(3,9),
                      ast.NewAssignNode(loc(3,9),
                        ast.NewVariableNode(loc(3,9), "s"),
                        ast.NewFuncallNode(loc(3,13),
                          ast.NewVariableNode(loc(3,13), "sprintf"),
                          []duck.IExprNode {
                            ast.NewStringLiteralNode(loc(3,21), "\"%d\""),
                            ast.NewVariableNode(loc(3,27), "i"),
                          },
                        ),
                      ),
                    ),
                    ast.NewExprStmtNode(loc(4,9),
                      ast.NewFuncallNode(loc(4,9),
                        ast.NewVariableNode(loc(4,9), "println"),
                        []duck.IExprNode {
                          ast.NewVariableNode(loc(4,17), "s"),
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
      funcdecls(),
      defconsts(),
      defstructs(),
      defunions(),
      typedefs(),
    ),
  )
  y, err := ParseExpr(s)
  xt.AssertNil(t, "", err)
  xt.AssertStringEqualsDiff(t, "for1", xt.JSON(y), xt.JSON(x))
//xt.AssertDeepEquals(t, "", y, x)
}
