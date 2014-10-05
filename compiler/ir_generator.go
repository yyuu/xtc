package compiler

import (
  "fmt"
  xtc_asm "bitbucket.org/yyuu/xtc/asm"
  xtc_ast "bitbucket.org/yyuu/xtc/ast"
  xtc_core "bitbucket.org/yyuu/xtc/core"
  xtc_entity "bitbucket.org/yyuu/xtc/entity"
  xtc_ir "bitbucket.org/yyuu/xtc/ir"
  xtc_typesys "bitbucket.org/yyuu/xtc/typesys"
)

type IRGenerator struct {
  errorHandler *xtc_core.ErrorHandler
  options *xtc_core.Options
  typeTable *xtc_typesys.TypeTable
  exprNestLevel int
  stmts []xtc_core.IStmt
  scopeStack []*xtc_entity.LocalScope
  breakStack []*xtc_asm.Label
  continueStack []*xtc_asm.Label
  jumpMap map[string]*jumpEntry
}

type jumpEntry struct {
  label *xtc_asm.Label
  numRefered int
  isDefined bool
  location xtc_core.Location
}

func newJumpEntry(label *xtc_asm.Label) *jumpEntry {
  loc := xtc_core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  return &jumpEntry { label, 0, false, loc }
}

func NewIRGenerator(errorHandler *xtc_core.ErrorHandler, options *xtc_core.Options, table *xtc_typesys.TypeTable) *IRGenerator {
  stmts := []xtc_core.IStmt { }
  scopeStack := []*xtc_entity.LocalScope { }
  breakStack := []*xtc_asm.Label { }
  continueStack := []*xtc_asm.Label { }
  jumpMap := make(map[string]*jumpEntry)
  return &IRGenerator { errorHandler, options, table, 0, stmts, scopeStack, breakStack, continueStack, jumpMap }
}

