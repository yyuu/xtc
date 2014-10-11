package compiler

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func setupTypeChecker(ast *xtc_ast.AST, table *xtc_typesys.TypeTable) *TypeChecker {
  errorHandler := xtc_core.NewErrorHandler(xtc_core.LOG_WARN)
  options := xtc_core.NewOptions("type_checker_test.go")

  resolver := NewTypeResolver(errorHandler, options, table)
  types := ast.ListTypes()
  entities := ast.ListEntities()
  resolver.defineTypes(types)
  xtc_ast.VisitTypeDefinitions(resolver, types)
  xtc_entity.VisitEntities(resolver, entities)
  if errorHandler.ErrorOccured() {
    panic("must not happen: test data is broken")
  }

  checker := NewTypeChecker(errorHandler, options, table)
  vs := ast.GetDefinedVariables()
  for i := range vs {
    checker.checkVariable(vs[i])
  }
  fs := ast.GetDefinedFunctions()
  for i := range fs {
    checker.currentFunction = fs[i]
    checker.checkReturnType(fs[i])
    checker.checkParamTypes(fs[i])
    xtc_ast.VisitStmtNode(checker, fs[i].GetBody())
  }
  return checker
}

func TestTypeCheckerVoidVoid(t *testing.T) {
/*
  void f() {
    return;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc, []xtc_core.ITypeRef { }, false),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, nil),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertFalse(t, "should be able to use void as the return value of void function", checker.errorHandler.ErrorOccured())
}

func TestTypeCheckerVoidInt(t *testing.T) {
/*
   void f() {
     return 12345;
   }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc, []xtc_core.ITypeRef { }, false),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, xtc_ast.NewIntegerLiteralNode(loc, "12345")),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertTrue(t, "should not be able to use int as the return value of void function", checker.errorHandler.ErrorOccured()) // must fail
}

func TestTypeCheckerVoidString(t *testing.T) {
/*
  void f() {
    return "foo";
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewVoidTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc, []xtc_core.ITypeRef { }, false),
            ),
          ),
          "f",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, xtc_ast.NewStringLiteralNode(loc, "foo")),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertTrue(t, "should not be able to use string as the return value of void function", checker.errorHandler.ErrorOccured()) // must fail
}

func TestTypeCheckerIntVoid(t *testing.T) {
/*
  int g() {
    return;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "g",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, nil),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertTrue(t, "should not be able to use void as the return value of int function", checker.errorHandler.ErrorOccured()) // must fail
}

func TestTypeCheckerIntChar(t *testing.T) {
/*
  int g() {
    return 'a';
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "g",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, xtc_ast.NewCharacterLiteralNode(loc, "97")),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertFalse(t, "should be able to use char as the return value of int function", checker.errorHandler.ErrorOccured())
}

func TestTypeCheckerIntInt(t *testing.T) {
/*
  int g() {
    return 0;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "g",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, xtc_ast.NewIntegerLiteralNode(loc, "0")),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertFalse(t, "should be able to use int as the return value of int function", checker.errorHandler.ErrorOccured())
}

func TestTypeCheckerIntLong(t *testing.T) {
/*
  int g() {
    return 0L;
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "g",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, xtc_ast.NewIntegerLiteralNode(loc, "0L")),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertFalse(t, "should be able to use long as the return value of int function", checker.errorHandler.ErrorOccured())
}

func TestTypeCheckerIntString(t *testing.T) {
/*
  int g() {
    return "bar";
  }
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(
        xtc_entity.NewDefinedFunction(
          false,
          xtc_ast.NewTypeNode(loc,
            xtc_typesys.NewFunctionTypeRef(
              xtc_typesys.NewIntTypeRef(loc),
              xtc_typesys.NewParamTypeRefs(loc,
                []xtc_core.ITypeRef {
                  xtc_typesys.NewIntTypeRef(loc),
                },
                false,
              ),
            ),
          ),
          "g",
          xtc_entity.NewParams(loc, xtc_entity.NewParameters(), false),
          xtc_ast.NewBlockNode(loc,
            xtc_entity.NewDefinedVariables(),
            []xtc_core.IStmtNode {
              xtc_ast.NewReturnNode(loc, xtc_ast.NewStringLiteralNode(loc, "bar")),
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
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupTypeChecker(ast, table)
  xt.AssertTrue(t, "should not be able to use string as the return value of int function", checker.errorHandler.ErrorOccured())
}
