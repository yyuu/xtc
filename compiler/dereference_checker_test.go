package compiler

import (
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
  "bitbucket.org/yyuu/xtc/xt"
)

func setupDereferenceChecker(ast *xtc_ast.AST, table *xtc_typesys.TypeTable) *DereferenceChecker {
  errorHandler := xtc_core.NewErrorHandler(xtc_core.LOG_WARN)
  options := xtc_core.NewOptions("dereference_checker_test.go")

  resolver := NewTypeResolver(errorHandler, options, table)
  types := ast.ListTypes()
  entities := ast.ListEntities()
  resolver.defineTypes(types)
  xtc_ast.VisitTypeDefinitions(resolver, types)
  xtc_entity.VisitEntities(resolver, entities)
  if errorHandler.ErrorOccured() {
    panic("must not happen: test data is broken")
  }

  checker := NewDereferenceChecker(errorHandler, options, table)
  vs := ast.GetDefinedVariables()
  for i := range vs {
    checker.checkToplevelVariable(vs[i])
  }
  fs := ast.GetDefinedFunctions()
  for i := range fs {
    xtc_ast.VisitStmtNode(checker, fs[i].GetBody())
  }
  return checker
}

func TestDereferenceCheckerCheckConstant(t *testing.T) {
/*
  int a = 12345;
  int b = 67890;
 */
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(
        xtc_entity.NewDefinedVariable(
          false,
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
          "a",
          xtc_ast.NewIntegerLiteralNode(loc, "12345"),
        ),
        xtc_entity.NewDefinedVariable(
          false,
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
          "b",
          xtc_ast.NewIntegerLiteralNode(loc, "67890"),
        ),
      ),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupDereferenceChecker(ast, table)
  xt.AssertFalse(t, "int literal should be able to be used as initializer", checker.errorHandler.ErrorOccured())
}

func TestDereferenceCheckerCheckNoneConstant(t *testing.T) {
/*
  int a = 12345;
  int b = a;
 */
  defer func() {
    if r := recover(); r != "" {
      t.Skipf("FIXME: entity of toplevel variable should be resolved eventually: %s", r)
    }
  }()
  loc := xtc_core.NewLocation("", 0, 0)
  ast := xtc_ast.NewAST(loc,
    xtc_ast.NewDeclaration(
      xtc_entity.NewDefinedVariables(
        xtc_entity.NewDefinedVariable(
          false,
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
          "a",
          xtc_ast.NewIntegerLiteralNode(loc, "12345"),
        ),
        xtc_entity.NewDefinedVariable(
          false,
          xtc_ast.NewTypeNode(loc, xtc_typesys.NewIntTypeRef(loc)),
          "b",
          xtc_ast.NewVariableNode(loc, "a"),
        ),
      ),
      xtc_entity.NewUndefinedVariables(),
      xtc_entity.NewDefinedFunctions(),
      xtc_entity.NewUndefinedFunctions(),
      xtc_entity.NewConstants(),
      xtc_ast.NewStructNodes(),
      xtc_ast.NewUnionNodes(),
      xtc_ast.NewTypedefNodes(),
    ),
  )
  table := xtc_typesys.NewTypeTableFor(xtc_core.PLATFORM_X86_LINUX)
  checker := setupDereferenceChecker(ast, table)
  xt.AssertTrue(t, "variable should not be able to be used as initializer", checker.errorHandler.ErrorOccured())
}
