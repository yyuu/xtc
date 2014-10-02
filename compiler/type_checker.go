package compiler

import (
  "fmt"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
)

type TypeChecker struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
  typeTable *bs_typesys.TypeTable
  currentFunction *bs_entity.DefinedFunction
}

func NewTypeChecker(errorHandler *bs_core.ErrorHandler, options *bs_core.Options, table *bs_typesys.TypeTable) *TypeChecker {
  return &TypeChecker { errorHandler, options, table, nil }
}

func (self *TypeChecker) Check(ast *bs_ast.AST) (*bs_ast.AST, error) {
  vs := ast.GetDefinedVariables()
  for i := range vs {
    self.checkVariable(vs[i])
  }
  fs := ast.GetDefinedFunctions()
  for i := range fs {
    self.currentFunction = fs[i]
    self.checkReturnType(fs[i])
    self.checkParamTypes(fs[i])
    bs_ast.VisitStmtNode(self, fs[i].GetBody())
  }
  if self.errorHandler.ErrorOccured() {
    return nil, fmt.Errorf("found %d error(s).", self.errorHandler.GetErrors())
  }
  return ast, nil
}

func (self *TypeChecker) checkVariable(v *bs_entity.DefinedVariable) {
  if self.isInvalidVariableType(v.GetType()) {
    self.errorHandler.Errorf("invalid variable type: %s", v.GetType())
  }
  if v.HasInitializer() {
    if self.isInvalidLHSType(v.GetType()) {
      self.errorHandler.Errorf("invalid LHS type: %s", v.GetType())
    }
    bs_ast.VisitExprNode(self, v.GetInitializer())
    v.SetInitializer(self.implicitCast(v.GetType(), v.GetInitializer()))
  }
}

func (self *TypeChecker) isInvalidVariableType(t bs_core.IType) bool {
  return t.IsVoid() || (t.IsArray() && ! t.IsAllocatedArray())
}

func (self *TypeChecker) isInvalidLHSType(t bs_core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid() || t.IsArray()
}

func (self *TypeChecker) isInvalidRHSType(t bs_core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid()
}

func (self *TypeChecker) implicitCast(t bs_core.IType, expr bs_core.IExprNode) bs_core.IExprNode {
  if expr.GetType().IsSameType(t) {
    return expr
  } else {
    if expr.GetType().IsCastableTo(t) {
      if ! expr.GetType().IsCompatible(t) && ! self.isSafeIntegerCast(expr, t) {
        self.errorHandler.Errorf("%s incompatible implicit cast from %s to %s", expr.GetLocation(), expr.GetType(), t)
      }
      typeNode := bs_ast.NewTypeNode(expr.GetLocation(), bs_typesys.NewVoidTypeRef(expr.GetLocation()))
      typeNode.SetType(t)
      return bs_ast.NewCastNode(expr.GetLocation(), typeNode, expr)
    } else {
      self.errorHandler.Errorf("invalid cast error: %s to %s", expr.GetType(), t)
      return expr
    }
  }
}

func (self *TypeChecker) castOptionalArg(arg bs_core.IExprNode) bs_core.IExprNode {
  if ! arg.GetType().IsInteger() {
    return arg
  } else {
    var t bs_core.IType
    if arg.GetType().IsSigned() {
      t = self.typeTable.SignedStackType()
    } else {
      t = self.typeTable.UnsignedStackType()
    }
    if arg.GetType().Size() < t.Size() {
      return self.implicitCast(t, arg)
    } else {
      return arg
    }
  }
}

func (self *TypeChecker) isSafeIntegerCast(node bs_core.INode, t bs_core.IType) bool {
  if ! t.IsInteger() {
    return false
  } else {
    i, ok := t.(bs_typesys.IntegerType)
    if ! ok {
      return false
    }
    n, ok := node.(bs_ast.IntegerLiteralNode)
    if ! ok {
      return false
    }
    return i.IsInDomain(n.GetValue())
  }
}

func (self *TypeChecker) checkReturnType(f *bs_entity.DefinedFunction) {
  if self.isInvalidReturnType(f.GetReturnType()) {
    self.errorHandler.Errorf("returns invalid type: %s", f.GetReturnType())
  }
}

func (self *TypeChecker) isInvalidReturnType(t bs_core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsArray()
}

func (self *TypeChecker) checkParamTypes(f *bs_entity.DefinedFunction) {
  params := f.GetParameters()
  for i := range params {
    param := params[i]
    if self.isInvalidParameterType(param.GetType()) {
      self.errorHandler.Errorf("invalid parameter type: %s", param.GetType())
    }
  }
}

func (self *TypeChecker) isInvalidParameterType(t bs_core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid() || t.IsIncompleteArray()
}

func (self *TypeChecker) isInvalidStatementType(t bs_core.IType) bool {
  return t.IsStruct() || t.IsUnion()
}

