package compiler

import (
  "fmt"
  "testing"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/xt"
)

func assertTypeResolved(t *testing.T, s string, a *ast.AST) {
  entities := a.ListEntities()
  for i := range entities {
    switch ent := entities[i].(type) {
      case *entity.Constant: {
        xt.AssertTrue(t, fmt.Sprintf("%s: constant `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
      }
      case *entity.DefinedVariable: {
        xt.AssertTrue(t, fmt.Sprintf("%s: variable `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
      }
      case *entity.UndefinedVariable: {
        xt.AssertTrue(t, fmt.Sprintf("%s: variable `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
      }
      case *entity.DefinedFunction: {
        xt.AssertTrue(t, fmt.Sprintf("%s: function `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
        params := ent.GetParameters()
        for i := range params {
          xt.AssertTrue(t, fmt.Sprintf("%s: parameter of function `%s' is not type-resolved", s, ent.GetName()), params[i].GetTypeNode().IsResolved())
        }
      }
      case *entity.UndefinedFunction: {
        xt.AssertTrue(t, fmt.Sprintf("%s: function `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
        params := ent.GetParameters()
        for i := range params {
          xt.AssertTrue(t, fmt.Sprintf("%s: parameter of function `%s' is not type-resolved", s, ent.GetName()), params[i].GetTypeNode().IsResolved())
        }
      }
      default: {
        xt.AssertTrue(t, fmt.Sprintf("%s: unknown (%s) is not type-resolved", s, ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
      }
    }
  }
}
