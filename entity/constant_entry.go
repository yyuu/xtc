package entity

type ConstantEntry struct {
  s string
}

func NewConstantEntry(s string) *ConstantEntry {
  return &ConstantEntry { s }
}

func (self *ConstantEntry) IsConstantEntry() bool {
  return true
}
