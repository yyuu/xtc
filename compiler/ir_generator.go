package compiler

import (
  "fmt"
  bs_asm "bitbucket.org/yyuu/bs/asm"
  bs_ast "bitbucket.org/yyuu/bs/ast"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_entity "bitbucket.org/yyuu/bs/entity"
  bs_ir "bitbucket.org/yyuu/bs/ir"
  bs_typesys "bitbucket.org/yyuu/bs/typesys"
)

type IRGenerator struct {
  errorHandler *bs_core.ErrorHandler
  options *bs_core.Options
  typeTable *bs_typesys.TypeTable
  exprNestLevel int
  stmts []bs_core.IStmt
  scopeStack []*bs_entity.LocalScope
  breakStack []*bs_asm.Label
  continueStack []*bs_asm.Label
  jumpMap map[string]*jumpEntry
}

type jumpEntry struct {
  label *bs_asm.Label
  numRefered int
  isDefined bool
  location bs_core.Location
}

func newJumpEntry(label *bs_asm.Label) *jumpEntry {
  loc := bs_core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  return &jumpEntry { label, 0, false, loc }
}

func NewIRGenerator(errorHandler *bs_core.ErrorHandler, options *bs_core.Options, table *bs_typesys.TypeTable) *IRGenerator {
  stmts := []bs_core.IStmt { }
  scopeStack := []*bs_entity.LocalScope { }
  breakStack := []*bs_asm.Label { }
  continueStack := []*bs_asm.Label { }
  jumpMap := make(map[string]*jumpEntry)
  return &IRGenerator { errorHandler, options, table, 0, stmts, scopeStack, breakStack, continueStack, jumpMap }
}

func (self *IRGenerator) Generate(ast *bs_ast.AST) (*bs_ir.IR, error) {
  vs := ast.GetDefinedVariables()
  for i := range vs {
    if vs[i].HasInitializer() {
      vs[i].SetIR(self.transformExpr(vs[i].GetInitializer()))
    }
  }
  fs := ast.GetDefinedFunctions()
  for i := range fs {
    fs[i].SetIR(self.compileFunctionBody(fs[i]))
  }
  ir := ast.GenerateIR()
  if self.errorHandler.ErrorOccured() {
    return nil, fmt.Errorf("found %d error(s).", self.errorHandler.GetErrors())
  }
  return ir, nil
}

func (self *IRGenerator) compileFunctionBody(f *bs_entity.DefinedFunction) []bs_core.IStmt {
  self.stmts = []bs_core.IStmt { }
  self.scopeStack = []*bs_entity.LocalScope { }
  self.breakStack = []*bs_asm.Label { }
  self.continueStack = []*bs_asm.Label { }
  self.jumpMap = make(map[string]*jumpEntry)
  self.transformStmt(f.GetBody())
  self.checkJumpLinks(self.jumpMap)
  return self.stmts
}

func (self *IRGenerator) transformStmt(node bs_core.IStmtNode) {
  bs_ast.VisitStmtNode(self, node)
}

func (self *IRGenerator) transformStmtExpr(node bs_core.IExprNode) {
  bs_ast.VisitExprNode(self, node)
}

func (self *IRGenerator) transformExpr(node bs_core.IExprNode) bs_core.IExpr {
  self.exprNestLevel++
  e := bs_ast.VisitExprNode(self, node)
  self.exprNestLevel--
  return e.(bs_core.IExpr)
}

func (self *IRGenerator) isStatement() bool {
  return self.exprNestLevel == 0
}

func (self *IRGenerator) assign(loc bs_core.Location, lhs bs_core.IExpr, rhs bs_core.IExpr) {
  self.stmts = append(self.stmts, bs_ir.NewAssign(loc, self.addressOf(lhs), rhs))
}

func (self *IRGenerator) tmpVar(t bs_core.IType) *bs_entity.DefinedVariable {
  ref := self.typeTable.GetTypeRef(t)
  typeNode := bs_ast.NewTypeNode(bs_core.NewLocation("[builtin:ir_generator]", 0, 0), ref)
  typeNode.SetType(t)
  return self.currentScope().AllocateTmp(typeNode)
}

func (self *IRGenerator) label(label *bs_asm.Label) {
  loc := bs_core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  self.stmts = append(self.stmts, bs_ir.NewLabelStmt(loc, label))
}

func (self *IRGenerator) jump(target *bs_asm.Label) {
  loc := bs_core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  self.stmts = append(self.stmts, bs_ir.NewJump(loc, target))
}

