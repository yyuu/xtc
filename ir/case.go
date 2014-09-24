package ir

type Case struct {
  ClassName string
  Value int64
  Label string // FIXME:
}

func NewCase(value int64, label string) *Case {
  return &Case { "ir.Case", value, label }
}
