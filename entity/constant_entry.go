package entity

type ConstantEntry struct {
  value string
}

func NewConstantEntry(value string) *ConstantEntry {
  return &ConstantEntry { value }
}

func (self *ConstantEntry) GetValue() string {
  return self.value
}