func (self *TypeChecker) mustBeInteger(expr bs_core.IExprNode, op string) bool {
  if ! expr.GetType().IsInteger() {
    self.errorHandler.Errorf("%s wrong operand type for %s: %s", expr.GetLocation(), op, expr.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) mustBeScalar(expr bs_core.IExprNode, op string) bool {
  if ! expr.GetType().IsScalar() {
    self.errorHandler.Errorf("%s wrong operand type for %s: %s", expr.GetLocation(), op, expr.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) checkCond(cond bs_core.IExprNode) {
  self.mustBeScalar(cond, "condition expression")
}

func (self *TypeChecker) expectsComparableScalars(node bs_core.IBinaryOpNode) {
  if ! self.mustBeScalar(node.GetLeft(), node.GetOperator()) {
    return
  }
  if ! self.mustBeScalar(node.GetRight(), node.GetOperator()) {
    return
  }
  if node.GetLeft().GetType().IsPointer() {
    right := self.forcePointerType(node.GetLeft(), node.GetRight())
    node.SetRight(right)
    node.SetType(node.GetLeft().GetType())
    return
  }
  if node.GetRight().GetType().IsPointer() {
    left := self.forcePointerType(node.GetRight(), node.GetLeft())
    node.SetLeft(left)
    node.SetType(node.GetRight().GetType())
    return
  }
  self.arithmeticImplicitCast(node)
}

func (self *TypeChecker) forcePointerType(master bs_core.IExprNode, slave bs_core.IExprNode) bs_core.IExprNode {
  if master.GetType().IsCompatible(slave.GetType()) {
    return slave
  } else {
    self.errorHandler.Warnf("incompatible implicit cast from %s to %s", slave.GetType(), master.GetType())
    typeNode := bs_ast.NewTypeNode(master.GetLocation(), bs_typesys.NewVoidTypeRef(master.GetLocation()))
    typeNode.SetType(master.GetType())
    return bs_ast.NewCastNode(master.GetLocation(), typeNode, slave)
  }
}

func (self *TypeChecker) arithmeticImplicitCast(node bs_core.IBinaryOpNode) {
  r := self.integralPromotion(node.GetRight().GetType())
  l := self.integralPromotion(node.GetLeft().GetType())
  target := self.usualArithmeticConversion(l, r)
  if ! l.IsSameType(target) {
    typeNode := bs_ast.NewTypeNode(node.GetLocation(), bs_typesys.NewVoidTypeRef(node.GetLocation()))
    node.SetLeft(bs_ast.NewCastNode(node.GetLocation(), typeNode, node.GetLeft()))
  }
  if ! r.IsSameType(target) {
    typeNode := bs_ast.NewTypeNode(node.GetLocation(), bs_typesys.NewVoidTypeRef(node.GetLocation()))
    node.SetLeft(bs_ast.NewCastNode(node.GetLocation(), typeNode, node.GetRight()))
  }
  node.SetType(target)
}

func (self *TypeChecker) integralPromotion(t bs_core.IType) bs_core.IType {
  if ! t.IsInteger() {
    self.errorHandler.Errorf("integral promotion for %s", t)
  }
  intType := self.typeTable.SignedInt()
  if t.Size() < intType.Size() {
    return intType
  } else {
    return t
  }
}

func (self *TypeChecker) integralPromotedExpr(expr bs_core.IExprNode) bs_core.IExprNode {
  t := self.integralPromotion(expr.GetType())
  if t.IsSameType(expr.GetType()) {
    return expr
  } else {
    typeNode := bs_ast.NewTypeNode(expr.GetLocation(), bs_typesys.NewVoidTypeRef(expr.GetLocation()))
    return bs_ast.NewCastNode(expr.GetLocation(), typeNode, expr)
  }
}

func (self *TypeChecker) usualArithmeticConversion(l bs_core.IType, r bs_core.IType) bs_core.IType {
  s_int := self.typeTable.SignedInt()
  u_int := self.typeTable.UnsignedInt()
  s_long := self.typeTable.SignedLong()
  u_long := self.typeTable.UnsignedLong()
  if (l.IsSameType(u_int) && r.IsSameType(s_long)) || (r.IsSameType(u_int) && l.IsSameType(s_long)) {
    return u_long
  } else {
    if l.IsSameType(u_long) || r.IsSameType(u_long) {
      return u_long
    } else {
      if l.IsSameType(s_long) || r.IsSameType(s_long) {
        return s_long
      } else {
        if l.IsSameType(u_int) || r.IsSameType(u_int) {
          return u_int
        } else {
          return s_int
        }
      }
    }
  }
}

func (self *TypeChecker) expectsScalarLHS(node bs_core.IUnaryArithmeticOpNode) {
  if node.GetExpr().IsParameter() {
    // parameter is always a scalar.
  } else {
    if node.GetExpr().GetType().IsArray() {
      self.errorHandler.Errorf("%s wrong operand type for %s: %s", node.GetLocation(), node.GetOperator(), node.GetExpr())
      return
    } else {
      self.mustBeScalar(node.GetExpr(), node.GetOperator())
    }
  }
  if node.GetExpr().GetType().IsInteger() {
    opType := self.integralPromotion(node.GetExpr().GetType())
    if ! node.GetExpr().GetType().IsSameType(opType) {
      node.SetOpType(opType)
    }
    node.SetAmount(1)
  } else {
    if node.GetExpr().GetType().IsPointer() {
      if node.GetExpr().GetType().GetBaseType().IsVoid() {
        self.errorHandler.Errorf("%s wrong operand type for %s: %s", node.GetLocation(), node.GetOperator(), node.GetExpr())
        return
      }
      node.SetAmount(node.GetExpr().GetType().GetBaseType().Size())
    } else {
      panic("must not happen")
    }
  }
}

func (self *TypeChecker) checkLHS(lhs bs_core.IExprNode) bool {
  if lhs.IsParameter() {
    // parameter is always assignable.
    return true
  } else {
    if self.isInvalidLHSType(lhs.GetType()) {
      self.errorHandler.Errorf("%s invalid LHS expression type: %s", lhs.GetLocation(), lhs.GetType())
      return false
    } else {
      return true
    }
  }
}

func (self *TypeChecker) checkRHS(rhs bs_core.IExprNode) bool {
  if self.isInvalidRHSType(rhs.GetType()) {
    self.errorHandler.Errorf("%s invalid RHS expression type: %s", rhs.GetLocation(), rhs.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) expectsSameIntegerOrPointerDiff(node bs_core.IBinaryOpNode) {
  if node.GetLeft().IsPointer() && node.GetRight().IsPointer() {
    if node.GetOperator() == "+" {
      self.errorHandler.Errorf("%s invalid operation: pointer + pointer", node.GetLocation())
      return
    }
    node.SetType(self.typeTable.PtrDiffType())
  }
}

func (self *TypeChecker) expectsSameInteger(node bs_core.IBinaryOpNode) {
  if ! self.mustBeInteger(node.GetLeft(), node.GetOperator()) {
    return
  }
  if ! self.mustBeInteger(node.GetRight(), node.GetOperator()) {
    return
  }
  self.arithmeticImplicitCast(node)
}

func (self *TypeChecker) VisitStmtNode(unknown bs_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.BlockNode: {
      vars := node.GetVariables()
      for i := range vars {
        self.checkVariable(vars[i])
      }
      bs_ast.VisitStmtNodes(self, node.GetStmts())
    }
    case *bs_ast.ExprStmtNode: {
      bs_ast.VisitExprNode(self, node.GetExpr())
      if self.isInvalidStatementType(node.GetExpr().GetType()) {
        self.errorHandler.Errorf("%s invalid statement type: %s", node.GetLocation(), node.GetExpr().GetType())
      }
    }
    case *bs_ast.IfNode: {
      visitIfNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *bs_ast.WhileNode: {
      visitWhileNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *bs_ast.ForNode: {
      visitForNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *bs_ast.SwitchNode: {
      visitSwitchNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *bs_ast.ReturnNode: {
      visitReturnNode(self, node)
      if self.currentFunction.IsVoid() {
        if node.GetExpr() != nil {
          self.errorHandler.Errorf("%s returning value from void function", node.GetLocation())
        }
        if node.GetExpr().GetType().IsVoid() {
          self.errorHandler.Errorf("%s returning void", node.GetLocation())
        }
        node.SetExpr(self.implicitCast(self.currentFunction.GetReturnType(), node.GetExpr()))
      }
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *TypeChecker) VisitExprNode(unknown bs_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.AssignNode: {
      visitAssignNode(self, node)
      if self.checkLHS(node.GetLHS()) {
        if self.checkRHS(node.GetRHS()) {
          node.SetRHS(self.implicitCast(node.GetLHS().GetType(), node.GetRHS()))
        }
      }
    }
    case *bs_ast.OpAssignNode: {
      visitOpAssignNode(self, node)
      if self.checkLHS(node.GetLHS()) {
        if self.checkRHS(node.GetRHS()) {
          if node.GetLHS().GetType().IsPointer() {
            self.mustBeInteger(node.GetRHS(), node.GetOperator())
            node.SetRHS(self.integralPromotedExpr(node.GetRHS()))
          } else {
            if self.mustBeInteger(node.GetLHS(), node.GetOperator()) {
              if self.mustBeInteger(node.GetRHS(), node.GetOperator()) {
                l := self.integralPromotion(node.GetLHS().GetType())
                r := self.integralPromotion(node.GetRHS().GetType())
                opType := self.usualArithmeticConversion(l, r)
                if ! opType.IsCompatible(l) && self.isSafeIntegerCast(node.GetRHS(), opType) {
                  self.errorHandler.Warnf("%s incompatible implicit cast from %s to %s", node.GetLocation(), opType, l)
                }
                if ! r.IsSameType(opType) {
                  typeNode := bs_ast.NewTypeNode(node.GetLocation(), bs_typesys.NewVoidTypeRef(node.GetLocation()))
                  node.SetRHS(bs_ast.NewCastNode(node.GetLocation(), typeNode, node.GetRHS()))
                }
              }
            }
          }
        }
      }
    }
    case *bs_ast.CondExprNode: {
      visitCondExprNode(self, node)
      self.checkCond(node.GetCond())
      t := node.GetThenExpr().GetType()
      e := node.GetElseExpr().GetType()
      if ! t.IsSameType(e) {
        if t.IsCompatible(e) {
          // insert cast on thenBody
          typeNode := bs_ast.NewTypeNode(node.GetLocation(), bs_typesys.NewVoidTypeRef(node.GetLocation()))
          cast := bs_ast.NewCastNode(node.GetLocation(), typeNode, node.GetThenExpr())
          node.SetThenExpr(cast)
        } else {
          if e.IsCompatible(t) {
            // insert cast on elseBody
            typeNode := bs_ast.NewTypeNode(node.GetLocation(), bs_typesys.NewVoidTypeRef(node.GetLocation()))
            cast := bs_ast.NewCastNode(node.GetLocation(), typeNode, node.GetElseExpr())
            node.SetElseExpr(cast)
          } else {
            self.errorHandler.Errorf("%s invalid cast from %s to %s", node.GetLocation(), e, t)
          }
        }
      }
    }
    case *bs_ast.BinaryOpNode: {
      visitBinaryOpNode(self, node)
      if node.GetOperator() == "+" || node.GetOperator() == "-" {
        self.expectsSameIntegerOrPointerDiff(node)
      } else {
        switch node.GetOperator() {
          case "*", "/", "%", "&", "|", "^", "<<", ">>": {
            self.expectsSameInteger(node)
          }
          case "==", "!=", "<", "<=", ">", ">=": {
            self.expectsComparableScalars(node)
          }
          default: {
            self.errorHandler.Errorf("unknown binary operator: %s", node.GetOperator())
          }
        }
      }
    }
    case *bs_ast.LogicalAndNode: {
      visitLogicalAndNode(self, node)
      self.expectsComparableScalars(node)
    }
    case *bs_ast.LogicalOrNode: {
      visitLogicalOrNode(self, node)
      self.expectsComparableScalars(node)
    }
    case *bs_ast.UnaryOpNode: {
      visitUnaryOpNode(self, node)
      if node.GetOperator() == "!" {
        self.mustBeScalar(node.GetExpr(), node.GetOperator())
      } else {
        self.mustBeInteger(node.GetExpr(), node.GetOperator())
      }
    }
    case *bs_ast.PrefixOpNode: {
      visitPrefixOpNode(self, node)
      self.expectsScalarLHS(node)
    }
    case *bs_ast.SuffixOpNode: {
      visitSuffixOpNode(self, node)
      self.expectsScalarLHS(node)
    }
    case *bs_ast.FuncallNode: {
      visitFuncallNode(self, node)
      t := node.GetFunctionType()
      if ! t.AcceptsArgc(node.NumArgs()) {
        self.errorHandler.Errorf("%s wrong number of arguments: %d", node.GetLocation(), node.NumArgs())
      } else {
        args := node.GetArgs()
        paramDescs := t.GetParamTypes().GetParamDescs()
        if len(args) < len(paramDescs) {
          self.errorHandler.Errorf("%s missing argument: %d for %d", node.GetLocation(), len(args), len(paramDescs))
        }
        newArgs := []bs_core.IExprNode { }
        for i := range args {
          arg := args[i]
          if i < len(paramDescs) {
            // mandatory args
            if self.checkRHS(arg) {
              arg = self.implicitCast(paramDescs[i], arg)
            }
          } else {
            // optionary args
            if self.checkRHS(arg) {
              arg = self.castOptionalArg(arg)
            }
          }
          newArgs = append(newArgs, arg)
        }
        node.SetArgs(newArgs)
      }
    }
    case *bs_ast.ArefNode: {
      visitArefNode(self, node)
      self.mustBeInteger(node.GetIndex(), "[]")
    }
    case *bs_ast.CastNode: {
      visitCastNode(self, node)
      if ! node.GetExpr().GetType().IsCastableTo(node.GetType()) {
        self.errorHandler.Errorf("%s invalid cast from %s to %s", node.GetLocation(), node.GetExpr().GetType(), node.GetType())
      }
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *TypeChecker) VisitTypeDefinition(unknown bs_core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
