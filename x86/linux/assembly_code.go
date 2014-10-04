package x86_linux

import (
  "fmt"
  "strings"
  bs_asm "bitbucket.org/yyuu/bs/asm"
  bs_core "bitbucket.org/yyuu/bs/core"
  bs_x86 "bitbucket.org/yyuu/bs/x86"
)

type AssemblyCode struct {
  NaturalType int
  LabelSymbols *bs_asm.SymbolTable
  virtualStack *bs_x86.VirtualStack
  Assemblies []bs_core.IAssembly
  CommentIndentLevel int
  Statistics *bs_asm.Statistics
}

func NewAssemblyCode(naturalType int, labelSymbols *bs_asm.SymbolTable) *AssemblyCode {
  return &AssemblyCode {
    NaturalType: naturalType,
    LabelSymbols: labelSymbols,
    virtualStack: bs_x86.NewVirtualStack(naturalType),
    Assemblies: []bs_core.IAssembly { },
    CommentIndentLevel: 0,
    Statistics: nil,
  }
}

func (self *AssemblyCode) GetAssemblies() []bs_core.IAssembly {
  return self.Assemblies
}

func (self *AssemblyCode) addAll(assemblies []bs_core.IAssembly) {
  self.Assemblies = append(self.Assemblies, assemblies...)
}

func (self *AssemblyCode) ToSource() string {
  sources := make([]string, len(self.Assemblies))
  for i := range self.Assemblies {
    sources[i] = self.Assemblies[i].ToSource(self.LabelSymbols)
  }
  return strings.Join(sources, "\n") + "\n"
}

func (self *AssemblyCode) GetStatistics() *bs_asm.Statistics {
  if self.Statistics == nil {
    self.Statistics = bs_asm.CollectStatistics(self.Assemblies)
  }
  return self.Statistics
}

func (self *AssemblyCode) doesUses(reg *bs_x86.Register) bool {
  if reg == nil { panic("reg is nil") }
  return self.GetStatistics().DoesRegisterUsed(reg)
}

func (self *AssemblyCode) comment(str string) {
  self.Assemblies = append(self.Assemblies, bs_asm.NewComment(str, self.CommentIndentLevel))
}

func (self *AssemblyCode) indentComment() {
  self.CommentIndentLevel++
}

func (self *AssemblyCode) unindentComment() {
  self.CommentIndentLevel--
}

func (self *AssemblyCode) label1(label *bs_asm.Label) {
  if label == nil { panic("label is nil") }
  self.Assemblies = append(self.Assemblies, label)
}

func (self *AssemblyCode) label2(sym bs_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.label1(bs_asm.NewLabel(sym))
}

func (self *AssemblyCode) reduceLabels() {
  stats := self.GetStatistics()
  result := []bs_core.IAssembly { }
  for i := range self.Assemblies {
    asm := self.Assemblies[i]
    if asm.IsLabel() && !stats.DoesSymbolUsed(asm.(*bs_asm.Label).GetSymbol()) {
      // noop
    } else {
      result = append(result, asm)
    }
  }
  self.Assemblies = result
}

func (self *AssemblyCode) directive(direc string) {
  self.Assemblies = append(self.Assemblies, bs_asm.NewDirective(direc))
}