func (self *IRGenerator) Generate(ast *xtc_ast.AST) (*xtc_ir.IR, error) {
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

func (self *IRGenerator) compileFunctionBody(f *xtc_entity.DefinedFunction) []xtc_core.IStmt {
  self.stmts = []xtc_core.IStmt { }
  self.scopeStack = []*xtc_entity.LocalScope { }
  self.breakStack = []*xtc_asm.Label { }
  self.continueStack = []*xtc_asm.Label { }
  self.jumpMap = make(map[string]*jumpEntry)
  self.transformStmt(f.GetBody())
  self.checkJumpLinks(self.jumpMap)
  return self.stmts
}

func (self *IRGenerator) transformStmt(node xtc_core.IStmtNode) {
  xtc_ast.VisitStmtNode(self, node)
}

func (self *IRGenerator) transformStmtExpr(node xtc_core.IExprNode) {
  xtc_ast.VisitExprNode(self, node)
}

func (self *IRGenerator) transformExpr(node xtc_core.IExprNode) xtc_core.IExpr {
  self.exprNestLevel++
  e := xtc_ast.VisitExprNode(self, node)
  self.exprNestLevel--
  return e.(xtc_core.IExpr)
}

func (self *IRGenerator) isStatement() bool {
  return self.exprNestLevel == 0
}

func (self *IRGenerator) assign(loc xtc_core.Location, lhs xtc_core.IExpr, rhs xtc_core.IExpr) {
  self.stmts = append(self.stmts, xtc_ir.NewAssign(loc, self.addressOf(lhs), rhs))
}

func (self *IRGenerator) tmpVar(t xtc_core.IType) *xtc_entity.DefinedVariable {
  ref := self.typeTable.GetTypeRef(t)
  typeNode := xtc_ast.NewTypeNode(xtc_core.NewLocation("[builtin:ir_generator]", 0, 0), ref)
  typeNode.SetType(t)
  return self.currentScope().AllocateTmp(typeNode)
}

func (self *IRGenerator) label(label *xtc_asm.Label) {
  loc := xtc_core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  self.stmts = append(self.stmts, xtc_ir.NewLabelStmt(loc, label))
}

func (self *IRGenerator) jump(target *xtc_asm.Label) {
  loc := xtc_core.NewLocation("[builtin:ir_generator]", 0, 0) // FIXME:
  self.stmts = append(self.stmts, xtc_ir.NewJump(loc, target))
}

func (self *IRGenerator) cjump(loc xtc_core.Location, cond xtc_core.IExpr, thenLabel *xtc_asm.Label, elseLabel *xtc_asm.Label) {
  self.stmts = append(self.stmts, xtc_ir.NewCJump(loc, cond, thenLabel, elseLabel))
}

func (self *IRGenerator) pushBreak(label *xtc_asm.Label) {
  self.breakStack = append(self.breakStack, label)
}

func (self *IRGenerator) popBreak() *xtc_asm.Label {
  if len(self.breakStack) < 1 {
    self.errorHandler.Fatal("unmatched push/pop for break stack")
  }
  label := self.currentBreakTarget()
  self.breakStack = self.breakStack[0:len(self.breakStack)-1]
  return label
}

func (self *IRGenerator) currentBreakTarget() *xtc_asm.Label {
  if len(self.breakStack) < 1 {
    self.errorHandler.Fatal("break from out of loop")
  }
  return self.breakStack[len(self.breakStack)-1]
}

func (self *IRGenerator) pushContinue(label *xtc_asm.Label) {
  self.continueStack = append(self.continueStack, label)
}

func (self *IRGenerator) popContinue() *xtc_asm.Label {
  if len(self.continueStack) < 1 {
    self.errorHandler.Fatal("unmatched push/pop for continue stack")
  }
  label := self.currentContinueTarget()
  self.continueStack = self.continueStack[0:len(self.continueStack)-1]
  return label
}

func (self *IRGenerator) currentContinueTarget() *xtc_asm.Label {
  if len(self.continueStack) < 1 {
    self.errorHandler.Fatal("continue from out of loop")
  }
  return self.continueStack[len(self.continueStack)-1]
}

func (self *IRGenerator) transformIndex(node *xtc_ast.ArefNode) xtc_core.IExpr {
  if node.IsMultiDimension() {
    return xtc_ir.NewBin(self.int_t(), xtc_ir.OP_ADD,
                     self.transformExpr(node.GetIndex()),
                     xtc_ir.NewBin(self.int_t(), xtc_ir.OP_MUL,
                               xtc_ir.NewInt(self.int_t(), int64(node.GetLength())),
                               self.transformIndex(node.GetExpr().(*xtc_ast.ArefNode))))

  } else {
    return self.transformExpr(node.GetIndex())
  }
}

func (self *IRGenerator) transformOpAssign(loc xtc_core.Location, op int, lhsType xtc_core.IType, lhs xtc_core.IExpr, rhs xtc_core.IExpr) xtc_core.IExpr {
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

func (self *IRGenerator) bin(op int, leftType xtc_core.IType, left xtc_core.IExpr, right xtc_core.IExpr) *xtc_ir.Bin {
  if self.isPointerArithmetic(op, leftType) {
    return xtc_ir.NewBin(left.GetTypeId(), op, left,
                     xtc_ir.NewBin(right.GetTypeId(), xtc_ir.OP_MUL,
                               right,
                               self.ptrBaseSize(leftType)))
  } else {
    return xtc_ir.NewBin(left.GetTypeId(), op, left, right)
  }
}

func (self *IRGenerator) isPointerDiff(op int, l xtc_core.IType, r xtc_core.IType) bool {
  return op == xtc_ir.OP_SUB && l.IsPointer() && r.IsPointer()
}

func (self *IRGenerator) isPointerArithmetic(op int, operandType xtc_core.IType) bool {
  switch op {
    case xtc_ir.OP_ADD: return operandType.IsPointer()
    case xtc_ir.OP_SUB: return operandType.IsPointer()
    default:        return false
  }
}

func (self *IRGenerator) ptrBaseSize(t xtc_core.IType) xtc_core.IExpr {
  return xtc_ir.NewInt(self.ptrdiff_t(), int64(t.GetBaseType().Size()))
}

func (self *IRGenerator) binOp(uniOp string) int {
  if uniOp == "++" {
    return xtc_ir.OP_ADD
  } else {
    if uniOp == "--" {
      return xtc_ir.OP_SUB
    } else {
      panic("must not happen")
    }
  }
}

func (self *IRGenerator) addressOf(expr xtc_core.IExpr) xtc_core.IExpr {
  return expr.GetAddressNode(self.ptr_t())
}

func (self *IRGenerator) ref(ent xtc_core.IEntity) *xtc_ir.Var {
  return xtc_ir.NewVar(self.varType(ent.GetType()), ent)
}

func (self *IRGenerator) mem(ent xtc_core.IEntity) *xtc_ir.Mem {
  return xtc_ir.NewMem(self.asmType(ent.GetType().GetBaseType()), self.ref(ent))
}

func (self *IRGenerator) mem2(expr xtc_core.IExpr, t xtc_core.IType) *xtc_ir.Mem {
  return xtc_ir.NewMem(self.asmType(t), expr)
}

func (self *IRGenerator) ptrdiff(n int64) *xtc_ir.Int {
  return xtc_ir.NewInt(self.ptrdiff_t(), n)
}

func (self *IRGenerator) size(n int64) *xtc_ir.Int {
  return xtc_ir.NewInt(self.size_t(), n)
}

func (self *IRGenerator) imm(operandType xtc_core.IType, n int64) *xtc_ir.Int {
  if operandType.IsPointer() {
    return xtc_ir.NewInt(self.ptrdiff_t(), n)
  } else {
    return xtc_ir.NewInt(self.int_t(), n)
  }
}

func (self *IRGenerator) pointerTo(t xtc_core.IType) xtc_core.IType {
  return self.typeTable.PointerTo(t)
}

func (self *IRGenerator) asmType(t xtc_core.IType) int {
  if t.IsVoid() {
    return self.int_t()
  } else {
    return xtc_asm.TypeGet(t.Size())
  }
}

func (self *IRGenerator) varType(t xtc_core.IType) int {
  if ! t.IsScalar() {
    return 0
  } else {
    return xtc_asm.TypeGet(t.Size())
  }
}

func (self *IRGenerator) int_t() int {
  return xtc_asm.TypeGet(self.typeTable.GetIntSize())
}

func (self *IRGenerator) size_t() int {
  return xtc_asm.TypeGet(self.typeTable.GetLongSize())
}

func (self *IRGenerator) ptr_t() int {
  return xtc_asm.TypeGet(self.typeTable.GetPointerSize())
}

func (self *IRGenerator) ptrdiff_t() int {
  return xtc_asm.TypeGet(self.typeTable.GetLongSize())
}

func (self *IRGenerator) currentScope() *xtc_entity.LocalScope {
  return self.scopeStack[len(self.scopeStack)-1]
}

func (self *IRGenerator) pushScope(scope *xtc_entity.LocalScope) {
  self.scopeStack = append(self.scopeStack, scope)
}

func (self *IRGenerator) popScope() *xtc_entity.LocalScope {
  scope := self.currentScope()
  self.scopeStack = self.scopeStack[0:len(self.scopeStack)-1]
  return scope
}

func (self *IRGenerator) defineLabel(name string, loc xtc_core.Location) *xtc_asm.Label {
  ent := self.getJumpEntry(name)
  if ent.isDefined {
    self.errorHandler.Errorf("duplicated jump label in %s(): %s", name, name)
  }
  ent.isDefined = true
  ent.location = loc
  return ent.label
}

func (self *IRGenerator) referLabel(name string) *xtc_asm.Label {
  ent := self.getJumpEntry(name)
  ent.numRefered++
  return ent.label
}

func (self *IRGenerator) getJumpEntry(name string) *jumpEntry {
  ent := self.jumpMap[name]
  if ent == nil {
    ent = newJumpEntry(xtc_asm.NewUnnamedLabel())
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

func (self *IRGenerator) VisitStmtNode(unknown xtc_core.IStmtNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.BlockNode: {
      self.pushScope(node.GetScope().(*xtc_entity.LocalScope))
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
    case *xtc_ast.ExprStmtNode: {
      e := xtc_ast.VisitExprNode(self, node.GetExpr())
      if e != nil {
        self.errorHandler.Warnf("%s useless expression", node.GetLocation())
      }
      return nil
    }
    case *xtc_ast.IfNode: {
      thenLabel := xtc_asm.NewUnnamedLabel()
      elseLabel := xtc_asm.NewUnnamedLabel()
      endLabel := xtc_asm.NewUnnamedLabel()
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
    case *xtc_ast.SwitchNode: {
      caseNodes := node.GetCases()
      cases := make([]*xtc_ir.Case, len(caseNodes))
      endLabel := xtc_asm.NewUnnamedLabel()
      defaultLabel := endLabel
      cond := self.transformExpr(node.GetCond())
      for i := range caseNodes {
        c := caseNodes[i].(*xtc_ast.CaseNode)
        if c.IsDefault() {
          defaultLabel = c.GetLabel()
        } else {
          values := c.GetValues()
          for j := range values {
            v := self.transformExpr(values[j])
            cases[i] = xtc_ir.NewCase(v.(*xtc_ir.Int).GetValue(), c.GetLabel())
          }
        }
      }
      self.stmts = append(self.stmts, xtc_ir.NewSwitch(node.GetLocation(), cond, cases, defaultLabel, endLabel))
      self.pushBreak(endLabel)
      for i := range caseNodes {
        c := caseNodes[i].(*xtc_ast.CaseNode)
        self.label(c.GetLabel())
        self.transformStmt(c.GetBody())
      }
      self.popBreak()
      self.label(endLabel)
      return nil
    }
    case *xtc_ast.CaseNode: {
      panic("must not happen")
    }
    case *xtc_ast.WhileNode: {
      begLabel := xtc_asm.NewUnnamedLabel()
      bodyLabel := xtc_asm.NewUnnamedLabel()
      endLabel := xtc_asm.NewUnnamedLabel()
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
    case *xtc_ast.DoWhileNode: {
      begLabel := xtc_asm.NewUnnamedLabel()
      contLabel := xtc_asm.NewUnnamedLabel()
      endLabel := xtc_asm.NewUnnamedLabel()
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
    case *xtc_ast.ForNode: {
      begLabel := xtc_asm.NewUnnamedLabel()
      bodyLabel := xtc_asm.NewUnnamedLabel()
      contLabel := xtc_asm.NewUnnamedLabel()
      endLabel := xtc_asm.NewUnnamedLabel()
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
    case *xtc_ast.BreakNode: {
      self.jump(self.currentBreakTarget())
      return nil
    }
    case *xtc_ast.ContinueNode: {
      self.jump(self.currentContinueTarget())
      return nil
    }
    case *xtc_ast.LabelNode: {
      stmt := xtc_ir.NewLabelStmt(node.GetLocation(), self.defineLabel(node.GetName(), node.GetLocation()))
      self.stmts = append(self.stmts, stmt)
      if node.GetStmt() != nil {
        self.transformStmt(node.GetStmt())
      }
      return nil
    }
    case *xtc_ast.GotoNode: {
      self.jump(self.referLabel(node.GetTarget()))
    }
    case *xtc_ast.ReturnNode: {
      var expr xtc_core.IExpr
      if node.GetExpr() != nil {
        expr = self.transformExpr(node.GetExpr())
      }
      self.stmts = append(self.stmts, xtc_ir.NewReturn(node.GetLocation(), expr))
      return nil
    }
    default: {
      visitStmtNode(self, unknown)
    }
  }
  return nil
}

func (self *IRGenerator) VisitExprNode(unknown xtc_core.IExprNode) interface{} {
  switch node := unknown.(type) {
    case *xtc_ast.CondExprNode: {
      thenLabel := xtc_asm.NewUnnamedLabel()
      elseLabel := xtc_asm.NewUnnamedLabel()
      endLabel := xtc_asm.NewUnnamedLabel()
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
    case *xtc_ast.LogicalAndNode: {
      rightLabel := xtc_asm.NewUnnamedLabel()
      endLabel := xtc_asm.NewUnnamedLabel()
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
    case *xtc_ast.LogicalOrNode: {
      rightLabel := xtc_asm.NewUnnamedLabel()
      endLabel := xtc_asm.NewUnnamedLabel()
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
    case *xtc_ast.AssignNode: {
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
    case *xtc_ast.OpAssignNode: {
      rhs := self.transformExpr(node.GetRHS())
      lhs := self.transformExpr(node.GetLHS())
      t := node.GetLHS().GetType()
      op := xtc_ir.OpInternBinary(node.GetOperator(), t.IsSigned())
      return self.transformOpAssign(node.GetLocation(), op, t, lhs, rhs)
    }
    case *xtc_ast.PrefixOpNode: {
      t := node.GetExpr().GetType()
      return self.transformOpAssign(node.GetLocation(),
                                    self.binOp(node.GetOperator()),
                                    t,
                                    self.transformExpr(node.GetExpr()), self.imm(t, 1))
    }
    case *xtc_ast.SuffixOpNode: {
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
    case *xtc_ast.FuncallNode: {
      numArgs := node.NumArgs()
      args := make([]xtc_core.IExpr, numArgs)
      for i := range args {
        args[i] = self.transformExpr(node.GetArg(i))
      }
      call := xtc_ir.NewCall(self.asmType(node.GetType()),
                         self.transformExpr(node.GetExpr()),
                         args)
      if self.isStatement() {
        self.stmts = append(self.stmts, xtc_ir.NewExprStmt(node.GetLocation(), call))
      } else {
        tmp := self.tmpVar(node.GetType())
        self.assign(node.GetLocation(), self.ref(tmp), call)
        return self.ref(tmp)
      }
    }
    case *xtc_ast.BinaryOpNode: {
      right := self.transformExpr(node.GetRight())
      left := self.transformExpr(node.GetLeft())
      op := xtc_ir.OpInternBinary(node.GetOperator(), node.GetType().IsSigned())
      t := node.GetType()
      r := node.GetRight().GetType()
      l := node.GetLeft().GetType()
      if self.isPointerDiff(op, l, r) {
        tmp := xtc_ir.NewBin(self.asmType(t), op, left, right)
        return xtc_ir.NewBin(self.asmType(t), xtc_ir.OP_S_DIV, tmp, self.ptrBaseSize(l))
      } else {
        if self.isPointerArithmetic(op, l) {
          return xtc_ir.NewBin(self.asmType(t), op,
                           left, 
                           xtc_ir.NewBin(self.asmType(r), xtc_ir.OP_MUL,
                                     right,
                                     self.ptrBaseSize(l)))
        } else {
          if self.isPointerArithmetic(op, r) {
            return xtc_ir.NewBin(self.asmType(t), op,
                             xtc_ir.NewBin(self.asmType(l), xtc_ir.OP_MUL, left, self.ptrBaseSize(r)),
                             right)
          } else {
            return xtc_ir.NewBin(self.asmType(t), op, left, right)
          }
        }
      }
    }
    case *xtc_ast.UnaryOpNode: {
      if node.GetOperator() == "+" {
        return self.transformExpr(node.GetExpr())
      } else {
        return xtc_ir.NewUni(self.asmType(node.GetType()),
                         xtc_ir.OpInternUnary(node.GetOperator()),
                         self.transformExpr(node.GetExpr()))
      }
    }
    case *xtc_ast.ArefNode: {
      expr := self.transformExpr(node.GetBaseExpr())
      offset := xtc_ir.NewBin(self.ptrdiff_t(), xtc_ir.OP_MUL, self.size(int64(node.GetElementSize())), self.transformIndex(node))
      addr := xtc_ir.NewBin(self.ptr_t(), xtc_ir.OP_ADD, expr, offset)
      return self.mem2(addr, node.GetType())
    }
    case *xtc_ast.MemberNode: {
      expr := self.addressOf(self.transformExpr(node.GetExpr()))
      offset := self.ptrdiff(int64(node.GetOffset()))
      addr := xtc_ir.NewBin(self.ptr_t(), xtc_ir.OP_ADD, expr, offset)
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *xtc_ast.PtrMemberNode: {
      expr := self.transformExpr(node.GetExpr())
      offset := self.ptrdiff(int64(node.GetOffset()))
      addr := xtc_ir.NewBin(self.ptr_t(), xtc_ir.OP_ADD, expr, offset)
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *xtc_ast.DereferenceNode: {
      addr := self.transformExpr(node.GetExpr())
      if node.IsLoadable() {
        return self.mem2(addr, node.GetType())
      } else {
        return addr
      }
    }
    case *xtc_ast.AddressNode: {
      e := self.transformExpr(node.GetExpr())
      if node.GetExpr().IsLoadable() {
        return self.addressOf(e)
      } else {
        return e
      }
    }
    case *xtc_ast.CastNode: {
      if node.IsEffectiveCast() {
        if node.GetExpr().GetType().IsSigned() {
          return xtc_ir.NewUni(self.asmType(node.GetType()), xtc_ir.OP_S_CAST, self.transformExpr(node.GetExpr()))
        } else {
          return xtc_ir.NewUni(self.asmType(node.GetType()), xtc_ir.OP_U_CAST, self.transformExpr(node.GetExpr()))
        }
      } else {
        if self.isStatement() {
          self.transformStmtExpr(node.GetExpr())
        } else {
          return self.transformExpr(node.GetExpr())
        }
      }
    }
    case *xtc_ast.SizeofExprNode: {
      return xtc_ir.NewInt(self.size_t(), int64(node.GetExpr().GetType().AllocSize()))
    }
    case *xtc_ast.SizeofTypeNode: {
      return xtc_ir.NewInt(self.size_t(), int64(node.GetOperandType().AllocSize()))
    }
    case *xtc_ast.VariableNode: {
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
    case *xtc_ast.IntegerLiteralNode: {
      return xtc_ir.NewInt(self.asmType(node.GetType()), node.GetValue())
    }
    case *xtc_ast.StringLiteralNode: {
      return xtc_ir.NewStr(self.asmType(node.GetType()), node.GetEntry())
    }
    default: {
      visitExprNode(self, unknown)
    }
  }
  return nil
}

func (self *IRGenerator) VisitTypeDefinition(unknown xtc_core.ITypeDefinition) interface{} {
  visitTypeDefinition(self, unknown)
  return nil
}
