package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

type DereferenceChecker struct {
  errorHandler *core.ErrorHandler
  options *core.Options
  typeTable *typesys.TypeTable
}

func NewDereferenceChecker(errorHandler *core.ErrorHandler, options *core.Options, table *typesys.TypeTable) *DereferenceChecker {
  return &DereferenceChecker { errorHandler, options, table }
}

func (self *DereferenceChecker) Check(a *ast.AST) {
  self.errorHandler.Debug("starting dereference checker.")
  vs := a.GetDefinedVariables()
  for i := range vs {
    self.checkToplevelVariable(vs[i])
  }
  fs := a.GetDefinedFunctions()
  for i := range fs {
    ast.VisitStmtNode(self, fs[i].GetBody())
  }
  if self.errorHandler.ErrorOccured() {
    self.errorHandler.Fatalf("found %d error(s).", self.errorHandler.GetErrors())
  } else {
    self.errorHandler.Debug("finished dereference checker.")
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
    self.errorHandler.Errorf("%s not a constant", expr.GetLocation())
  }
}

func (self *DereferenceChecker) checkVariable(v *entity.DefinedVariable) {
  if v.HasInitializer() {
    ast.VisitExprNode(self, v.GetInitializer())
  }
}

func (self *DereferenceChecker) handleImplicitAddress(node core.IExprNode) {
  if ! node.IsLoadable() {
    t := node.GetType()
    if t.IsArray() {
      // int[4] ary; ary; should generate int*
      node.SetType(self.typeTable.PointerTo(t.GetBaseType()))
    } else {
      node.SetType(self.typeTable.PointerTo(t))
    }
  }
}

func (self *DereferenceChecker) checkMemberRef(loc core.Location, t core.IType, memb string) {
  if ! t.IsCompositeType() {
    self.errorHandler.Errorf("%s accessing member `%s' for non-struct/union: %s", loc, t, memb)
  }
  ct := t.(core.ICompositeType)
  if ! ct.HasMember(memb) {
    self.errorHandler.Errorf("%s %s does not have member: %s", loc, t, memb)
  }
}

func (self *DereferenceChecker) VisitStmtNode(unknown core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *ast.BlockNode: {
      vs := node.GetVariables()
      for i := range vs {
        self.checkVariable(vs[i])
      }
      ast.VisitStmtNodes(self, node.GetStmts())
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *DereferenceChecker) VisitExprNode(unknown core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *ast.AssignNode: {
      visitAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Errorf("%s invalid lhs expression", node.GetLocation())
      }
    }
    case *ast.OpAssignNode: {
      visitOpAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Errorf("%s invalid lhs expression", node.GetLocation())
      }
    }
    case *ast.PrefixOpNode: {
      visitPrefixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Errorf("%s cannot increment/decrement", node.GetExpr().GetLocation())
      }
    }
    case *ast.SuffixOpNode: {
      visitSuffixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Errorf("%s cannot increment/decrement", node.GetExpr().GetLocation())
      }
    }
    case *ast.FuncallNode: {
      visitFuncallNode(self, node)
      if ! node.GetExpr().IsCallable() {
        self.errorHandler.Errorf("%s calling object is not a function", node.GetLocation())
      }
    }
    case *ast.ArefNode: {
      visitArefNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s indexing non-array/pointer expression", node.GetLocation())
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
        self.errorHandler.Errorf("%s undereferable error", node.GetLocation())
      }
      self.checkMemberRef(node.GetLocation(), node.GetDereferedType(), node.GetMember())
      self.handleImplicitAddress(node)
    }
    case *ast.DereferenceNode: {
      visitDereferenceNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s undereferable error", node.GetLocation())
      }
      self.handleImplicitAddress(node)
    }
    case *ast.AddressNode: {
      visitAddressNode(self, node)
      if ! node.GetExpr().IsLvalue() {
        self.errorHandler.Errorf("%s invalid expression for &", node.GetLocation())
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
      self.handleImplicitAddress(node)
    }
    case *ast.CastNode: {
      visitCastNode(self, node)
      if node.GetType().IsArray() {
        self.errorHandler.Errorf("%s cast specifies array type", node.GetLocation())
      }
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *DereferenceChecker) VisitTypeDefinition(unknown core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
