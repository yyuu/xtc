package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

type TypeChecker struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
  currentFunction *entity.DefinedFunction
}

func NewTypeChecker(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *TypeChecker {
  return &TypeChecker { errorHandler, table, nil }
}

func (self *TypeChecker) Check(a *ast.AST) {
  self.errorHandler.Debug("starting type checker.")
  vs := a.GetDefinedVariables()
  for i := range vs {
    self.checkVariable(vs[i])
  }
  fs := a.GetDefinedFunctions()
  for i := range fs {
    self.currentFunction = fs[i]
    self.checkReturnType(fs[i])
    self.checkParamTypes(fs[i])
    ast.VisitStmt(self, fs[i].GetBody())
  }
  self.errorHandler.Debug("finished type checker.")
}

func (self *TypeChecker) checkVariable(v *entity.DefinedVariable) {
  if self.isInvalidVariableType(v.GetType()) {
    self.errorHandler.Fatalf("invalid variable type")
  }
  if v.HasInitializer() {
    if self.isInvalidLHSType(v.GetType()) {
      self.errorHandler.Fatalf("invalid LHS type: %s", v.GetType())
    }
    ast.VisitExpr(self, v.GetInitializer())
    v.SetInitializer(self.implicitCast(v.GetType(), v.GetInitializer()))
  }
}

func (self *TypeChecker) isInvalidVariableType(t core.IType) bool {
  return t.IsVoid() || (t.IsArray() && ! t.IsAllocatedArray())
}

func (self *TypeChecker) isInvalidLHSType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid() || t.IsArray()
}

func (self *TypeChecker) isInvalidRHSType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid()
}

func (self *TypeChecker) implicitCast(t core.IType, expr core.IExprNode) core.IExprNode {
  if expr.GetType().IsSameType(t) {
    return expr
  } else {
    if expr.GetType().IsCastableTo(t) {
      if ! expr.GetType().IsCompatible(t) && ! self.isSafeIntegerCast(expr, t) {
        self.errorHandler.Warnf("%s incompatible inplicit cast from %s to %s", expr.GetLocation(), expr.GetType(), t)
      }
      typeNode := ast.NewTypeNode(expr.GetLocation(), typesys.NewVoidTypeRef(expr.GetLocation()))
      typeNode.SetType(t)
      return ast.NewCastNode(expr.GetLocation(), typeNode, expr)
    } else {
      self.errorHandler.Fatalf("invalid cast error: %s to %s", expr.GetType(), t)
      return expr
    }
  }
}

func (self *TypeChecker) castOptionalArg(arg core.IExprNode) core.IExprNode {
  if ! arg.GetType().IsInteger() {
    return arg
  } else {
    var t core.IType
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

func (self *TypeChecker) isSafeIntegerCast(node core.INode, t core.IType) bool {
  if ! t.IsInteger() {
    return false
  } else {
    i, ok := t.(typesys.IntegerType)
    if ! ok {
      return false
    }
    n, ok := node.(ast.IntegerLiteralNode)
    if ! ok {
      return false
    }
    return i.IsInDomain(n.GetValue())
  }
}

func (self *TypeChecker) checkReturnType(f *entity.DefinedFunction) {
  if self.isInvalidReturnType(f.GetReturnType()) {
    self.errorHandler.Fatalf("returns invalid type: %s", f.GetReturnType())
  }
}

func (self *TypeChecker) isInvalidReturnType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsArray()
}

func (self *TypeChecker) checkParamTypes(f *entity.DefinedFunction) {
  params := f.GetParameters()
  for i := range params {
    param := params[i]
    if self.isInvalidParameterType(param.GetType()) {
      self.errorHandler.Fatalf("invalid parameter type: %s", param.GetType())
    }
  }
}

func (self *TypeChecker) isInvalidParameterType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion() || t.IsVoid() || t.IsIncompleteArray()
}

func (self *TypeChecker) isInvalidStatementType(t core.IType) bool {
  return t.IsStruct() || t.IsUnion()
}

