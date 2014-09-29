package compiler

import (
  "fmt"
  "bitbucket.org/yyuu/bs/asm"
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
  breakStack []*asm.Label
  continueStack []*asm.Label
  jumpMap map[string]*jumpEntry
}

type jumpEntry struct {
  label *asm.Label
  numRefered int
  isDefined bool
  location core.Location
}

func newJumpEntry(label *asm.Label) *jumpEntry {
  loc := core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  return &jumpEntry { label, 0, false, loc }
}

func NewIRGenerator(errorHandler *core.ErrorHandler, table *typesys.TypeTable) *IRGenerator {
  stmts := []core.IStmt { }
  scopeStack := []*entity.LocalScope { }
  breakStack := []*asm.Label { }
  continueStack := []*asm.Label { }
  jumpMap := make(map[string]*jumpEntry)
  return &IRGenerator { errorHandler, table, 0, stmts, scopeStack, breakStack, continueStack, jumpMap }
}

func (self *IRGenerator) Generate(a *ast.AST) *ir.IR {
  self.errorHandler.Debug("starting IR generator.")
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
  x := a.GenerateIR()
  self.errorHandler.Debug("finished IR generator.")
  return x
}

func (self *IRGenerator) compileFunctionBody(f *entity.DefinedFunction) []core.IStmt {
  self.stmts = []core.IStmt { }
  self.scopeStack = []*entity.LocalScope { }
  self.breakStack = []*asm.Label { }
  self.continueStack = []*asm.Label { }
  self.jumpMap = make(map[string]*jumpEntry)
  self.transformStmt(f.GetBody())
  self.checkJumpLinks(self.jumpMap)
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
  ref := self.typeTable.GetTypeRef(t)
  typeNode := ast.NewTypeNode(core.NewLocation("[builtin:ir_generator]", 0, 0), ref)
  typeNode.SetType(t)
  return self.currentScope().AllocateTmp(typeNode)
}

func (self *IRGenerator) label(label *asm.Label) {
  loc := core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  self.stmts = append(self.stmts, ir.NewLabelStmt(loc, label))
}

func (self *IRGenerator) jump(target *asm.Label) {
  loc := core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  self.stmts = append(self.stmts, ir.NewJump(loc, target))
}

func (self *IRGenerator) cjump(loc core.Location, cond core.IExpr, thenLabel *asm.Label, elseLabel *asm.Label) {
  self.stmts = append(self.stmts, ir.NewCJump(loc, cond, thenLabel, elseLabel))
}

func (self *IRGenerator) pushBreak(label *asm.Label) {
  self.breakStack = append(self.breakStack, label)
}

func (self *IRGenerator) popBreak() *asm.Label {
  if len(self.breakStack) < 1 {
    panic("unmatched push/pop for break stack")
  }
  label := self.currentBreakTarget()
  self.breakStack = self.breakStack[0:len(self.breakStack)-1]
  return label
}

func (self *IRGenerator) currentBreakTarget() *asm.Label {
  if len(self.breakStack) < 1 {
    panic("break from out of loop")
  }
  return self.breakStack[len(self.breakStack)-1]
}

func (self *IRGenerator) pushContinue(label *asm.Label) {
  self.continueStack = append(self.continueStack, label)
}

func (self *IRGenerator) popContinue() *asm.Label {
  if len(self.continueStack) < 1 {
    panic("unmatched push/pop for continue stack")
  }
  label := self.currentContinueTarget()
  self.continueStack = self.continueStack[0:len(self.continueStack)-1]
  return label
}

