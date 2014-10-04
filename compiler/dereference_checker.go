package compiler

import (
  "fmt"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
)

type DereferenceChecker struct {
  errorHandler *xtc_core.ErrorHandler
  options *xtc_core.Options
  typeTable *xtc_typesys.TypeTable
}

func NewDereferenceChecker(errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options, table *xtc_typesys.TypeTable) *DereferenceChecker {
  return &DereferenceChecker { errorHandler, options, table }
}

func (self *DereferenceChecker) Check(ast *xtc_ast.AST) (*xtc_ast.AST, error) {
  vs := ast.GetDefinedVariables()
  for i := range vs {
    self.checkToplevelVariable(vs[i])
  }
  fs := ast.GetDefinedFunctions()
  for i := range fs {
    xtc_ast.VisitStmtNode(self, fs[i].GetBody())
  }
  if self.errorHandler.ErrorOccured() {
    return nil, fmt.Errorf("found %d error(s).", self.errorHandler.GetErrors())
  }
  return ast, nil
}

func (self *DereferenceChecker) checkToplevelVariable(v *xtc_entity.DefinedVariable) {
  self.checkVariable(v)
  if v.HasInitializer() {
    self.checkConstant(v.GetInitializer())
  }
}

func (self *DereferenceChecker) checkConstant(expr xtc_core.IExprNode) {
  if ! expr.IsConstant() {
    self.errorHandler.Errorf("%s not a constant", expr.GetLocation())
  }
}

func (self *DereferenceChecker) checkVariable(v *xtc_entity.DefinedVariable) {
  if v.HasInitializer() {
    xtc_ast.VisitExprNode(self, v.GetInitializer())
  }
}

func (self *DereferenceChecker) handleImplicitAddress(node xtc_core.IExprNode) {
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

func (self *DereferenceChecker) checkMemberRef(loc xtc_core.Location, t xtc_core.IType, memb string) {
  if ! t.IsCompositeType() {
    self.errorHandler.Errorf("%s accessing member `%s' for non-struct/union: %s", loc, t, memb)
  }
  ct := t.(xtc_core.ICompositeType)
  if ! ct.HasMember(memb) {
    self.errorHandler.Errorf("%s %s does not have member: %s", loc, t, memb)
  }
}

func (self *DereferenceChecker) VisitStmtNode(unknown xtc_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.BlockNode: {
      vs := node.GetVariables()
      for i := range vs {
        self.checkVariable(vs[i])
      }
      xtc_ast.VisitStmtNodes(self, node.GetStmts())
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *DereferenceChecker) VisitExprNode(unknown xtc_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.AssignNode: {
      visitAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Errorf("%s invalid lhs expression", node.GetLocation())
      }
    }
    case *xtc_ast.OpAssignNode: {
      visitOpAssignNode(self, node)
      if ! node.GetLHS().IsAssignable() {
        self.errorHandler.Errorf("%s invalid lhs expression", node.GetLocation())
      }
    }
    case *xtc_ast.PrefixOpNode: {
      visitPrefixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Errorf("%s cannot increment/decrement", node.GetExpr().GetLocation())
      }
    }
    case *xtc_ast.SuffixOpNode: {
      visitSuffixOpNode(self, node)
      if ! node.GetExpr().IsAssignable() {
        self.errorHandler.Errorf("%s cannot increment/decrement", node.GetExpr().GetLocation())
      }
    }
    case *xtc_ast.FuncallNode: {
      visitFuncallNode(self, node)
      if ! node.GetExpr().IsCallable() {
        self.errorHandler.Errorf("%s calling object is not a function", node.GetLocation())
      }
    }
    case *xtc_ast.ArefNode: {
      visitArefNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s indexing non-array/pointer expression", node.GetLocation())
      }
      self.handleImplicitAddress(node)
    }
    case *xtc_ast.MemberNode: {
      visitMemberNode(self, node)
      self.checkMemberRef(node.GetLocation(), node.GetExpr().GetType(), node.GetMember())
      self.handleImplicitAddress(node)
    }
    case *xtc_ast.PtrMemberNode: {
      visitPtrMemberNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s undereferable error", node.GetLocation())
      }
      self.checkMemberRef(node.GetLocation(), node.GetDereferedType(), node.GetMember())
      self.handleImplicitAddress(node)
    }
    case *xtc_ast.DereferenceNode: {
      visitDereferenceNode(self, node)
      if ! node.GetExpr().IsPointer() {
        self.errorHandler.Errorf("%s undereferable error", node.GetLocation())
      }
      self.handleImplicitAddress(node)
    }
    case *xtc_ast.AddressNode: {
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
    case *xtc_ast.VariableNode: {
      visitVariableNode(self, node)
      if node.GetEntity().IsConstant() {
        self.checkConstant(node.GetEntity().(*xtc_entity.Constant).GetValue())
      }
      self.handleImplicitAddress(node)
    }
    case *xtc_ast.CastNode: {
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

func (self *DereferenceChecker) VisitTypeDefinition(unknown xtc_core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
