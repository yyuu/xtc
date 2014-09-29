package sysdep

import (
  "fmt"
  bs_asm "bitbucket.org/yyuu/bs/asm"
  bs_core "bitbucket.org/yyuu/bs/core"
)

type LinuxX86AssemblyCode struct {
  NaturalType int
  StackWordSize int64
  LabelSymbols *bs_asm.SymbolTable
  virtualStack *x86VirtualStack
  Assemblies []bs_core.IAssembly
  commentIndentLevel int
}

func NewLinuxX86AssemblyCode(naturalType int, stackWordSize int64, labelSymbols *bs_asm.SymbolTable) *LinuxX86AssemblyCode {
  virtualStack := newX86VirtualStack(naturalType, stackWordSize)
  assemblies := []bs_core.IAssembly { }
  return &LinuxX86AssemblyCode { naturalType, stackWordSize, labelSymbols, virtualStack, assemblies, 0 }
}

func (self *LinuxX86AssemblyCode) comment(str string) {
  self.Assemblies = append(self.Assemblies, bs_asm.NewComment(str, self.commentIndentLevel))
}

func (self *LinuxX86AssemblyCode) indentComment() {
  self.commentIndentLevel++
}

func (self *LinuxX86AssemblyCode) unindentComment() {
  self.commentIndentLevel--
}

func (self *LinuxX86AssemblyCode) label(label *bs_asm.Label) {
  self.Assemblies = append(self.Assemblies, label)
}

func (self *LinuxX86AssemblyCode) directive(direc string) {
  self.Assemblies = append(self.Assemblies, bs_asm.NewDirective(direc))
}

func (self *LinuxX86AssemblyCode) insn1(op string) {
  instruction := bs_asm.NewInstruction(op, "", []bs_core.IOperand { }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *LinuxX86AssemblyCode) insn2(op string, a bs_core.IOperand) {
  instruction := bs_asm.NewInstruction(op, "", []bs_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *LinuxX86AssemblyCode) insn3(op string, suffix string, a bs_core.IOperand) {
  instruction := bs_asm.NewInstruction(op, suffix, []bs_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *LinuxX86AssemblyCode) insn4(t int, op string, a bs_core.IOperand) {
  instruction := bs_asm.NewInstruction(op, self.typeSuffix(t), []bs_core.IOperand { a }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *LinuxX86AssemblyCode) insn5(op string, suffix string, a bs_core.IOperand, b bs_core.IOperand) {
  instruction := bs_asm.NewInstruction(op, suffix, []bs_core.IOperand { a, b }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *LinuxX86AssemblyCode) insn6(t int, op string, a bs_core.IOperand, b bs_core.IOperand) {
  instruction := bs_asm.NewInstruction(op, self.typeSuffix(t), []bs_core.IOperand { a, b }, false)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *LinuxX86AssemblyCode) typeSuffix2(t1, t2 int) string {
  return self.typeSuffix(t1) + self.typeSuffix(t2)
}

func (self *LinuxX86AssemblyCode) typeSuffix(t int) string {
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

func (self *LinuxX86AssemblyCode) _file(name string) {
  self.directive(fmt.Sprintf(".file\t%s", name))
}

func (self *LinuxX86AssemblyCode) _text() {
  self.directive(fmt.Sprintf("\t.text"))
}

func (self *LinuxX86AssemblyCode) _data() {
  self.directive(fmt.Sprintf("\t.data"))
}

func (self *LinuxX86AssemblyCode) _section(name string) {
  self.directive(fmt.Sprintf("\t.section\t%s", name))
}

func (self *LinuxX86AssemblyCode) _section2(name, flags, t, group, linkage string) {
  self.directive(fmt.Sprintf("\t.section\t%s,%s,%s,%s,%s", name, flags, t, group, linkage))
}

func (self *LinuxX86AssemblyCode) _globl(sym bs_core.ISymbol) {
  self.directive(fmt.Sprintf(".globl %s", sym.GetName()))
}

func (self *LinuxX86AssemblyCode) _local(sym bs_core.ISymbol) {
  self.directive(fmt.Sprintf(".local %s", sym.GetName()))
}

func (self *LinuxX86AssemblyCode) _hidden(sym bs_core.ISymbol) {
  self.directive(fmt.Sprintf("\t.hidden\t%s", sym.GetName()))
}

func (self *LinuxX86AssemblyCode) _comm(sym bs_core.ISymbol, size, alignment int64) {
  self.directive(fmt.Sprintf("\t.comm\t%s,%d,%d", sym.GetName(), size, alignment))
}

func (self *LinuxX86AssemblyCode) _align(n int64) {
  self.directive(fmt.Sprintf("\t.align\t%d", n))
}

func (self *LinuxX86AssemblyCode) _type(sym bs_core.ISymbol, t string) {
  self.directive(fmt.Sprintf("\t.type\t%s,%s", sym.GetName(), t))
}

func (self *LinuxX86AssemblyCode) _size(sym bs_core.ISymbol, size string) {
  self.directive(fmt.Sprintf("\t.size\t%s,%s", sym.GetName(), size))
}

func (self *LinuxX86AssemblyCode) _byte(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".byte\t%s", val))
}

func (self *LinuxX86AssemblyCode) _value(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".value\t%s", val))
}

func (self *LinuxX86AssemblyCode) _long(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".long\t%s", val))
}

