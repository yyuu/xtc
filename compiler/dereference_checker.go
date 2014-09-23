package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

type DereferenceChecker struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
}

func NewDereferenceChecker(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *DereferenceChecker {
  return &DereferenceChecker { errorHandler, table }
}

func (self *DereferenceChecker) Check(a *ast.AST) {
  vs := a.GetDefinedVariables()
  for i := range vs {
    self.checkToplevelVariable(vs[i])
  }
  fs := a.GetDefinedFunctions()
  for i := range fs {
    ast.VisitStmt(self, fs[i].GetBody())
  }
}

func (self *DereferenceChecker) checkToplevelVariable(v *entity.DefinedVariable) {
  self.checkVariable(v)
  if v.HasInitializer() {
    self.checkConstant(v.GetInitializer())
  }
}

func (self *DereferenceChecker) checkConstant(expr core.IExprNode) {
  if ! expr.IsConstant() {
    self.errorHandler.Errorf("%s not a constant\n", expr.GetLocation())
  }
}

func (self *DereferenceChecker) checkVariable(v *entity.DefinedVariable) {
  if v.HasInitializer() {
    ast.VisitExpr(self, v.GetInitializer())
  }
}

func (self *DereferenceChecker) handleImplicitAddress(node core.IExprNode) {
  if ! node.IsLoadable() {
    t := node.GetType()
    if t.IsArray() {
      node.SetType(self.typeTable.PointerTo(t.GetBaseType()))
    } else {
      node.SetType(self.typeTable.PointerTo(t))
    }
  }
}

func (self *DereferenceChecker) checkMemberRef(loc core.Location, t core.IType, memb string) {
  if ! t.IsCompositeType() {
    self.errorHandler.Errorf("%s accessing member `%s' for non-struct/union: %s\n", loc, t, memb)
  }
  ct := t.(core.ICompositeType)
  if ! ct.HasMember(memb) {
    self.errorHandler.Errorf("%s %s does not have member: %s\n", loc, t, memb)
  }
}

func (self *DereferenceChecker) VisitNode(node core.INode) {
  switch typed := node.(type) {
    case *ast.BlockNode: {
      vs := typed.GetVariables()
      for i := range vs {
        self.checkVariable(vs[i])
      }
      ast.VisitStmts(self, typed.GetStmts())
    }
    case *ast.AssignNode: {
      visitAssignNode(self, typed)
      if ! typed.GetLhs().IsAssignable() {
        self.errorHandler.Fatalf("%s invalid lhs expression\n", typed.GetLocation())
      }
    }
    case *ast.OpAssignNode: {
      visitOpAssignNode(self, typed)
      if ! typed.GetLhs().IsAssignable() {
        self.errorHandler.Fatalf("%s invalid lhs expression\n", typed.GetLocation())
      }
    }
    case *ast.PrefixOpNode: {
      visitPrefixOpNode(self, typed)
      if ! typed.GetExpr().IsAssignable() {
        self.errorHandler.Fatalf("%s cannot increment/decrement\n", typed.GetExpr().GetLocation())
      }
    }
    case *ast.SuffixOpNode: {
      visitSuffixOpNode(self, typed)
      if ! typed.GetExpr().IsAssignable() {
        self.errorHandler.Fatalf("%s cannot increment/decrement\n", typed.GetExpr().GetLocation())
      }
    }
    case *ast.FuncallNode: {
      visitFuncallNode(self, typed)
      if ! typed.GetExpr().IsCallable() {
        self.errorHandler.Fatalf("%s calling object is not a function\n", typed.GetLocation())
      }
    }
    case *ast.ArefNode: {
      visitArefNode(self, typed)
      if ! typed.GetExpr().IsPointer() {
        self.errorHandler.Fatalf("%s indexing non-array/pointer expression\n", typed.GetLocation())
      }
      self.handleImplicitAddress(typed)
    }
    case *ast.MemberNode: {
      visitMemberNode(self, typed)
      self.checkMemberRef(typed.GetLocation(), typed.GetExpr().GetType(), typed.GetMember())
      self.handleImplicitAddress(typed)
    }
    case *ast.PtrMemberNode: {
      visitPtrMemberNode(self, typed)
      if ! typed.GetExpr().IsPointer() {
        self.errorHandler.Fatalf("%s undereferable error\n", typed.GetLocation())
      }
      self.checkMemberRef(typed.GetLocation(), typed.GetDereferedType(), typed.GetMember())
      self.handleImplicitAddress(typed)
    }
    case *ast.DereferenceNode: {
      visitDereferenceNode(self, typed)
      if ! typed.GetExpr().IsPointer() {
        self.errorHandler.Fatalf("%s undereferable error\n", typed.GetLocation())
      }
      self.handleImplicitAddress(typed)
    }
    case *ast.AddressNode: {
      visitAddressNode(self, typed)
      if ! typed.GetExpr().IsLvalue() {
        self.errorHandler.Fatalf("%s invalid expression for &\n", typed.GetLocation())
      }
      base := typed.GetExpr().GetType()
      if ! typed.GetExpr().IsLoadable() {
        typed.SetType(base)
      } else {
        typed.SetType(self.typeTable.PointerTo(base))
      }
    }
    case *ast.VariableNode: {
      visitVariableNode(self, typed)
      if typed.GetEntity().IsConstant() {
        self.checkConstant(typed.GetEntity().(*entity.Constant).GetValue())
      }
      self.handleImplicitAddress(typed)
    }
    case *ast.CastNode: {
      visitCastNode(self, typed)
      if typed.GetType().IsArray() {
        self.errorHandler.Fatalf("%s cast specifies array type\n", typed.GetLocation())
      }
    }
    default: {
      visitNode(self, node)
    }
  }
}
