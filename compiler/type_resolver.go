package compiler

import (
  "fmt"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
)

type TypeResolver struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
  typeTable *bs_typesys.TypeTable
}

func NewTypeResolver(errorHandler *bs_core.ErrorHandler, options *bs_core.Options, table *bs_typesys.TypeTable) *TypeResolver {
  return &TypeResolver { errorHandler, options, table }
}

func (self *TypeResolver) Resolve(ast *bs_ast.AST) (*bs_ast.AST, error) {
  types := ast.ListTypes()
  entities := ast.ListEntities()
  self.defineTypes(types)
  bs_ast.VisitTypeDefinitions(self, types)
  bs_entity.VisitEntities(self, entities)
  if self.errorHandler.ErrorOccured() {
    return nil, fmt.Errorf("found %d error(s).", self.errorHandler.GetErrors())
  }
  return ast, nil
}

func (self *TypeResolver) defineTypes(deftypes []bs_core.ITypeDefinition) {
  for i := range deftypes {
    def := deftypes[i]
    if self.typeTable.IsDefined(def.GetTypeRef()) {
      self.errorHandler.Errorf("duplicated type definition: %s", def.GetTypeRef())
    }
    self.typeTable.PutType(def.GetTypeRef(), def.DefiningType())
  }
}

func (self *TypeResolver) bindType(n bs_core.ITypeNode) {
  if ! n.IsResolved() {
    ref := n.GetTypeRef()
    t := self.typeTable.GetType(ref)
    n.SetType(t)
  }
}

func (self *TypeResolver) resolveCompositeType(def bs_core.ICompositeTypeDefinition) {
  ref := def.GetTypeRef()
  ct, ok := self.typeTable.GetType(ref).(bs_core.ICompositeType)
  if ! ok {
    self.errorHandler.Errorf("cannot intern struct/union: %s", def.GetName())
  }
  members := ct.GetMembers()
  for i := range members {
    slot := members[i]
    self.bindType(slot.GetTypeNode())
  }
}

func (self *TypeResolver) resolveFunctionHeader(fun bs_core.IFunction, params []*bs_entity.Parameter) {
  self.bindType(fun.GetTypeNode())
  for i := range params {
    param := params[i]
    t := self.typeTable.GetParamType(param.GetTypeRef())
    param.GetTypeNode().SetType(t)
  }
}

func (self *TypeResolver) VisitStmtNode(unknown bs_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.BlockNode: {
      variables := node.GetVariables()
      for i := range variables {
        bs_entity.VisitEntity(self, variables[i])
      }
      bs_ast.VisitStmtNodes(self, node.GetStmts())
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *TypeResolver) VisitExprNode(unknown bs_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.CastNode: {
      self.bindType(node.GetTypeNode())
      visitCastNode(self, node)
    }
    case *bs_ast.IntegerLiteralNode: {
      self.bindType(node.GetTypeNode())
    }
    case *bs_ast.SizeofExprNode: {
      self.bindType(node.GetTypeNode())
      visitSizeofExprNode(self, node)
    }
    case *bs_ast.SizeofTypeNode: {
      self.bindType(node.GetOperandTypeNode())
      self.bindType(node.GetTypeNode())
      visitSizeofTypeNode(self, node)
    }
    case *bs_ast.StringLiteralNode: {
      self.bindType(node.GetTypeNode())
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *TypeResolver) VisitTypeDefinition(unknown bs_core.ITypeDefinition) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.StructNode: {
      self.resolveCompositeType(node)
    }
    case *bs_ast.TypedefNode: {
      self.bindType(node.GetTypeNode())
      self.bindType(node.GetRealTypeNode())
    }
    case *bs_ast.UnionNode: {
      self.resolveCompositeType(node)
    }
    default: {
      visitTypeDefinition(self, unknown)
    }
  }
  return nil
}

func (self *TypeResolver) VisitEntity(unknown bs_core.IEntity) interface{} {
  switch ent := unknown.(type) {
    case *bs_entity.DefinedVariable: {
      self.bindType(ent.GetTypeNode())
      if ent.HasInitializer() {
        bs_ast.VisitExprNode(self, ent.GetInitializer())
      }
    }
    case *bs_entity.UndefinedVariable: {
      self.bindType(ent.GetTypeNode())
    }
    case *bs_entity.Constant: {
      self.bindType(ent.GetTypeNode())
      bs_ast.VisitExprNode(self, ent.GetValue())
    }
    case *bs_entity.DefinedFunction: {
      self.resolveFunctionHeader(ent, ent.GetParameters())
      bs_ast.VisitStmtNode(self, ent.GetBody())
    }
    case *bs_entity.UndefinedFunction: {
      self.resolveFunctionHeader(ent, ent.GetParameters())
    }
    default: {
      visitEntity(self, unknown)
    }
  }
  return nil
}
