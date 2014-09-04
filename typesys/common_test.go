package typesys

type location struct {
  sourceName string
  lineNumber int
  lineOffset int
}

func (self location) GetSourceName() string {
  return self.sourceName
}

func (self location) GetLineNumber() int {
  return self.lineNumber
}

func (self location) GetLineOffset() int {
  return self.lineOffset
}

func (self location) String() string {
  panic("location#String called")
}

func (self location) MarshalJSON() ([]byte, error) {
  panic("location#MarshalJSON called")
}
