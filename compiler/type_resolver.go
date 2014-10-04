package compiler

import (
  "fmt"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
)

type TypeResolver struct {
  errorHandler *xtc_core.ErrorHandler
  options *xtc_core.Options
  typeTable *xtc_typesys.TypeTable
}

func NewTypeResolver(errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options, table *xtc_typesys.TypeTable) *TypeResolver {
  return &TypeResolver { errorHandler, options, table }
}

func (self *TypeResolver) Resolve(ast *xtc_ast.AST) (*xtc_ast.AST, error) {
  types := ast.ListTypes()
  entities := ast.ListEntities()
  self.defineTypes(types)
  xtc_ast.VisitTypeDefinitions(self, types)
  xtc_entity.VisitEntities(self, entities)
  if self.errorHandler.ErrorOccured() {
    return nil, fmt.Errorf("found %d error(s).", self.errorHandler.GetErrors())
  }
  return ast, nil
}

func (self *TypeResolver) defineTypes(deftypes []xtc_core.ITypeDefinition) {
  for i := range deftypes {
    def := deftypes[i]
    if self.typeTable.IsDefined(def.GetTypeRef()) {
      self.errorHandler.Errorf("duplicated type definition: %s", def.GetTypeRef())
    }
    self.typeTable.PutType(def.GetTypeRef(), def.DefiningType())
  }
}

func (self *TypeResolver) bindType(n xtc_core.ITypeNode) {
  if ! n.IsResolved() {
    ref := n.GetTypeRef()
    t := self.typeTable.GetType(ref)
    n.SetType(t)
  }
}

func (self *TypeResolver) resolveCompositeType(def xtc_core.ICompositeTypeDefinition) {
  ref := def.GetTypeRef()
  ct, ok := self.typeTable.GetType(ref).(xtc_core.ICompositeType)
  if ! ok {
    self.errorHandler.Errorf("cannot intern struct/union: %s", def.GetName())
  }
  members := ct.GetMembers()
  for i := range members {
    slot := members[i]
    self.bindType(slot.GetTypeNode())
  }
}

func (self *TypeResolver) resolveFunctionHeader(fun xtc_core.IFunction, params []*xtc_entity.Parameter) {
  self.bindType(fun.GetTypeNode())
  for i := range params {
    param := params[i]
    t := self.typeTable.GetParamType(param.GetTypeRef())
    param.GetTypeNode().SetType(t)
  }
}

func (self *TypeResolver) VisitStmtNode(unknown xtc_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.BlockNode: {
      variables := node.GetVariables()
      for i := range variables {
        xtc_entity.VisitEntity(self, variables[i])
      }
      xtc_ast.VisitStmtNodes(self, node.GetStmts())
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *TypeResolver) VisitExprNode(unknown xtc_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.CastNode: {
      self.bindType(node.GetTypeNode())
      visitCastNode(self, node)
    }
    case *xtc_ast.IntegerLiteralNode: {
      self.bindType(node.GetTypeNode())
    }
    case *xtc_ast.SizeofExprNode: {
      self.bindType(node.GetTypeNode())
      visitSizeofExprNode(self, node)
    }
    case *xtc_ast.SizeofTypeNode: {
      self.bindType(node.GetOperandTypeNode())
      self.bindType(node.GetTypeNode())
      visitSizeofTypeNode(self, node)
    }
    case *xtc_ast.StringLiteralNode: {
      self.bindType(node.GetTypeNode())
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *TypeResolver) VisitTypeDefinition(unknown xtc_core.ITypeDefinition) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.StructNode: {
      self.resolveCompositeType(node)
    }
    case *xtc_ast.TypedefNode: {
      self.bindType(node.GetTypeNode())
      self.bindType(node.GetRealTypeNode())
    }
    case *xtc_ast.UnionNode: {
      self.resolveCompositeType(node)
    }
    default: {
      visitTypeDefinition(self, unknown)
    }
  }
  return nil
}

func (self *TypeResolver) VisitEntity(unknown xtc_core.IEntity) interface{} {
  switch ent := unknown.(type) {
    case *xtc_entity.DefinedVariable: {
      self.bindType(ent.GetTypeNode())
      if ent.HasInitializer() {
        xtc_ast.VisitExprNode(self, ent.GetInitializer())
      }
    }
    case *xtc_entity.UndefinedVariable: {
      self.bindType(ent.GetTypeNode())
    }
    case *xtc_entity.Constant: {
      self.bindType(ent.GetTypeNode())
      xtc_ast.VisitExprNode(self, ent.GetValue())
    }
    case *xtc_entity.DefinedFunction: {
      self.resolveFunctionHeader(ent, ent.GetParameters())
      xtc_ast.VisitStmtNode(self, ent.GetBody())
    }
    case *xtc_entity.UndefinedFunction: {
      self.resolveFunctionHeader(ent, ent.GetParameters())
    }
    default: {
      visitEntity(self, unknown)
    }
  }
  return nil
}
