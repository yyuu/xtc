package compiler

import (
  "fmt"
  "testing"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  "bitbucket.org/yyuu/xtc/xt"
)

func assertTypeResolved(t *testing.T, s string, a *xtc_ast.AST) {
  v := &typeVisitor { t }
  entities := a.ListEntities()
  for i := range entities {
    switch ent := entities[i].(type) {
      case *xtc_entity.Constant: {
        xt.AssertTrue(t, fmt.Sprintf("%s: constant `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
        xtc_ast.VisitExprNode(v, ent.GetValue())
      }
      case *xtc_entity.DefinedVariable: {
        xt.AssertTrue(t, fmt.Sprintf("%s: variable `%s' (%s) is not type-resolved", s, ent.GetName(), ent.GetTypeRef()), ent.GetTypeNode().IsResolved())
        if ent.HasInitializer() {
          xtc_ast.VisitExprNode(v, ent.GetInitializer())
        }
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
        xtc_ast.VisitStmtNode(v, ent.GetBody())
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

type typeVisitor struct {
  t *testing.T
}

func (self *typeVisitor) assertTrue(key string, got bool) {
  xt.AssertTrue(self.t, key, got)
}

func (self *typeVisitor) VisitStmtNode(unknown xtc_core.IStmtNode) interface{} {
  visitStmtNode(self, unknown)
  return nil
}

func (self *typeVisitor) VisitExprNode(unknown xtc_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.CastNode: {
      self.assertTrue("cast: type not resolved", node.GetTypeNode().IsResolved())
      visitCastNode(self, node)
    }
    case *xtc_ast.IntegerLiteralNode: {
      self.assertTrue("integer: type not resolved", node.GetTypeNode().IsResolved())
      visitIntegerLiteralNode(self, node)
    }
    case *xtc_ast.SizeofExprNode: {
      self.assertTrue("sizeof(expr): type not resolved", node.GetTypeNode().IsResolved())
      visitSizeofExprNode(self, node)
    }
    case *xtc_ast.SizeofTypeNode: {
      self.assertTrue("sizeof(type): type not resolved", node.GetTypeNode().IsResolved())
      self.assertTrue("sizeof(type): type not resolved", node.GetOperandTypeNode().IsResolved())
      visitSizeofTypeNode(self, node)
    }
    case *xtc_ast.StringLiteralNode: {
      self.assertTrue("string: type not resolved", node.GetTypeNode().IsResolved())
      visitStringLiteralNode(self, node)
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *typeVisitor) VisitTypeDefinition(unknown xtc_core.ITypeDefinition) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.StructNode: {
      self.assertTrue("struct: type not resolved", node.GetTypeNode().IsResolved())
      members := node.GetMembers()
      for i := range members {
        self.assertTrue("struct: type not resolved", members[i].GetTypeNode().IsResolved())
      }
      visitStructNode(self, node)
    }
    case *xtc_ast.TypedefNode: {
      self.assertTrue("typedef: type not resolved", node.GetTypeNode().IsResolved())
      self.assertTrue("typedef: type not resolved", node.GetRealTypeNode().IsResolved())
      visitTypedefNode(self, node)
    }
    case *xtc_ast.UnionNode: {
      self.assertTrue("union: type not resolved", node.GetTypeNode().IsResolved())
      members := node.GetMembers()
      for i := range members {
        self.assertTrue("union: type not resolved", members[i].GetTypeNode().IsResolved())
      }
      visitUnionNode(self, node)
    }
    default: {
      visitTypeDefinition(self, unknown)
    }
  }
  return nil
}