func (self *AssemblyCode) insn1(op string) {
  instruction := bs_asm.NewInstruction(op, "", []bs_core.IOperand { }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn2(op string, a bs_core.IOperand) {
  if a == nil { panic("a is nil") }
  instruction := bs_asm.NewInstruction(op, "", []bs_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn3(op string, suffix string, a bs_core.IOperand) {
  if a == nil { panic("a is nil") }
  instruction := bs_asm.NewInstruction(op, suffix, []bs_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn4(t int, op string, a bs_core.IOperand) {
  if a == nil { panic("a is nil") }
  instruction := bs_asm.NewInstruction(op, self.typeSuffix(t), []bs_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn5(op string, suffix string, a bs_core.IOperand, b bs_core.IOperand) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  instruction := bs_asm.NewInstruction(op, suffix, []bs_core.IOperand { a, b }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) insn6(t int, op string, a bs_core.IOperand, b bs_core.IOperand) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  instruction := bs_asm.NewInstruction(op, self.typeSuffix(t), []bs_core.IOperand { a, b }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) typeSuffix2(t1, t2 int) string {
  return self.typeSuffix(t1) + self.typeSuffix(t2)
}

func (self *AssemblyCode) typeSuffix(t int) string {
  switch t {
    case bs_asm.TYPE_INT8:  return "b"
    case bs_asm.TYPE_INT16: return "w"
    case bs_asm.TYPE_INT32: return "l"
    case bs_asm.TYPE_INT64: return "q"
    default: {
      panic(fmt.Errorf("unknown register type: %d", t))
    }
  }
}

func (self *AssemblyCode) _file(name string) {
  self.directive(fmt.Sprintf(".file\t%q", name))
}

func (self *AssemblyCode) _text() {
  self.directive(fmt.Sprintf("\t.text"))
}

func (self *AssemblyCode) _data() {
  self.directive(fmt.Sprintf("\t.data"))
}

func (self *AssemblyCode) _section(name string) {
  self.directive(fmt.Sprintf("\t.section\t%s", name))
}

func (self *AssemblyCode) _section2(name, flags, t, group, linkage string) {
  self.directive(fmt.Sprintf("\t.section\t%s,%s,%s,%s,%s", name, flags, t, group, linkage))
}

func (self *AssemblyCode) _globl(sym bs_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf(".globl %s", sym.GetName()))
}

func (self *AssemblyCode) _local(sym bs_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf(".local %s", sym.GetName()))
}

func (self *AssemblyCode) _hidden(sym bs_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.hidden\t%s", sym.GetName()))
}

func (self *AssemblyCode) _comm(sym bs_core.ISymbol, size, alignment int64) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.comm\t%s,%d,%d", sym.GetName(), size, alignment))
}

func (self *AssemblyCode) _align(n int64) {
  self.directive(fmt.Sprintf("\t.align\t%d", n))
}

func (self *AssemblyCode) _type(sym bs_core.ISymbol, t string) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.type\t%s,%s", sym.GetName(), t))
}

func (self *AssemblyCode) _size(sym bs_core.ISymbol, size string) {
  if sym == nil { panic("sym is nil") }
  self.directive(fmt.Sprintf("\t.size\t%s,%s", sym.GetName(), size))
}

func (self *AssemblyCode) _byte(val bs_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".byte\t%s", val))
}

func (self *AssemblyCode) _value(val bs_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".value\t%s", val))
}

func (self *AssemblyCode) _long(val bs_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".long\t%s", val))
}

func (self *AssemblyCode) _quad(val bs_core.ILiteral) {
  if val == nil { panic("val is nil") }
  self.directive(fmt.Sprintf(".quad\t%s", val))
}

func (self *AssemblyCode) _string(str string) {
  self.directive(fmt.Sprintf("\t.string\t%q", str))
}

func (self *AssemblyCode) virtualPush(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.virtualStack.Extend(STACK_WORD_SIZE)
  self.mov3(reg, self.virtualStack.Top())
}

func (self *AssemblyCode) virtualPop(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.mov2(self.virtualStack.Top(), reg)
  self.virtualStack.Rewind(STACK_WORD_SIZE)
}

