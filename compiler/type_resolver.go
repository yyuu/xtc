package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

type TypeResolver struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
}

func NewTypeResolver(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *TypeResolver {
  return &TypeResolver { errorHandler, table }
}

func (self *TypeResolver) Resolve(a *ast.AST) {
  self.errorHandler.Debug("starting type resolver.")
  types := a.ListTypes()
  entities := a.ListEntities()

  self.defineTypes(types)
  for i := range types {
    ast.VisitNode(self, types[i])
  }
  for i := range entities {
    entity.VisitEntity(self, entities[i])
  }
  self.errorHandler.Debug("finished type resolver.")
}

func (self *TypeResolver) defineTypes(deftypes []core.ITypeDefinition) {
  for i := range deftypes {
    def := deftypes[i]
    if self.typeTable.IsDefined(def.GetTypeRef()) {
      panic(fmt.Errorf("duplicated type definition: %s", def.GetTypeRef()))
    }
    self.typeTable.PutType(def.GetTypeRef(), def.DefiningType())
  }
}

func (self *TypeResolver) bindType(n core.ITypeNode) {
  if ! n.IsResolved() {
    ref := n.GetTypeRef()
    t := self.typeTable.GetType(ref)
    n.SetType(t)
  }
}

func (self *TypeResolver) resolveCompositeType(def core.ICompositeTypeDefinition) {
  ref := def.GetTypeRef()
  ct, ok := self.typeTable.GetType(ref).(core.ICompositeType)
  if ! ok {
    panic(fmt.Errorf("cannot intern struct/union: %s", def.GetName()))
  }
  members := ct.GetMembers()
  for i := range members {
    slot := members[i]
    self.bindType(slot.GetTypeNode())
  }
}

func (self *TypeResolver) resolveFunctionHeader(fun core.IFunction, params []*entity.Parameter) {
  self.bindType(fun.GetTypeNode())
  for i := range params {
    param := params[i]
    t := self.typeTable.GetParamType(param.GetTypeRef())
    param.GetTypeNode().SetType(t)
  }
}

func (self *TypeResolver) VisitNode(unknown core.INode) interface{} {
  switch node := unknown.(type) {
    case *ast.BlockNode: {
      variables := node.GetVariables()
      for i := range variables {
        entity.VisitEntity(self, variables[i])
      }
      ast.VisitStmts(self, node.GetStmts())
    }
    case *ast.CastNode: {
      self.bindType(node.GetTypeNode())
      visitCastNode(self, node)
    }
    case *ast.IntegerLiteralNode: {
      self.bindType(node.GetTypeNode())
    }
    case *ast.SizeofExprNode: {
      self.bindType(node.GetTypeNode())
      visitSizeofExprNode(self, node)
    }
    case *ast.SizeofTypeNode: {
      self.bindType(node.GetOperandTypeNode())
      self.bindType(node.GetTypeNode())
      visitSizeofTypeNode(self, node)
    }
    case *ast.StringLiteralNode: {
      self.bindType(node.GetTypeNode())
    }
    case *ast.StructNode: {
      self.resolveCompositeType(node)
    }
    case *ast.TypedefNode: {
      self.bindType(node.GetTypeNode())
      self.bindType(node.GetRealTypeNode())
    }
    case *ast.UnionNode: {
      self.resolveCompositeType(node)
    }
    default: {
      visitNode(self, unknown)
    }
  }
  return nil
}

func (self *TypeResolver) VisitEntity(unknown core.IEntity) interface{} {
  switch ent := unknown.(type) {
    case *entity.DefinedVariable: {
      self.bindType(ent.GetTypeNode())
      if ent.HasInitializer() {
        ast.VisitNode(self, ent.GetInitializer())
      }
    }
    case *entity.UndefinedVariable: {
      self.bindType(ent.GetTypeNode())
    }
    case *entity.Constant: {
      self.bindType(ent.GetTypeNode())
      ast.VisitExpr(self, ent.GetValue())
    }
    case *entity.DefinedFunction: {
      self.resolveFunctionHeader(ent, ent.GetParameters())
      ast.VisitStmt(self, ent.GetBody())
    }
    case *entity.UndefinedFunction: {
      self.resolveFunctionHeader(ent, ent.GetParameters())
    }
    default: {
      visitEntity(self, unknown)
    }
  }
  return nil
}