func (self *TypeChecker) mustBeInteger(expr core.IExprNode, op string) bool {
  if ! expr.GetType().IsInteger() {
    self.errorHandler.Fatalf("%s wrong operand type for %s: %s", expr.GetLocation(), op, expr.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) mustBeScalar(expr core.IExprNode, op string) bool {
  if ! expr.GetType().IsScalar() {
    self.errorHandler.Fatalf("%s wrong operand type for %s: %s", expr.GetLocation(), op, expr.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) checkCond(cond core.IExprNode) {
  self.mustBeScalar(cond, "condition expression")
}

func (self *TypeChecker) expectsComparableScalars(node core.IBinaryOpNode) {
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

func (self *TypeChecker) forcePointerType(master core.IExprNode, slave core.IExprNode) core.IExprNode {
  if master.GetType().IsCompatible(slave.GetType()) {
    return slave
  } else {
    self.errorHandler.Warnf("incompatible implicit cast from %s to %s", slave.GetType(), master.GetType())
    typeNode := ast.NewTypeNode(master.GetLocation(), typesys.NewVoidTypeRef(master.GetLocation()))
    typeNode.SetType(master.GetType())
    return ast.NewCastNode(master.GetLocation(), typeNode, slave)
  }
}

func (self *TypeChecker) arithmeticImplicitCast(node core.IBinaryOpNode) {
  r := self.integralPromotion(node.GetRight().GetType())
  l := self.integralPromotion(node.GetLeft().GetType())
  target := self.usualArithmeticConversion(l, r)
  if ! l.IsSameType(target) {
    typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
    node.SetLeft(ast.NewCastNode(node.GetLocation(), typeNode, node.GetLeft()))
  }
  if ! r.IsSameType(target) {
    typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
    node.SetLeft(ast.NewCastNode(node.GetLocation(), typeNode, node.GetRight()))
  }
  node.SetType(target)
}

func (self *TypeChecker) integralPromotion(t core.IType) core.IType {
  if ! t.IsInteger() {
    self.errorHandler.Fatalf("integral promotion for %s", t)
  }
  intType := self.typeTable.SignedInt()
  if t.Size() < intType.Size() {
    return intType
  } else {
    return t
  }
}

func (self *TypeChecker) integralPromotedExpr(expr core.IExprNode) core.IExprNode {
  t := self.integralPromotion(expr.GetType())
  if t.IsSameType(expr.GetType()) {
    return expr
  } else {
    typeNode := ast.NewTypeNode(expr.GetLocation(), typesys.NewVoidTypeRef(expr.GetLocation()))
    return ast.NewCastNode(expr.GetLocation(), typeNode, expr)
  }
}

func (self *TypeChecker) usualArithmeticConversion(l core.IType, r core.IType) core.IType {
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

func (self *TypeChecker) expectsScalarLHS(node core.IUnaryArithmeticOpNode) {
  if node.GetExpr().IsParameter() {
    // parameter is always a scalar.
  } else {
    if node.GetExpr().GetType().IsArray() {
      self.errorHandler.Fatalf("%s wrong operand type for %s: %s", node.GetLocation(), node.GetOperator(), node.GetExpr())
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
        self.errorHandler.Fatalf("%s wrong operand type for %s: %s", node.GetLocation(), node.GetOperator(), node.GetExpr())
        return
      }
      node.SetAmount(node.GetExpr().GetType().GetBaseType().Size())
    } else {
      panic("must not happen")
    }
  }
}

func (self *TypeChecker) checkLHS(lhs core.IExprNode) bool {
  if lhs.IsParameter() {
    // parameter is always assignable.
    return true
  } else {
    if self.isInvalidLHSType(lhs.GetType()) {
      self.errorHandler.Fatalf("%s invalid LHS expression type: %s", lhs.GetLocation(), lhs.GetType())
      return false
    } else {
      return true
    }
  }
}

func (self *TypeChecker) checkRHS(rhs core.IExprNode) bool {
  if self.isInvalidRHSType(rhs.GetType()) {
    self.errorHandler.Fatalf("%s invalid RHS expression type: %s", rhs.GetLocation(), rhs.GetType())
    return false
  } else {
    return true
  }
}

func (self *TypeChecker) expectsSameIntegerOrPointerDiff(node core.IBinaryOpNode) {
  if node.GetLeft().IsPointer() && node.GetRight().IsPointer() {
    if node.GetOperator() == "+" {
      self.errorHandler.Fatalf("%s invalid operation: pointer + pointer", node.GetLocation())
      return
    }
    node.SetType(self.typeTable.PtrDiffType())
  }
}

func (self *TypeChecker) expectsSameInteger(node core.IBinaryOpNode) {
  if ! self.mustBeInteger(node.GetLeft(), node.GetOperator()) {
    return
  }
  if ! self.mustBeInteger(node.GetRight(), node.GetOperator()) {
    return
  }
  self.arithmeticImplicitCast(node)
}

func (self *TypeChecker) VisitNode(unknown core.INode) interface{} {
  switch node := unknown.(type) {
    case *ast.BlockNode: {
      vars := node.GetVariables()
      for i := range vars {
        self.checkVariable(vars[i])
      }
      ast.VisitStmts(self, node.GetStmts())
    }
    case *ast.ExprStmtNode: {
      ast.VisitExpr(self, node.GetExpr())
      if self.isInvalidStatementType(node.GetExpr().GetType()) {
        self.errorHandler.Fatalf("%s invalid statement type: %s", node.GetLocation(), node.GetExpr().GetType())
      }
    }
    case *ast.IfNode: {
      visitIfNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.WhileNode: {
      visitWhileNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.ForNode: {
      visitForNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.SwitchNode: {
      visitSwitchNode(self, node)
      self.checkCond(node.GetCond())
    }
    case *ast.ReturnNode: {
      visitReturnNode(self, node)
      if self.currentFunction.IsVoid() {
        if node.GetExpr() != nil {
          self.errorHandler.Fatalf("%s returning value from void function", node.GetLocation())
        }
        if node.GetExpr().GetType().IsVoid() {
          self.errorHandler.Fatalf("%s returning void", node.GetLocation())
        }
        node.SetExpr(self.implicitCast(self.currentFunction.GetReturnType(), node.GetExpr()))
      }
    }
    case *ast.AssignNode: {
      visitAssignNode(self, node)
      if self.checkLHS(node.GetLHS()) {
        if self.checkRHS(node.GetRHS()) {
          node.SetRHS(self.implicitCast(node.GetLHS().GetType(), node.GetRHS()))
        }
      }
    }
    case *ast.OpAssignNode: {
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
                  typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
                  node.SetRHS(ast.NewCastNode(node.GetLocation(), typeNode, node.GetRHS()))
                }
              }
            }
          }
        }
      }
    }
    case *ast.CondExprNode: {
      visitCondExprNode(self, node)
      self.checkCond(node.GetCond())
      t := node.GetThenExpr().GetType()
      e := node.GetElseExpr().GetType()
      if ! t.IsSameType(e) {
        if t.IsCompatible(e) {
          // insert cast on thenBody
          typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
          cast := ast.NewCastNode(node.GetLocation(), typeNode, node.GetThenExpr())
          node.SetThenExpr(cast)
        } else {
          if e.IsCompatible(t) {
            // insert cast on elseBody
            typeNode := ast.NewTypeNode(node.GetLocation(), typesys.NewVoidTypeRef(node.GetLocation()))
            cast := ast.NewCastNode(node.GetLocation(), typeNode, node.GetElseExpr())
            node.SetElseExpr(cast)
          } else {
            self.errorHandler.Fatalf("%s invalid cast from %s to %s", node.GetLocation(), e, t)
          }
        }
      }
    }
    case *ast.BinaryOpNode: {
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
            panic(fmt.Errorf("unknown binary operator: %s", node.GetOperator()))
          }
        }
      }
    }
    case *ast.LogicalAndNode: {
      visitLogicalAndNode(self, node)
      self.expectsComparableScalars(node)
    }
    case *ast.LogicalOrNode: {
      visitLogicalOrNode(self, node)
      self.expectsComparableScalars(node)
    }
    case *ast.UnaryOpNode: {
      visitUnaryOpNode(self, node)
      if node.GetOperator() == "!" {
        self.mustBeScalar(node.GetExpr(), node.GetOperator())
      } else {
        self.mustBeInteger(node.GetExpr(), node.GetOperator())
      }
    }
    case *ast.PrefixOpNode: {
      visitPrefixOpNode(self, node)
      self.expectsScalarLHS(node)
    }
    case *ast.SuffixOpNode: {
      visitSuffixOpNode(self, node)
      self.expectsScalarLHS(node)
    }
    case *ast.FuncallNode: {
      visitFuncallNode(self, node)
      t := node.GetFunctionType()
      if ! t.AcceptsArgc(node.NumArgs()) {
        self.errorHandler.Fatalf("%s wrong number of arguments: %d", node.GetLocation(), node.NumArgs())
      } else {
        args := node.GetArgs()
        paramDescs := t.GetParamTypes().GetParamDescs()
        if len(args) < len(paramDescs) {
          panic(fmt.Errorf("%s missing argument: %d for %d", node.GetLocation(), len(args), len(paramDescs)))
        }
        newArgs := []core.IExprNode { }
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
    case *ast.ArefNode: {
      visitArefNode(self, node)
      self.mustBeInteger(node.GetIndex(), "[]")
    }
    case *ast.CastNode: {
      visitCastNode(self, node)
      if ! node.GetExpr().GetType().IsCastableTo(node.GetType()) {
        self.errorHandler.Fatalf("%s invalid cast from %s to %s", node.GetLocation(), node.GetExpr().GetType(), node.GetType())
      }
    }
    default: {
      visitNode(self, unknown)
    }
  }
  return nil
}
