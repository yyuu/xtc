package entity

type ConstantTable struct {
  table map[string]*ConstantEntry
}

func NewConstantTable() *ConstantTable {
  return &ConstantTable { make(map[string]*ConstantEntry) }
}

func (self *ConstantTable) Intern(s string) *ConstantEntry {
  ent, ok := self.table[s]
  if ! ok {
    ent = NewConstantEntry(s)
    self.table[s] = ent
  }
  return ent
}

func (self ConstantTable) GetEntries() []*ConstantEntry {
  entries := []*ConstantEntry { }
  for _, ent := range self.table {
    entries = append(entries, ent)
  }
  return entries
}

func (self *ConstantTable) IsEmpty() bool {
  return len(self.table) < 1
}
