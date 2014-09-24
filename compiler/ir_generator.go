package compiler

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/ir"
  "bitbucket.org/yyuu/bs/typesys"
)

type IRGenerator struct {
  errorHandler *core.ErrorHandler
  typeTable *typesys.TypeTable
  exprNestLevel int
  stmts []core.IStmt
  scopeStack []*entity.LocalScope
//breakStack []*Label { }
//continueStack []*Label { }
//jumpMap make(map[string]*JumpEntry)
}

func NewIRGenerator(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *IRGenerator {
  stmts := []core.IStmt { }
  scopeStack := []*entity.LocalScope { }
  return &IRGenerator { errorHandler, table, 0, stmts, scopeStack }
}

func (self *IRGenerator) Generate(a *ast.AST) *ir.IR {
  vs := a.GetDefinedVariables()
  for i := range vs {
    if vs[i].HasInitializer() {
      vs[i].SetIR(self.transformExpr(vs[i].GetInitializer()))
    }
  }
  fs := a.GetDefinedFunctions()
  for i := range fs {
    fs[i].SetIR(self.compileFunctionBody(fs[i]))
  }
  return a.GenerateIR()
}

func (self *IRGenerator) compileFunctionBody(f *entity.DefinedFunction) []core.IStmt {
  self.stmts = []core.IStmt { }
  self.scopeStack = []*entity.LocalScope { }
//self.breakStack = []*Label { }
//self.continueStack = []*Label { }
//self.jumpMap = make(map[string]*JumpEntry)
//self.transformStmt(f.GetBody())
//self.checkJumpLinks(jumpMap)
  return self.stmts
}

func (self *IRGenerator) transformStmt(node core.IStmtNode) {
  ast.VisitStmt(self, node)
}

func (self *IRGenerator) transformStmtExpr(node core.IExprNode) {
  ast.VisitExpr(self, node)
}

func (self *IRGenerator) transformExpr(node core.IExprNode) core.IExpr {
  self.exprNestLevel++
  e := ast.VisitExpr(self, node)
  self.exprNestLevel--
  return e.(core.IExpr)
}

func (self *IRGenerator) isStatement() bool {
  return self.exprNestLevel == 0
}

func (self *IRGenerator) assign(loc core.Location, lhs core.IExpr, rhs core.IExpr) {
  self.stmts = append(self.stmts, ir.NewAssign(loc, self.addressOf(lhs), rhs))
}

func (self *IRGenerator) tmpVar(t core.IType) *entity.DefinedVariable {
  last := len(self.scopeStack)-1
  ref := self.typeTable.GetTypeRef(t)
  typeNode := ast.NewTypeNode(core.NewLocation("[builtin:ir_generator]", 0, 0), ref)
  return self.scopeStack[last].AllocateTmp(typeNode)
}

func (self *IRGenerator) transformIndex(node *ast.ArefNode) core.IExpr {
  if node.IsMultiDimension() {
    return ir.NewBin(self.int_t(), ir.OP_ADD,
                     self.transformExpr(node.GetIndex()),
                     ir.NewBin(self.int_t(), ir.OP_MUL,
                               ir.NewInt(self.int_t(), int64(node.GetLength())),
                               self.transformIndex(node.GetExpr().(*ast.ArefNode))))

  } else {
    return self.transformExpr(node.GetIndex())
  }
}

func (self *IRGenerator) transformOpAssign(loc core.Location, op int, lhsType core.IType, lhs core.IExpr, rhs core.IExpr) core.IExpr {
  if lhs.IsVar() {
    self.assign(loc, lhs, self.bin(op, lhsType, lhs, rhs))
    if self.isStatement() {
      return nil
    } else {
      return lhs
    }
  } else {
    a := self.tmpVar(self.pointerTo(lhsType))
    self.assign(loc, self.ref(a), self.addressOf(lhs))
    self.assign(loc, self.mem(a), self.bin(op, lhsType, self.mem(a), rhs))
    if self.isStatement() {
      return nil
    } else {
      return self.mem(a)
    }
  }
}

func (self *IRGenerator) bin(op int, leftType core.IType, left core.IExpr, right core.IExpr) *ir.Bin {
  if self.isPointerArithmetic(op, leftType) {
    return ir.NewBin(left.GetType(), op, left,
                     ir.NewBin(right.GetType(), ir.OP_MUL,
                               right,
                               self.ptrBaseSize(leftType)))
  } else {
    return ir.NewBin(left.GetType(), op, left, right)
  }
}

func (self *IRGenerator) isPointerDiff(op int, l core.IType, r core.IType) bool {
  return op == ir.OP_SUB && l.IsPointer() && r.IsPointer()
}

func (self *IRGenerator) isPointerArithmetic(op int, operandType core.IType) bool {
  switch op {
    case ir.OP_ADD: return operandType.IsPointer()
    case ir.OP_SUB: return operandType.IsPointer()
    default:        return false
  }
}

func (self *IRGenerator) ptrBaseSize(t core.IType) core.IExpr {
  return ir.NewInt(self.ptrdiff_t(), int64(t.GetBaseType().Size()))
}

func (self *IRGenerator) binOp(uniOp string) int {
  if uniOp == "++" {
    return ir.OP_ADD
  } else {
    if uniOp == "--" {
      return ir.OP_SUB
    } else {
      panic("must not happen")
    }
  }
}

func (self *IRGenerator) addressOf(expr core.IExpr) core.IExpr {
  return expr.GetAddressNode(self.ptr_t())
}

func (self *IRGenerator) ref(ent core.IEntity) *ir.Var {
  return ir.NewVar(self.varType(ent.GetType()), ent)
}

func (self *IRGenerator) mem(ent core.IEntity) *ir.Mem {
  return ir.NewMem(self.asmType(ent.GetType().GetBaseType()), self.ref(ent))
}

func (self *IRGenerator) mem2(expr core.IExpr, t core.IType) *ir.Mem {
  return ir.NewMem(self.asmType(t), expr)
}

func (self *IRGenerator) ptrdiff(n int64) *ir.Int {
  return ir.NewInt(self.ptrdiff_t(), n)
}

func (self *IRGenerator) size(n int64) *ir.Int {
  return ir.NewInt(self.size_t(), n)
}

func (self *IRGenerator) imm(operandType core.IType, n int64) *ir.Int {
  if operandType.IsPointer() {
    return ir.NewInt(self.ptrdiff_t(), n)
  } else {
    return ir.NewInt(self.int_t(), n)
  }
}

func (self *IRGenerator) pointerTo(t core.IType) core.IType {
  return self.typeTable.PointerTo(t)
}

func (self *IRGenerator) asmType(t core.IType) core.IType {
  return t // FIXME: asm type
}

func (self *IRGenerator) varType(t core.IType) core.IType {
  return t // FIXME: asm type
}

func (self *IRGenerator) int_t() core.IType {
  return self.typeTable.UnsignedInt() // FIXME: asm type
}

func (self *IRGenerator) size_t() core.IType {
  return self.typeTable.UnsignedLong() // FIXME: asm type
}

func (self *IRGenerator) ptr_t() core.IType {
  return self.typeTable.UnsignedLong() // FIXME: asm type
}

func (self *IRGenerator) ptrdiff_t() core.IType {
  return self.typeTable.PtrDiffType() // FIXME: asm type
}

func (self *IRGenerator) VisitNode(unknown core.INode) interface{} {
  switch node := unknown.(type) {
//  case *ast.BlockNode: {
//  }
//  case *ast.ExprStmtNode: {
//    e := ast.VisitExpr(self, node.GetExpr())
//    if e != nil {
//      self.errorHandler.Warnf("%s useless expression\n", node.GetLocation())
//    }
//    return nil
//  }
//  case *ast.IfNode: {
//  }
//  case *ast.SwitchNode: {
//  }
    case *ast.CaseNode: {
      panic("must not happen")
    }
//  case *ast.WhileNode: {
//  }
//  case *ast.DoWhileNode: {
//  }
//  case *ast.ForNode: {
//  }
//  case *ast.BreakNode: {
//  }
//  case *ast.ContinueNode: {
//  }
//  case *ast.LabelNode: {
//  }
//  case *ast.GotoNode: {
//  }
//  case *ast.ReturnNode: {
//  }
//  case *ast.CondExprNode: {
//  }
//  case *ast.LogicalAndNode: {
//  }
//  case *ast.LogicalOrNode: {
//  }
    case *ast.AssignNode: {
      lloc := node.GetLHS().GetLocation()
      rloc := node.GetRHS().GetLocation()
      if self.isStatement() {
        rhs := self.transformExpr(node.GetRHS())
        self.assign(lloc, self.transformExpr(node.GetLHS()), rhs)
        return nil
      } else {
        tmp := self.tmpVar(node.GetRHS().GetType())
        self.assign(rloc, self.ref(tmp), self.transformExpr(node.GetRHS()))
        self.assign(lloc, self.transformExpr(node.GetLHS()), self.ref(tmp))
        return self.ref(tmp)
      }
    }
    case *ast.OpAssignNode: {
      rhs := self.transformExpr(node.GetRHS())
      lhs := self.transformExpr(node.GetLHS())
      t := node.GetLHS().GetType()
      op := ir.OpInternBinary(node.GetOperator(), t.IsSigned())
      return self.transformOpAssign(node.GetLocation(), op, t, lhs, rhs)
    }
    case *ast.PrefixOpNode: {
      t := node.GetExpr().GetType()
      return self.transformOpAssign(node.GetLocation(),
                                    self.binOp(node.GetOperator()),
                                    t,
                                    self.transformExpr(node.GetExpr()), self.imm(t, 1))
    }
    case *ast.SuffixOpNode: {
      expr := self.transformExpr(node.GetExpr())
      t := node.GetExpr().GetType()
      op := self.binOp(node.GetOperator())
      loc := node.GetLocation()
      if self.isStatement() {
        self.transformOpAssign(loc, op, t, expr, self.imm(t, 1))
        return nil
      } else {
        if expr.IsVar() {
          v := self.tmpVar(t)
          self.assign(loc, self.ref(v), expr)
          self.assign(loc, expr, self.bin(op, t, self.ref(v), self.imm(t, 1)))
          return self.ref(v)
        } else {
          a := self.tmpVar(self.pointerTo(t))
          v := self.tmpVar(t)
          self.assign(loc, self.ref(a), self.addressOf(expr))
          self.assign(loc, self.ref(v), self.mem(a))
          self.assign(loc, self.mem(a), self.bin(op, t, self.mem(a), self.imm(t, 1)))
          return self.ref(v)
        }
      }
    }
    case *ast.FuncallNode: {
      numArgs := node.NumArgs()
      args := make([]core.IExpr, numArgs)
      for i := range args {
        a := node.GetArg(numArgs-1-i)
        args[i] = self.transformExpr(a)
      }
      call := ir.NewCall(self.asmType(node.GetType()),
                         self.transformExpr(node.GetExpr()),
                         args)
      if self.isStatement() {
        self.stmts = append(self.stmts, ir.NewExprStmt(node.GetLocation(), call))
      } else {
        tmp := self.tmpVar(node.GetType())
        self.assign(node.GetLocation(), self.ref(tmp), call)
        return self.ref(tmp)
      }
    }
    case *ast.BinaryOpNode: {
      right := self.transformExpr(node.GetRight())
      left := self.transformExpr(node.GetLeft())
      op := ir.OpInternBinary(node.GetOperator(), node.GetType().IsSigned())
      t := node.GetType()
      r := node.GetRight().GetType()
      l := node.GetLeft().GetType()
      if self.isPointerDiff(op, l, r) {
        tmp := ir.NewBin(self.asmType(t), op, left, right)
        return ir.NewBin(self.asmType(t), ir.OP_S_DIV, tmp, self.ptrBaseSize(l))
      } else {
        if self.isPointerArithmetic(op, l) {
          return ir.NewBin(self.asmType(t), op,
                           left, 
                           ir.NewBin(self.asmType(r), ir.OP_MUL,
                                     right,
                                     self.ptrBaseSize(l)))
        } else {
          if self.isPointerArithmetic(op, r) {
            return ir.NewBin(self.asmType(t), op,
                             ir.NewBin(self.asmType(l), ir.OP_MUL, left, self.ptrBaseSize(r)),
                             right)
          } else {
            return ir.NewBin(self.asmType(t), op, left, right)
          }
        }
      }
    }
    case *ast.UnaryOpNode: {
      if node.GetOperator() == "+" {
        return self.transformExpr(node.GetExpr())
      } else {
        return ir.NewUni(self.asmType(node.GetType()),
                         ir.OpInternUnary(node.GetOperator()),
                         self.transformExpr(node.GetExpr()))
      }
    }
    case *ast.ArefNode: {
      expr := self.transformExpr(node.GetBaseExpr())
      offset := ir.NewBin(self.ptrdiff_t(), ir.OP_MUL, self.size(int64(node.GetElementSize())), self.transformIndex(node))
      addr := ir.NewBin(self.ptr_t(), ir.OP_ADD, expr, offset)
      return self.mem2(addr, node.GetType())
    }
    case *ast.MemberNode: {
      expr := self.addressOf(self.transformExpr(node.GetExpr()))
      offset := self.ptrdiff(int64(node.GetOffset()))
      addr := ir.NewBin(self.ptr_t(), ir.OP_ADD, expr, offset)
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *ast.PtrMemberNode: {
      expr := self.transformExpr(node.GetExpr())
      offset := self.ptrdiff(int64(node.GetOffset()))
      addr := ir.NewBin(self.ptr_t(), ir.OP_ADD, expr, offset)
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *ast.DereferenceNode: {
      addr := self.transformExpr(node.GetExpr())
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *ast.AddressNode: {
      e := self.transformExpr(node.GetExpr())
      if node.GetExpr().IsLoadable() {
        return self.addressOf(e)
      } else {
        return e
      }
    }
    case *ast.CastNode: {
      if node.IsEffectiveCast() {
        if node.GetExpr().GetType().IsSigned() {
          return ir.NewUni(self.asmType(node.GetType()), ir.OP_S_CAST, self.transformExpr(node.GetExpr()))
        } else {
          return ir.NewUni(self.asmType(node.GetType()), ir.OP_U_CAST, self.transformExpr(node.GetExpr()))
        }
      } else {
        if self.isStatement() {
          self.transformStmtExpr(node.GetExpr())
        } else {
          return self.transformExpr(node.GetExpr())
        }
      }
    }
    case *ast.SizeofExprNode: {
      return ir.NewInt(self.size_t(), int64(node.GetExpr().GetType().AllocSize()))
    }
    case *ast.SizeofTypeNode: {
      return ir.NewInt(self.size_t(), int64(node.GetOperandType().AllocSize()))
    }
    case *ast.VariableNode: {
      if node.GetEntity().IsConstant() {
        return self.transformExpr(node.GetEntity().GetValue())
      } else {
        v := self.ref(node.GetEntity())
        if node.IsLoadable() {
          return v
        } else {
          return self.addressOf(v)
        }
      }
    }
    case *ast.IntegerLiteralNode: {
      return ir.NewInt(self.asmType(node.GetType()), node.GetValue())
    }
    case *ast.StringLiteralNode: {
      return ir.NewStr(self.asmType(node.GetType()), node.GetEntry())
    }
    default: {
      visitNode(self, unknown)
    }
  }
  return nil
}
