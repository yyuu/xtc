package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
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
//types := a.ListTypes()
//entities := a.ListEntities()

//self.defineTypes(types)
//for i := range types {
//  ast.VisitNode(self, types[i])
//}
//for i := range entities {
//  entity.VisitNode(self, entities[i])
//}
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

func (self *TypeResolver) VisitNode(node core.INode) {
  fmt.Println("FIXME: TypeResolver#Visit called:", node)
//switch typed := unknown.(type) {
//  case ast.StructNode: {
//  }
//  case ast.UnionNode: {
//  }
//  case ast.TypedefNode: {
//  }
//  case entity.DefinedVariable: {
//  }
//  case entity.UndefinedVariable: {
//  }
//  case entity.DefinedFunction: {
//  }
//  case entity.UndefinedFunction: {
//  }
//  case ast.BlockNode: {
//  }
//  case ast.CastNode: {
//  }
//  case ast.SizeofExprNode: {
//  }
//  case ast.SizeofTypeNode: {
//  }
//  case ast.IntegerLiteralNode: {
//    self.bindType(typed.GetTypeNode())
//  }
//  case ast.StringLiteralNode: {
//    self.bindType(typed.GetTypeNode())
//  }
//}
}

func (self *TypeResolver) bindType(x core.ITypeNode) {
  fmt.Println("FIXME: TypeResolver#bindType called:", x)
}