func (self *IRGenerator) currentContinueTarget() *asm.Label {
  if len(self.continueStack) < 1 {
    panic("continue from out of loop")
  }
  return self.continueStack[len(self.continueStack)-1]
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
    return ir.NewBin(left.GetTypeId(), op, left,
                     ir.NewBin(right.GetTypeId(), ir.OP_MUL,
                               right,
                               self.ptrBaseSize(leftType)))
  } else {
    return ir.NewBin(left.GetTypeId(), op, left, right)
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

func (self *IRGenerator) asmType(t core.IType) int {
  if t.IsVoid() {
    return self.int_t()
  } else {
    return asm.TypeGet(t.Size())
  }
}

func (self *IRGenerator) varType(t core.IType) int {
  if ! t.IsScalar() {
    return 0
  } else {
    return asm.TypeGet(t.Size())
  }
}

func (self *IRGenerator) int_t() int {
  return asm.TypeGet(self.typeTable.GetIntSize())
}

func (self *IRGenerator) size_t() int {
  return asm.TypeGet(self.typeTable.GetLongSize())
}

func (self *IRGenerator) ptr_t() int {
  return asm.TypeGet(self.typeTable.GetPointerSize())
}

func (self *IRGenerator) ptrdiff_t() int {
  return asm.TypeGet(self.typeTable.GetLongSize())
}

func (self *IRGenerator) currentScope() *entity.LocalScope {
  return self.scopeStack[len(self.scopeStack)-1]
}

func (self *IRGenerator) pushScope(scope *entity.LocalScope) {
  self.scopeStack = append(self.scopeStack, scope)
}

func (self *IRGenerator) popScope() *entity.LocalScope {
  scope := self.currentScope()
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

func (self *IRGenerator) defineLabel(name string, loc core.Location) *asm.Label {
  ent := self.getJumpEntry(name)
  if ent.isDefined {
    panic(fmt.Errorf("duplicated jump label in %s(): %s", name, name))
  }
  ent.isDefined = true
  ent.location = loc
  return ent.label
}

func (self *IRGenerator) referLabel(name string) *asm.Label {
  ent := self.getJumpEntry(name)
  ent.numRefered++
  return ent.label
}

func (self *IRGenerator) getJumpEntry(name string) *jumpEntry {
  ent := self.jumpMap[name]
  if ent == nil {
    ent = newJumpEntry(asm.NewUnnamedLabel())
    self.jumpMap[name] = ent
  }
  return ent
}

func (self *IRGenerator) checkJumpLinks(jumpMap map[string]*jumpEntry) {
  for name, jump := range jumpMap {
    if jump.isDefined {
      self.errorHandler.Fatalf("%s undefined label: %s", jump.location, name)
    }
    if jump.numRefered == 0 {
      self.errorHandler.Fatalf("%s useless label: %s", jump.location, name)
    }
  }
}

func (self *IRGenerator) VisitNode(unknown core.INode) interface{} {
  switch node := unknown.(type) {
    case *ast.BlockNode: {
      self.pushScope(node.GetScope())
      vs := node.GetVariables()
      for i := range vs {
        if vs[i].HasInitializer() {
          if vs[i].IsPrivate() {
            vs[i].SetIR(self.transformExpr(vs[i].GetInitializer()))
          } else {
            self.assign(node.GetLocation(), self.ref(vs[i]), self.transformExpr(vs[i].GetInitializer()))
          }
        }
      }
      stmts := node.GetStmts()
      for i := range stmts {
        self.transformStmt(stmts[i])
      }
      self.popScope()
    }
    case *ast.ExprStmtNode: {
      e := ast.VisitExpr(self, node.GetExpr())
      if e != nil {
        self.errorHandler.Warnf("%s useless expression", node.GetLocation())
      }
      return nil
    }
    case *ast.IfNode: {
      thenLabel := asm.NewUnnamedLabel()
      elseLabel := asm.NewUnnamedLabel()
      endLabel := asm.NewUnnamedLabel()
      cond := self.transformExpr(node.GetCond())
      if node.GetElseBody() == nil {
        self.cjump(node.GetLocation(), cond, thenLabel, endLabel)
        self.label(thenLabel)
        self.transformStmt(node.GetThenBody())
        self.label(endLabel)
      } else {
        self.cjump(node.GetLocation(), cond, thenLabel, endLabel)
        self.label(thenLabel)
        self.transformStmt(node.GetThenBody())
        self.jump(endLabel)
        self.label(elseLabel)
        self.transformStmt(node.GetElseBody())
        self.label(endLabel)
      }
      return nil
    }
    case *ast.SwitchNode: {
      caseNodes := node.GetCases()
      cases := make([]*ir.Case, len(caseNodes))
      endLabel := asm.NewUnnamedLabel()
      defaultLabel := endLabel
      cond := self.transformExpr(node.GetCond())
      for i := range caseNodes {
        c := caseNodes[i].(*ast.CaseNode)
        if c.IsDefault() {
          defaultLabel = c.GetLabel()
        } else {
          values := c.GetValues()
          for j := range values {
            v := self.transformExpr(values[j])
            cases[i] = ir.NewCase(v.(*ir.Int).GetValue(), c.GetLabel())
          }
        }
      }
      self.stmts = append(self.stmts, ir.NewSwitch(node.GetLocation(), cond, cases, defaultLabel, endLabel))
      self.pushBreak(endLabel)
      for i := range caseNodes {
        c := caseNodes[i].(*ast.CaseNode)
        self.label(c.GetLabel())
        self.transformStmt(c.GetBody())
      }
      self.popBreak()
      self.label(endLabel)
      return nil
    }
    case *ast.CaseNode: {
      panic("must not happen")
    }
    case *ast.WhileNode: {
      begLabel := asm.NewUnnamedLabel()
      bodyLabel := asm.NewUnnamedLabel()
      endLabel := asm.NewUnnamedLabel()
      self.label(begLabel)
      self.cjump(node.GetLocation(), self.transformExpr(node.GetCond()), bodyLabel, endLabel)
      self.label(bodyLabel)
      self.pushContinue(begLabel)
      self.pushBreak(endLabel)
      self.transformStmt(node.GetBody())
      self.popBreak()
      self.popContinue()
      self.jump(begLabel)
      self.label(endLabel)
      return nil
    }
    case *ast.DoWhileNode: {
      begLabel := asm.NewUnnamedLabel()
      contLabel := asm.NewUnnamedLabel()
      endLabel := asm.NewUnnamedLabel()
      self.pushContinue(contLabel)
      self.pushBreak(endLabel)
      self.label(begLabel)
      self.transformStmt(node.GetBody())
      self.popBreak()
      self.popContinue()
      self.label(contLabel)
      self.cjump(node.GetLocation(), self.transformExpr(node.GetCond()), begLabel, endLabel)
      self.label(endLabel)
      return nil
    }
    case *ast.ForNode: {
      begLabel := asm.NewUnnamedLabel()
      bodyLabel := asm.NewUnnamedLabel()
      contLabel := asm.NewUnnamedLabel()
      endLabel := asm.NewUnnamedLabel()
      self.transformStmtExpr(node.GetInit())
      self.label(begLabel)
      self.cjump(node.GetLocation(), self.transformExpr(node.GetCond()), bodyLabel, endLabel)
      self.label(bodyLabel)
      self.pushContinue(contLabel)
      self.pushBreak(endLabel)
      self.transformStmt(node.GetBody())
      self.popBreak()
      self.popContinue()
      self.label(contLabel)
      self.transformStmtExpr(node.GetIncr())
      self.jump(begLabel)
      self.label(endLabel)
      return nil
    }
    case *ast.BreakNode: {
      self.jump(self.currentBreakTarget())
      return nil
    }
    case *ast.ContinueNode: {
      self.jump(self.currentContinueTarget())
      return nil
    }
    case *ast.LabelNode: {
      stmt := ir.NewLabelStmt(node.GetLocation(), self.defineLabel(node.GetName(), node.GetLocation()))
      self.stmts = append(self.stmts, stmt)
      if node.GetStmt() != nil {
        self.transformStmt(node.GetStmt())
      }
      return nil
    }
    case *ast.GotoNode: {
      self.jump(self.referLabel(node.GetTarget()))
    }
    case *ast.ReturnNode: {
      var expr core.IExpr
      if node.GetExpr() != nil {
        expr = self.transformExpr(node.GetExpr())
      }
      self.stmts = append(self.stmts, ir.NewReturn(node.GetLocation(), expr))
      return nil
    }
    case *ast.CondExprNode: {
      thenLabel := asm.NewUnnamedLabel()
      elseLabel := asm.NewUnnamedLabel()
      endLabel := asm.NewUnnamedLabel()
      v := self.tmpVar(node.GetType())
      cond := self.transformExpr(node.GetCond())
      self.cjump(node.GetLocation(), cond, thenLabel, elseLabel)
      self.label(thenLabel)
      self.assign(node.GetThenExpr().GetLocation(), self.ref(v), self.transformExpr(node.GetThenExpr()))
      self.jump(endLabel)
      self.label(elseLabel)
      self.assign(node.GetElseExpr().GetLocation(), self.ref(v), self.transformExpr(node.GetElseExpr()))
      self.jump(endLabel)
      self.label(endLabel)
      if self.isStatement() {
        return nil
      } else {
        return self.ref(v)
      }
    }
    case *ast.LogicalAndNode: {
      rightLabel := asm.NewUnnamedLabel()
      endLabel := asm.NewUnnamedLabel()
      v := self.tmpVar(node.GetType())
      self.assign(node.GetLeft().GetLocation(), self.ref(v), self.transformExpr(node.GetLeft()))
      self.cjump(node.GetLocation(), self.ref(v), rightLabel, endLabel)
      self.label(rightLabel)
      self.assign(node.GetRight().GetLocation(), self.ref(v), self.transformExpr(node.GetRight()))
      self.label(endLabel)
      if self.isStatement() {
        return nil
      } else {
        return self.ref(v)
      }
    }
    case *ast.LogicalOrNode: {
      rightLabel := asm.NewUnnamedLabel()
      endLabel := asm.NewUnnamedLabel()
      v := self.tmpVar(node.GetType())
      self.assign(node.GetLeft().GetLocation(), self.ref(v), self.transformExpr(node.GetLeft()))
      self.cjump(node.GetLocation(), self.ref(v), endLabel, rightLabel)
      self.label(rightLabel)
      self.assign(node.GetRight().GetLocation(), self.ref(v), self.transformExpr(node.GetRight()))
      self.label(endLabel)
      if self.isStatement() {
        return nil
      } else {
        return self.ref(v)
      }
    }
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