func (self *AssemblyCode) jmp(label *bs_asm.Label) {
  if label == nil { panic("label is nil") }
  self.insn2("jmp", bs_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *AssemblyCode) jnz(label *bs_asm.Label) {
  if label == nil { panic("label is nil") }
  self.insn2("jnz", bs_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *AssemblyCode) je(label *bs_asm.Label) {
  if label == nil { panic("label is nil") }
  self.insn2("je", bs_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *AssemblyCode) cmp(a bs_core.IOperand, b *bs_x86.Register) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  self.insn6(b.GetTypeId(), "cmp", a, b)
}

func (self *AssemblyCode) sete(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("sete", reg)
}

func (self *AssemblyCode) setne(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setne", reg)
}

func (self *AssemblyCode) seta(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("seta", reg)
}

func (self *AssemblyCode) setae(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setae", reg)
}

func (self *AssemblyCode) setb(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setb", reg)
}

func (self *AssemblyCode) setbe(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setbe", reg)
}

func (self *AssemblyCode) setg(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setg", reg)
}

func (self *AssemblyCode) setge(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setge", reg)
}

func (self *AssemblyCode) setl(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setl", reg)
}

func (self *AssemblyCode) setle(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("setle", reg)
}

func (self *AssemblyCode) test(a *bs_x86.Register, b *bs_x86.Register) {
  if a == nil { panic("a is nil") }
  if b == nil { panic("b is nil") }
  self.insn6(b.GetTypeId(), "test", a, b)
}

func (self *AssemblyCode) push(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn3("push", self.typeSuffix(self.NaturalType), reg)
}

func (self *AssemblyCode) pop(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn3("pop", self.typeSuffix(self.NaturalType), reg)
}

// call function by relative address
func (self *AssemblyCode) call(sym bs_core.ISymbol) {
  if sym == nil { panic("sym is nil") }
  self.insn2("call", bs_asm.NewDirectMemoryReference(sym))
}

// call function by absolute address
func (self *AssemblyCode) callAbsolute(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn2("call", bs_asm.NewAbsoluteAddress(reg))
}

func (self *AssemblyCode) ret() {
  self.insn1("ret")
}

func (self *AssemblyCode) mov1(src *bs_x86.Register, dest *bs_x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(self.NaturalType, "mov", src, dest)
}

// load
func (self *AssemblyCode) mov2(src bs_core.IOperand, dest *bs_x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(dest.GetTypeId(), "mov", src, dest)
}

// save
func (self *AssemblyCode) mov3(src *bs_x86.Register, dest bs_core.IOperand) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(src.GetTypeId(), "mov", src, dest)
}

// for stack access
func (self *AssemblyCode) relocatableMov(src bs_core.IOperand, dest bs_core.IOperand) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  instruction := bs_asm.NewInstruction("mov", self.typeSuffix(self.NaturalType), []bs_core.IOperand { src, dest }, true)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *AssemblyCode) movsx(src *bs_x86.Register, dest *bs_x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn5("movs", self.typeSuffix2(src.GetTypeId(), dest.GetTypeId()), src, dest)
}

func (self *AssemblyCode) movzx(src *bs_x86.Register, dest *bs_x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn5("movz", self.typeSuffix2(src.GetTypeId(), dest.GetTypeId()), src, dest)
}

func (self *AssemblyCode) movzb(src *bs_x86.Register, dest *bs_x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn5("movz", "b"+self.typeSuffix(dest.GetTypeId()), src, dest)
}

func (self *AssemblyCode) lea(src bs_core.IOperand, dest *bs_x86.Register) {
  if src == nil { panic("src is nil") }
  if dest == nil { panic("dest is nil") }
  self.insn6(self.NaturalType, "lea", src, dest)
}

func (self *AssemblyCode) neg(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn4(reg.GetTypeId(), "neg", reg)
}

func (self *AssemblyCode) add(diff bs_core.IOperand, base *bs_x86.Register) {
  if diff == nil { panic("diff is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "add", diff, base)
}

func (self *AssemblyCode) sub(diff bs_core.IOperand, base *bs_x86.Register) {
  if diff == nil { panic("diff is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "sub", diff, base)
}

func (self *AssemblyCode) imul(m bs_core.IOperand, base *bs_x86.Register) {
  if m == nil { panic("m is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "imul", m, base)
}

func (self *AssemblyCode) cltd() {
  self.insn1("cltd")
}

func (self *AssemblyCode) div(base *bs_x86.Register) {
  if base == nil { panic("base is nil") }
  self.insn4(base.GetTypeId(), "div", base)
}

func (self *AssemblyCode) idiv(base *bs_x86.Register) {
  if base == nil { panic("base is nil") }
  self.insn4(base.GetTypeId(), "idiv", base)
}

func (self *AssemblyCode) not(reg *bs_x86.Register) {
  if reg == nil { panic("reg is nil") }
  self.insn4(reg.GetTypeId(), "not", reg)
}

func (self *AssemblyCode) and(bits bs_core.IOperand, base *bs_x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "and", bits, base)
}

func (self *AssemblyCode) or(bits bs_core.IOperand, base *bs_x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "or", bits, base)
}

func (self *AssemblyCode) xor(bits bs_core.IOperand, base *bs_x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "xor", bits, base)
}

func (self *AssemblyCode) sar(bits *bs_x86.Register, base *bs_x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "sar", bits, base)
}

func (self *AssemblyCode) sal(bits *bs_x86.Register, base *bs_x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "sal", bits, base)
}

func (self *AssemblyCode) shr(bits *bs_x86.Register, base *bs_x86.Register) {
  if bits == nil { panic("bits is nil") }
  if base == nil { panic("base is nil") }
  self.insn6(base.GetTypeId(), "shr", bits, base)
}
