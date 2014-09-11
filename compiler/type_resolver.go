package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/typesys"
)

type TypeResolver struct {
  typeTable *typesys.TypeTable
}

func NewTypeResolver(table *typesys.TypeTable) *TypeResolver {
  return &TypeResolver { table }
}

func (self *TypeResolver) Resolve(a *ast.AST) {
  types := a.ListTypes()
  entities := a.ListEntities()

  self.defineTypes(types)
  for i := range types {
    ast.VisitNode(self, types[i])
  }
  for i := range entities {
    entity.VisitEntity(self, entities[i])
  }
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

func (self *TypeResolver) VisitNode(node core.INode) {
  switch typed := node.(type) {
    case *ast.StructNode: {
      self.resolveCompositeType(typed)
    }
    case *ast.UnionNode: {
      self.resolveCompositeType(typed)
    }
    case *ast.TypedefNode: {
      self.bindType(typed.GetTypeNode())
      self.bindType(typed.GetRealTypeNode())
    }
    case *ast.BlockNode: {
      variables := typed.GetVariables()
      for i := range variables {
        entity.VisitEntity(self, variables[i])
      }
      stmts := typed.GetStmts()
      for i := range stmts {
        ast.VisitNode(self, stmts[i])
      }
    }
    case *ast.CastNode: {
      self.bindType(typed.GetTypeNode())
      // super
    }
    case *ast.SizeofExprNode: {
      self.bindType(typed.GetTypeNode())
      // super
    }
    case *ast.SizeofTypeNode: {
      self.bindType(typed.GetOperandTypeNode())
      self.bindType(typed.GetTypeNode())
      // super
    }
    case *ast.IntegerLiteralNode: {
      self.bindType(typed.GetTypeNode())
    }
    case *ast.StringLiteralNode: {
      self.bindType(typed.GetTypeNode())
    }
  }
}

func (self *TypeResolver) VisitEntity(ent core.IEntity) {
  switch typed := ent.(type) {
    case *entity.DefinedVariable: {
      self.bindType(typed.GetTypeNode())
      if typed.HasInitializer() {
        ast.VisitNode(self, typed.GetInitializer())
      }
    }
    case *entity.UndefinedVariable: {
      self.bindType(typed.GetTypeNode())
    }
    case *entity.DefinedFunction: {
      self.resolveFunctionHeader(typed, typed.GetParameters())
      ast.VisitNode(self, typed.GetBody())
    }
    case *entity.UndefinedFunction: {
      self.resolveFunctionHeader(typed, typed.GetParameters())
    }
  }
}
