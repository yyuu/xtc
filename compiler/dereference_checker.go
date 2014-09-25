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

func (self *DereferenceChecker) VisitNode(unknown core.INode) {
  switch node := unknown.(type) {
    case *ast.BlockNode: {
      vs := node.GetVariables()
      for i := range vs {
        self.checkVariable(vs[i])
      }
      ast.VisitStmts(self, node.GetStmts())
    }
    case *ast.AssignNode: {
      visitAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Fatalf("%s invalid lhs expression\n", node.GetLocation())
      }
    }
    case *ast.OpAssignNode: {
      visitOpAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Fatalf("%s invalid lhs expression\n", node.GetLocation())
      }
    }
    case *ast.PrefixOpNode: {
      visitPrefixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Fatalf("%s cannot increment/decrement\n", node.GetExpr().GetLocation())
      }
    }
    case *ast.SuffixOpNode: {
      visitSuffixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Fatalf("%s cannot increment/decrement\n", node.GetExpr().GetLocation())
      }
    }
    case *ast.FuncallNode: {
      visitFuncallNode(self, node)
      if ! node.GetExpr().IsCallable() {
        self.errorHandler.Fatalf("%s calling object is not a function\n", node.GetLocation())
      }
    }
    case *ast.ArefNode: {
      visitArefNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Fatalf("%s indexing non-array/pointer expression\n", node.GetLocation())
      }
      self.handleImplicitAddress(node)
    }
    case *ast.MemberNode: {
      visitMemberNode(self, node)
      self.checkMemberRef(node.GetLocation(), node.GetExpr().GetType(), node.GetMember())
      self.handleImplicitAddress(node)
    }
    case *ast.PtrMemberNode: {
      visitPtrMemberNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Fatalf("%s undereferable error\n", node.GetLocation())
      }
      self.checkMemberRef(node.GetLocation(), node.GetDereferedType(), node.GetMember())
      self.handleImplicitAddress(node)
    }
    case *ast.DereferenceNode: {
      visitDereferenceNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Fatalf("%s undereferable error\n", node.GetLocation())
      }
      self.handleImplicitAddress(node)
    }
    case *ast.AddressNode: {
      visitAddressNode(self, node)
      if ! node.GetExpr().IsLvalue() {
        self.errorHandler.Fatalf("%s invalid expression for &\n", node.GetLocation())
      }
      base := node.GetExpr().GetType()
      if ! node.GetExpr().IsLoadable() {
        node.SetType(base)
      } else {
        node.SetType(self.typeTable.PointerTo(base))
      }
    }
    case *ast.VariableNode: {
      visitVariableNode(self, node)
      if node.GetEntity().IsConstant() {
        self.checkConstant(node.GetEntity().(*entity.Constant).GetValue())
      }
//    self.handleImplicitAddress(node)
    }
    case *ast.CastNode: {
      visitCastNode(self, node)
      if node.GetType().IsArray() {
        self.errorHandler.Fatalf("%s cast specifies array type\n", node.GetLocation())
      }
    }
    default: {
      visitNode(self, unknown)
    }
  }
}
