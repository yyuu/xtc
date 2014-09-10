package entity

type ConstantEntry struct {
  s string
}

func (self *ConstantEntry) IsConstantEntry() bool {
  return true
}