func (self *IRGenerator) cjump(loc bs_core.Location, cond bs_core.IExpr, thenLabel *bs_asm.Label, elseLabel *bs_asm.Label) {
  self.stmts = append(self.stmts, bs_ir.NewCJump(loc, cond, thenLabel, elseLabel))
}

func (self *IRGenerator) pushBreak(label *bs_asm.Label) {
  self.breakStack = append(self.breakStack, label)
}

func (self *IRGenerator) popBreak() *bs_asm.Label {
  if len(self.breakStack) < 1 {
    self.errorHandler.Fatal("unmatched push/pop for break stack")
  }
  label := self.currentBreakTarget()
  self.breakStack = self.breakStack[0:len(self.breakStack)-1]
  return label
}

func (self *IRGenerator) currentBreakTarget() *bs_asm.Label {
  if len(self.breakStack) < 1 {
    self.errorHandler.Fatal("break from out of loop")
  }
  return self.breakStack[len(self.breakStack)-1]
}

func (self *IRGenerator) pushContinue(label *bs_asm.Label) {
  self.continueStack = append(self.continueStack, label)
}

func (self *IRGenerator) popContinue() *bs_asm.Label {
  if len(self.continueStack) < 1 {
    self.errorHandler.Fatal("unmatched push/pop for continue stack")
  }
  label := self.currentContinueTarget()
  self.continueStack = self.continueStack[0:len(self.continueStack)-1]
  return label
}

func (self *IRGenerator) currentContinueTarget() *bs_asm.Label {
  if len(self.continueStack) < 1 {
    self.errorHandler.Fatal("continue from out of loop")
  }
  return self.continueStack[len(self.continueStack)-1]
}

func (self *IRGenerator) transformIndex(node *bs_ast.ArefNode) bs_core.IExpr {
  if node.IsMultiDimension() {
    return bs_ir.NewBin(self.int_t(), bs_ir.OP_ADD,
                     self.transformExpr(node.GetIndex()),
                     bs_ir.NewBin(self.int_t(), bs_ir.OP_MUL,
                               bs_ir.NewInt(self.int_t(), int64(node.GetLength())),
                               self.transformIndex(node.GetExpr().(*bs_ast.ArefNode))))

  } else {
    return self.transformExpr(node.GetIndex())
  }
}

