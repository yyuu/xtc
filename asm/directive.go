package asm

type Directive struct {
  ClassName string
  Content string
}

func NewDirective(content string) Directive {
  return Directive { "asm.Directive", content }
}

func (self Directive) IsDirective() bool {
  return true
}