func (self *LinuxX86AssemblyCode) _quad(val bs_core.ILiteral) {
  self.directive(fmt.Sprintf(".quad\t%s", val))
}

func (self *LinuxX86AssemblyCode) _string(str string) {
  self.directive(fmt.Sprintf("\t.string\t%s", str))
}

func (self *LinuxX86AssemblyCode) virtualPush(reg *x86Register) {
  self.virtualStack.extend(self.StackWordSize)
  self.mov3(reg, self.virtualStack.top())
}

func (self *LinuxX86AssemblyCode) virtualPop(reg *x86Register) {
  self.mov2(self.virtualStack.top(), reg)
  self.virtualStack.rewind(self.StackWordSize)
}

func (self *LinuxX86AssemblyCode) jmp(label *bs_asm.Label) {
  self.insn2("jmp", bs_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *LinuxX86AssemblyCode) jnz(label *bs_asm.Label) {
  self.insn2("jnz", bs_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *LinuxX86AssemblyCode) je(label *bs_asm.Label) {
  self.insn2("je", bs_asm.NewDirectMemoryReference(label.GetSymbol()))
}

func (self *LinuxX86AssemblyCode) cmp(a bs_core.IOperand, b *x86Register) {
  self.insn6(b.GetType(), "cmp", a, b)
}

func (self *LinuxX86AssemblyCode) sete(reg *x86Register) {
  self.insn2("sete", reg)
}

func (self *LinuxX86AssemblyCode) setne(reg *x86Register) {
  self.insn2("setne", reg)
}

func (self *LinuxX86AssemblyCode) seta(reg *x86Register) {
  self.insn2("seta", reg)
}

func (self *LinuxX86AssemblyCode) setae(reg *x86Register) {
  self.insn2("setae", reg)
}

func (self *LinuxX86AssemblyCode) setb(reg *x86Register) {
  self.insn2("setb", reg)
}

func (self *LinuxX86AssemblyCode) setbe(reg *x86Register) {
  self.insn2("setbe", reg)
}

func (self *LinuxX86AssemblyCode) setg(reg *x86Register) {
  self.insn2("setg", reg)
}

func (self *LinuxX86AssemblyCode) setge(reg *x86Register) {
  self.insn2("setge", reg)
}

func (self *LinuxX86AssemblyCode) setl(reg *x86Register) {
  self.insn2("setl", reg)
}

func (self *LinuxX86AssemblyCode) setle(reg *x86Register) {
  self.insn2("setle", reg)
}

func (self *LinuxX86AssemblyCode) test(a *x86Register, b *x86Register) {
  self.insn6(b.GetType(), "test", a, b)
}

func (self *LinuxX86AssemblyCode) push(reg *x86Register) {
  self.insn3("push", self.typeSuffix(self.NaturalType), reg)
}

func (self *LinuxX86AssemblyCode) pop(reg *x86Register) {
  self.insn3("pop", self.typeSuffix(self.NaturalType), reg)
}

// call function by relative address
func (self *LinuxX86AssemblyCode) call(sym bs_core.ISymbol) {
  self.insn2("call", bs_asm.NewDirectMemoryReference(sym))
}

// call function by absolute address
func (self *LinuxX86AssemblyCode) callAbsolute(reg *x86Register) {
  self.insn2("call", bs_asm.NewAbsoluteAddress(reg))
}

func (self *LinuxX86AssemblyCode) ret() {
  self.insn1("ret")
}

func (self *LinuxX86AssemblyCode) mov1(src *x86Register, dest *x86Register) {
  self.insn6(self.NaturalType, "mov", src, dest)
}

// load
func (self *LinuxX86AssemblyCode) mov2(src bs_core.IOperand, dest *x86Register) {
  self.insn6(dest.GetType(), "mov", src, dest)
}

// save
func (self *LinuxX86AssemblyCode) mov3(src *x86Register, dest bs_core.IOperand) {
  self.insn6(src.GetType(), "mov", src, dest)
}

// for stack access
func (self *LinuxX86AssemblyCode) relocatableMov(src bs_core.IOperand, dest bs_core.IOperand) {
  instruction := bs_asm.NewInstruction("mov", self.typeSuffix(self.NaturalType), []bs_core.IOperand { src, dest }, true)
  self.Assemblies = append(self.Assemblies, instruction)
}

func (self *LinuxX86AssemblyCode) movsx(src *x86Register, dest *x86Register) {
  self.insn5("movs", self.typeSuffix2(src.GetType(), dest.GetType()), src, dest)
}

func (self *LinuxX86AssemblyCode) movzx(src *x86Register, dest *x86Register) {
  self.insn5("movz", self.typeSuffix2(src.GetType(), dest.GetType()), src, dest)
}

func (self *LinuxX86AssemblyCode) movzb(src *x86Register, dest *x86Register) {
  self.insn5("movz", "b"+self.typeSuffix(dest.GetType()), src, dest)
}

func (self *LinuxX86AssemblyCode) lea(src bs_core.IOperand, dest *x86Register) {
  self.insn6(self.NaturalType, "lea", src, dest)
}

func (self *LinuxX86AssemblyCode) neg(reg *x86Register) {
  self.insn4(reg.GetType(), "neg", reg)
}

func (self *LinuxX86AssemblyCode) add(diff bs_core.IOperand, base *x86Register) {
  self.insn6(base.GetType(), "add", diff, base)
}

func (self *LinuxX86AssemblyCode) sub(diff bs_core.IOperand, base *x86Register) {
  self.insn6(base.GetType(), "sub", diff, base)
}

func (self *LinuxX86AssemblyCode) imul(m bs_core.IOperand, base *x86Register) {
  self.insn6(base.GetType(), "imul", m, base)
}

func (self *LinuxX86AssemblyCode) cltd() {
  self.insn1("cltd")
}

func (self *LinuxX86AssemblyCode) div(base *x86Register) {
  self.insn4(base.GetType(), "div", base)
}

func (self *LinuxX86AssemblyCode) idiv(base *x86Register) {
  self.insn4(base.GetType(), "idiv", base)
}

func (self *LinuxX86AssemblyCode) not(reg *x86Register) {
  self.insn4(reg.GetType(), "not", reg)
}

func (self *LinuxX86AssemblyCode) and(bits bs_core.IOperand, base *x86Register) {
  self.insn6(base.GetType(), "and", bits, base)
}

func (self *LinuxX86AssemblyCode) or(bits bs_core.IOperand, base *x86Register) {
  self.insn6(base.GetType(), "or", bits, base)
}

func (self *LinuxX86AssemblyCode) xor(bits bs_core.IOperand, base *x86Register) {
  self.insn6(base.GetType(), "xor", bits, base)
}

func (self *LinuxX86AssemblyCode) sar(bits *x86Register, base *x86Register) {
  self.insn6(base.GetType(), "sar", bits, base)
}

func (self *LinuxX86AssemblyCode) sal(bits *x86Register, base *x86Register) {
  self.insn6(base.GetType(), "sal", bits, base)
}

func (self *LinuxX86AssemblyCode) shr(bits *x86Register, base *x86Register) {
  self.insn6(base.GetType(), "shr", bits, base)
}