func (self *IRGenerator) transformOpAssign(loc bs_core.Location, op int, lhsType bs_core.IType, lhs bs_core.IExpr, rhs bs_core.IExpr) bs_core.IExpr {
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

func (self *IRGenerator) bin(op int, leftType bs_core.IType, left bs_core.IExpr, right bs_core.IExpr) *bs_ir.Bin {
  if self.isPointerArithmetic(op, leftType) {
    return bs_ir.NewBin(left.GetTypeId(), op, left,
                     bs_ir.NewBin(right.GetTypeId(), bs_ir.OP_MUL,
                               right,
                               self.ptrBaseSize(leftType)))
  } else {
    return bs_ir.NewBin(left.GetTypeId(), op, left, right)
  }
}

func (self *IRGenerator) isPointerDiff(op int, l bs_core.IType, r bs_core.IType) bool {
  return op == bs_ir.OP_SUB && l.IsPointer() && r.IsPointer()
}

func (self *IRGenerator) isPointerArithmetic(op int, operandType bs_core.IType) bool {
  switch op {
    case bs_ir.OP_ADD: return operandType.IsPointer()
    case bs_ir.OP_SUB: return operandType.IsPointer()
    default:        return false
  }
}

func (self *IRGenerator) ptrBaseSize(t bs_core.IType) bs_core.IExpr {
  return bs_ir.NewInt(self.ptrdiff_t(), int64(t.GetBaseType().Size()))
}

func (self *IRGenerator) binOp(uniOp string) int {
  if uniOp == "++" {
    return bs_ir.OP_ADD
  } else {
    if uniOp == "--" {
      return bs_ir.OP_SUB
    } else {
      panic("must not happen")
    }
  }
}

func (self *IRGenerator) addressOf(expr bs_core.IExpr) bs_core.IExpr {
  return expr.GetAddressNode(self.ptr_t())
}

func (self *IRGenerator) ref(ent bs_core.IEntity) *bs_ir.Var {
  return bs_ir.NewVar(self.varType(ent.GetType()), ent)
}

func (self *IRGenerator) mem(ent bs_core.IEntity) *bs_ir.Mem {
  return bs_ir.NewMem(self.asmType(ent.GetType().GetBaseType()), self.ref(ent))
}

func (self *IRGenerator) mem2(expr bs_core.IExpr, t bs_core.IType) *bs_ir.Mem {
  return bs_ir.NewMem(self.asmType(t), expr)
}

func (self *IRGenerator) ptrdiff(n int64) *bs_ir.Int {
  return bs_ir.NewInt(self.ptrdiff_t(), n)
}

func (self *IRGenerator) size(n int64) *bs_ir.Int {
  return bs_ir.NewInt(self.size_t(), n)
}

func (self *IRGenerator) imm(operandType bs_core.IType, n int64) *bs_ir.Int {
  if operandType.IsPointer() {
    return bs_ir.NewInt(self.ptrdiff_t(), n)
  } else {
    return bs_ir.NewInt(self.int_t(), n)
  }
}

func (self *IRGenerator) pointerTo(t bs_core.IType) bs_core.IType {
  return self.typeTable.PointerTo(t)
}

func (self *IRGenerator) asmType(t bs_core.IType) int {
  if t.IsVoid() {
    return self.int_t()
  } else {
    return bs_asm.TypeGet(t.Size())
  }
}

func (self *IRGenerator) varType(t bs_core.IType) int {
  if ! t.IsScalar() {
    return 0
  } else {
    return bs_asm.TypeGet(t.Size())
  }
}

func (self *IRGenerator) int_t() int {
  return bs_asm.TypeGet(self.typeTable.GetIntSize())
}

func (self *IRGenerator) size_t() int {
  return bs_asm.TypeGet(self.typeTable.GetLongSize())
}

func (self *IRGenerator) ptr_t() int {
  return bs_asm.TypeGet(self.typeTable.GetPointerSize())
}

func (self *IRGenerator) ptrdiff_t() int {
  return bs_asm.TypeGet(self.typeTable.GetLongSize())
}

func (self *IRGenerator) currentScope() *bs_entity.LocalScope {
  return self.scopeStack[len(self.scopeStack)-1]
}

func (self *IRGenerator) pushScope(scope *bs_entity.LocalScope) {
  self.scopeStack = append(self.scopeStack, scope)
}

func (self *IRGenerator) popScope() *bs_entity.LocalScope {
  scope := self.currentScope()
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

func (self *IRGenerator) defineLabel(name string, loc bs_core.Location) *bs_asm.Label {
  ent := self.getJumpEntry(name)
  if ent.isDefined {
    self.errorHandler.Errorf("duplicated jump label in %s(): %s", name, name)
  }
  ent.isDefined = true
  ent.location = loc
  return ent.label
}

func (self *IRGenerator) referLabel(name string) *bs_asm.Label {
  ent := self.getJumpEntry(name)
  ent.numRefered++
  return ent.label
}

func (self *IRGenerator) getJumpEntry(name string) *jumpEntry {
  ent := self.jumpMap[name]
  if ent == nil {
    ent = newJumpEntry(bs_asm.NewUnnamedLabel())
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

func (self *IRGenerator) VisitStmtNode(unknown bs_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.BlockNode: {
      self.pushScope(node.GetScope().(*bs_entity.LocalScope))
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
    case *bs_ast.ExprStmtNode: {
      e := bs_ast.VisitExprNode(self, node.GetExpr())
      if e != nil {
        self.errorHandler.Warnf("%s useless expression", node.GetLocation())
      }
      return nil
    }
    case *bs_ast.IfNode: {
      thenLabel := bs_asm.NewUnnamedLabel()
      elseLabel := bs_asm.NewUnnamedLabel()
      endLabel := bs_asm.NewUnnamedLabel()
      cond := self.transformExpr(node.GetCond())
      if node.HasElseBody() {
        self.cjump(node.GetLocation(), cond, thenLabel, endLabel)
        self.label(thenLabel)
        self.transformStmt(node.GetThenBody())
        self.jump(endLabel)
        self.label(elseLabel)
        self.transformStmt(node.GetElseBody())
        self.label(endLabel)
      } else {
        self.cjump(node.GetLocation(), cond, thenLabel, endLabel)
        self.label(thenLabel)
        self.transformStmt(node.GetThenBody())
        self.label(endLabel)
      }
      return nil
    }
    case *bs_ast.SwitchNode: {
      caseNodes := node.GetCases()
      cases := make([]*bs_ir.Case, len(caseNodes))
      endLabel := bs_asm.NewUnnamedLabel()
      defaultLabel := endLabel
      cond := self.transformExpr(node.GetCond())
      for i := range caseNodes {
        c := caseNodes[i].(*bs_ast.CaseNode)
        if c.IsDefault() {
          defaultLabel = c.GetLabel()
        } else {
          values := c.GetValues()
          for j := range values {
            v := self.transformExpr(values[j])
            cases[i] = bs_ir.NewCase(v.(*bs_ir.Int).GetValue(), c.GetLabel())
          }
        }
      }
      self.stmts = append(self.stmts, bs_ir.NewSwitch(node.GetLocation(), cond, cases, defaultLabel, endLabel))
      self.pushBreak(endLabel)
      for i := range caseNodes {
        c := caseNodes[i].(*bs_ast.CaseNode)
        self.label(c.GetLabel())
        self.transformStmt(c.GetBody())
      }
      self.popBreak()
      self.label(endLabel)
      return nil
    }
    case *bs_ast.CaseNode: {
      panic("must not happen")
    }
    case *bs_ast.WhileNode: {
      begLabel := bs_asm.NewUnnamedLabel()
      bodyLabel := bs_asm.NewUnnamedLabel()
      endLabel := bs_asm.NewUnnamedLabel()
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
    case *bs_ast.DoWhileNode: {
      begLabel := bs_asm.NewUnnamedLabel()
      contLabel := bs_asm.NewUnnamedLabel()
      endLabel := bs_asm.NewUnnamedLabel()
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
    case *bs_ast.ForNode: {
      begLabel := bs_asm.NewUnnamedLabel()
      bodyLabel := bs_asm.NewUnnamedLabel()
      contLabel := bs_asm.NewUnnamedLabel()
      endLabel := bs_asm.NewUnnamedLabel()
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
    case *bs_ast.BreakNode: {
      self.jump(self.currentBreakTarget())
      return nil
    }
    case *bs_ast.ContinueNode: {
      self.jump(self.currentContinueTarget())
      return nil
    }
    case *bs_ast.LabelNode: {
      stmt := bs_ir.NewLabelStmt(node.GetLocation(), self.defineLabel(node.GetName(), node.GetLocation()))
      self.stmts = append(self.stmts, stmt)
      if node.GetStmt() != nil {
        self.transformStmt(node.GetStmt())
      }
      return nil
    }
    case *bs_ast.GotoNode: {
      self.jump(self.referLabel(node.GetTarget()))
    }
    case *bs_ast.ReturnNode: {
      var expr bs_core.IExpr
      if node.GetExpr() != nil {
        expr = self.transformExpr(node.GetExpr())
      }
      self.stmts = append(self.stmts, bs_ir.NewReturn(node.GetLocation(), expr))
      return nil
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *IRGenerator) VisitExprNode(unknown bs_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *bs_ast.CondExprNode: {
      thenLabel := bs_asm.NewUnnamedLabel()
      elseLabel := bs_asm.NewUnnamedLabel()
      endLabel := bs_asm.NewUnnamedLabel()
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
    case *bs_ast.LogicalAndNode: {
      rightLabel := bs_asm.NewUnnamedLabel()
      endLabel := bs_asm.NewUnnamedLabel()
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
    case *bs_ast.LogicalOrNode: {
      rightLabel := bs_asm.NewUnnamedLabel()
      endLabel := bs_asm.NewUnnamedLabel()
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
    case *bs_ast.AssignNode: {
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
    case *bs_ast.OpAssignNode: {
      rhs := self.transformExpr(node.GetRHS())
      lhs := self.transformExpr(node.GetLHS())
      t := node.GetLHS().GetType()
      op := bs_ir.OpInternBinary(node.GetOperator(), t.IsSigned())
      return self.transformOpAssign(node.GetLocation(), op, t, lhs, rhs)
    }
    case *bs_ast.PrefixOpNode: {
      t := node.GetExpr().GetType()
      return self.transformOpAssign(node.GetLocation(),
                                    self.binOp(node.GetOperator()),
                                    t,
                                    self.transformExpr(node.GetExpr()), self.imm(t, 1))
    }
    case *bs_ast.SuffixOpNode: {
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
    case *bs_ast.FuncallNode: {
      numArgs := node.NumArgs()
      args := make([]bs_core.IExpr, numArgs)
      for i := range args {
        a := node.GetArg(numArgs-1-i)
        args[i] = self.transformExpr(a)
      }
      call := bs_ir.NewCall(self.asmType(node.GetType()),
                         self.transformExpr(node.GetExpr()),
                         args)
      if self.isStatement() {
        self.stmts = append(self.stmts, bs_ir.NewExprStmt(node.GetLocation(), call))
      } else {
        tmp := self.tmpVar(node.GetType())
        self.assign(node.GetLocation(), self.ref(tmp), call)
        return self.ref(tmp)
      }
    }
    case *bs_ast.BinaryOpNode: {
      right := self.transformExpr(node.GetRight())
      left := self.transformExpr(node.GetLeft())
      op := bs_ir.OpInternBinary(node.GetOperator(), node.GetType().IsSigned())
      t := node.GetType()
      r := node.GetRight().GetType()
      l := node.GetLeft().GetType()
      if self.isPointerDiff(op, l, r) {
        tmp := bs_ir.NewBin(self.asmType(t), op, left, right)
        return bs_ir.NewBin(self.asmType(t), bs_ir.OP_S_DIV, tmp, self.ptrBaseSize(l))
      } else {
        if self.isPointerArithmetic(op, l) {
          return bs_ir.NewBin(self.asmType(t), op,
                           left, 
                           bs_ir.NewBin(self.asmType(r), bs_ir.OP_MUL,
                                     right,
                                     self.ptrBaseSize(l)))
        } else {
          if self.isPointerArithmetic(op, r) {
            return bs_ir.NewBin(self.asmType(t), op,
                             bs_ir.NewBin(self.asmType(l), bs_ir.OP_MUL, left, self.ptrBaseSize(r)),
                             right)
          } else {
            return bs_ir.NewBin(self.asmType(t), op, left, right)
          }
        }
      }
    }
    case *bs_ast.UnaryOpNode: {
      if node.GetOperator() == "+" {
        return self.transformExpr(node.GetExpr())
      } else {
        return bs_ir.NewUni(self.asmType(node.GetType()),
                         bs_ir.OpInternUnary(node.GetOperator()),
                         self.transformExpr(node.GetExpr()))
      }
    }
    case *bs_ast.ArefNode: {
      expr := self.transformExpr(node.GetBaseExpr())
      offset := bs_ir.NewBin(self.ptrdiff_t(), bs_ir.OP_MUL, self.size(int64(node.GetElementSize())), self.transformIndex(node))
      addr := bs_ir.NewBin(self.ptr_t(), bs_ir.OP_ADD, expr, offset)
      return self.mem2(addr, node.GetType())
    }
    case *bs_ast.MemberNode: {
      expr := self.addressOf(self.transformExpr(node.GetExpr()))
      offset := self.ptrdiff(int64(node.GetOffset()))
      addr := bs_ir.NewBin(self.ptr_t(), bs_ir.OP_ADD, expr, offset)
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *bs_ast.PtrMemberNode: {
      expr := self.transformExpr(node.GetExpr())
      offset := self.ptrdiff(int64(node.GetOffset()))
      addr := bs_ir.NewBin(self.ptr_t(), bs_ir.OP_ADD, expr, offset)
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *bs_ast.DereferenceNode: {
      addr := self.transformExpr(node.GetExpr())
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *bs_ast.AddressNode: {
      e := self.transformExpr(node.GetExpr())
      if node.GetExpr().IsLoadable() {
        return self.addressOf(e)
      } else {
        return e
      }
    }
    case *bs_ast.CastNode: {
      if node.IsEffectiveCast() {
        if node.GetExpr().GetType().IsSigned() {
          return bs_ir.NewUni(self.asmType(node.GetType()), bs_ir.OP_S_CAST, self.transformExpr(node.GetExpr()))
        } else {
          return bs_ir.NewUni(self.asmType(node.GetType()), bs_ir.OP_U_CAST, self.transformExpr(node.GetExpr()))
        }
      } else {
        if self.isStatement() {
          self.transformStmtExpr(node.GetExpr())
        } else {
          return self.transformExpr(node.GetExpr())
        }
      }
    }
    case *bs_ast.SizeofExprNode: {
      return bs_ir.NewInt(self.size_t(), int64(node.GetExpr().GetType().AllocSize()))
    }
    case *bs_ast.SizeofTypeNode: {
      return bs_ir.NewInt(self.size_t(), int64(node.GetOperandType().AllocSize()))
    }
    case *bs_ast.VariableNode: {
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
    case *bs_ast.IntegerLiteralNode: {
      return bs_ir.NewInt(self.asmType(node.GetType()), node.GetValue())
    }
    case *bs_ast.StringLiteralNode: {
      return bs_ir.NewStr(self.asmType(node.GetType()), node.GetEntry())
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *IRGenerator) VisitTypeDefinition(unknown bs_core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
