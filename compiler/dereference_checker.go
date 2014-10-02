package compiler

import (
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
)

type DereferenceChecker struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
  typeTable *bs_typesys.TypeTable
}

func NewDereferenceChecker(errorHandler *bs_core.ErrorHandler, options *bs_core.Options, table *bs_typesys.TypeTable) *DereferenceChecker {
  return &DereferenceChecker { errorHandler, options, table }
}

func (self *DereferenceChecker) Check(ast *bs_ast.AST) {
  self.errorHandler.Debug("starting dereference checker.")
  vs := ast.GetDefinedVariables()
  for i := range vs {
    self.checkToplevelVariable(vs[i])
  }
  fs := ast.GetDefinedFunctions()
  for i := range fs {
    bs_ast.VisitStmtNode(self, fs[i].GetBody())
  }
  if self.errorHandler.ErrorOccured() {
    self.errorHandler.Fatalf("found %d error(s).", self.errorHandler.GetErrors())
  } else {
    self.errorHandler.Debug("finished dereference checker.")
  }
}

func (self *DereferenceChecker) checkToplevelVariable(v *bs_entity.DefinedVariable) {
  self.checkVariable(v)
  if v.HasInitializer() {
    self.checkConstant(v.GetInitializer())
  }
}

func (self *DereferenceChecker) checkConstant(expr bs_core.IExprNode) {
  if ! expr.IsConstant() {
    self.errorHandler.Errorf("%s not a constant", expr.GetLocation())
  }
}

func (self *DereferenceChecker) checkVariable(v *bs_entity.DefinedVariable) {
  if v.HasInitializer() {
    bs_ast.VisitExprNode(self, v.GetInitializer())
  }
}

func (self *DereferenceChecker) handleImplicitAddress(node bs_core.IExprNode) {
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

func (self *DereferenceChecker) checkMemberRef(loc bs_core.Location, t bs_core.IType, memb string) {
  if ! t.IsCompositeType() {
    self.errorHandler.Errorf("%s accessing member `%s' for non-struct/union: %s", loc, t, memb)
  }
  ct := t.(bs_core.ICompositeType)
  if ! ct.HasMember(memb) {
    self.errorHandler.Errorf("%s %s does not have member: %s", loc, t, memb)
  }
}

func (self *DereferenceChecker) VisitStmtNode(unknown bs_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.BlockNode: {
      vs := node.GetVariables()
      for i := range vs {
        self.checkVariable(vs[i])
      }
      bs_ast.VisitStmtNodes(self, node.GetStmts())
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *DereferenceChecker) VisitExprNode(unknown bs_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.AssignNode: {
      visitAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Errorf("%s invalid lhs expression", node.GetLocation())
      }
    }
    case *bs_ast.OpAssignNode: {
      visitOpAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Errorf("%s invalid lhs expression", node.GetLocation())
      }
    }
    case *bs_ast.PrefixOpNode: {
      visitPrefixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Errorf("%s cannot increment/decrement", node.GetExpr().GetLocation())
      }
    }
    case *bs_ast.SuffixOpNode: {
      visitSuffixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Errorf("%s cannot increment/decrement", node.GetExpr().GetLocation())
      }
    }
    case *bs_ast.FuncallNode: {
      visitFuncallNode(self, node)
      if ! node.GetExpr().IsCallable() {
        self.errorHandler.Errorf("%s calling object is not a function", node.GetLocation())
      }
    }
    case *bs_ast.ArefNode: {
      visitArefNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s indexing non-array/pointer expression", node.GetLocation())
      }
      self.handleImplicitAddress(node)
    }
    case *bs_ast.MemberNode: {
      visitMemberNode(self, node)
      self.checkMemberRef(node.GetLocation(), node.GetExpr().GetType(), node.GetMember())
      self.handleImplicitAddress(node)
    }
    case *bs_ast.PtrMemberNode: {
      visitPtrMemberNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s undereferable error", node.GetLocation())
      }
      self.checkMemberRef(node.GetLocation(), node.GetDereferedType(), node.GetMember())
      self.handleImplicitAddress(node)
    }
    case *bs_ast.DereferenceNode: {
      visitDereferenceNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s undereferable error", node.GetLocation())
      }
      self.handleImplicitAddress(node)
    }
    case *bs_ast.AddressNode: {
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
    case *bs_ast.VariableNode: {
      visitVariableNode(self, node)
      if node.GetEntity().IsConstant() {
        self.checkConstant(node.GetEntity().(*bs_entity.Constant).GetValue())
      }
      self.handleImplicitAddress(node)
    }
    case *bs_ast.CastNode: {
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

func (self *DereferenceChecker) VisitTypeDefinition(unknown bs_core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
