package compiler

import (
  "fmt"
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  "bitbucket.org/yyuu/xtc/xt"
)

func assertTypeResolved(t *testing.T, s string, a *xtc_ast.AST) {
  entities := a.ListEntities()
  for i := range entities {
    switch ent := entities[i].(type) {
      case *xtc_entity.Constant: {
        xt.AssertTrue(t, fmt.Sprintf("%s: constant `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
      }
      case *xtc_entity.DefinedVariable: {
        xt.AssertTrue(t, fmt.Sprintf("%s: variable `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
      }
      case *xtc_entity.UndefinedVariable: {
        xt.AssertTrue(t, fmt.Sprintf("%s: variable `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
      }
      case *xtc_entity.DefinedFunction: {
        xt.AssertTrue(t, fmt.Sprintf("%s: function `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
        params := ent.GetParameters()
        for i := range params {
          xt.AssertTrue(t, fmt.Sprintf("%s: parameter of function `%s' is not type-resolved", s, ent.GetName()), params[i].GetTypeNode().IsResolved())
        }
      }
      case *xtc_entity.UndefinedFunction: {
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
