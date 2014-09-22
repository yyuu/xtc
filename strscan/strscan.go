package strscan

import (
  "regexp"
  "unicode/utf8"
)

type StringScanner struct {
  String string
  offset int
  lineNumber int
  lineOffset int
  matched string
}

func New(s string) *StringScanner {
  return &StringScanner {
    String: s, 
    offset: 0,
    lineNumber: 0,
    lineOffset: 0,
    matched: "",
  }
}

func (self *StringScanner) Check(pattern string) string {
  // This returns the value that scan would return, without advancing the scan pointer. The match register is affected, though.
  re := mustCompile(pattern)
  if !self.IsEOS() {
    return re.FindString(self.String[self.offset:])
  }
  return ""
}

func (self *StringScanner) CheckUntil(pattern string) string {
  offset := self.offset
  re := mustCompile(pattern)
  for !self.IsEOS() {
    matched := re.FindString(self.String[offset:])
    if matched != "" {
      self.matched = matched
      return self.Peek(offset - self.offset + len(matched))
    }
    _, size := utf8.DecodeRuneInString(self.String[offset:])
    offset += size
  }
  return ""
}

func (self *StringScanner) IsEOS() bool {
  return len(self.String) <= self.offset
}

func (self *StringScanner) Match(pattern string) int {
  // Tests whether the given pattern is matched from the current scan pointer. Returns the length of the match, or nil. The scan pointer is not advanced.
  re := mustCompile(pattern)
  if !self.IsEOS() {
    matched := re.FindString(self.String[self.offset:])
    if matched != "" {
      self.matched = matched
      return len(matched)
    }
  }
  return 0
}

func (self *StringScanner) Matched() string {
  return self.matched
}

func (self *StringScanner) Peek(length int) string {
  // Extracts a string corresponding to string[pos,len], without advancing the scan pointer.
  var s string
  for i, offset := 0, self.offset; i < length && offset < len(self.String); i++ {
    r, size := utf8.DecodeRuneInString(self.String[offset:])
    offset += size
    p := make([]byte, size)
    utf8.EncodeRune(p, r)
    s += string(p)
  }
  return s
}

func (self *StringScanner) Pos() int {
  return self.offset
}

func (self *StringScanner) Seek(offset int) int {
  self.offset = offset
  return offset
}

func (self *StringScanner) Scan(pattern string) string {
  // Tries to match with pattern at the current position. If there’s a match, the scanner advances the “scan pointer” and returns the matched string. Otherwise, the scanner returns nil.
  re := mustCompile(pattern)
  if !self.IsEOS() {
    matched := re.FindString(self.String[self.offset:])
    if matched != "" {
      self.offset += len(matched)
      self.matched = matched
      return matched
    }
  }
  return ""
}

func (self *StringScanner) ScanUntil(pattern string) string {
  offset := self.offset
  re := mustCompile(pattern)
  for !self.IsEOS() {
    matched := re.FindString(self.String[self.offset:])
    if matched != "" {
      self.offset += len(matched)
      self.matched = self.String[offset:self.offset]
      return self.matched
    }
    self.skipRune()
  }
  return ""
}

func (self *StringScanner) Skip(pattern string) int {
  // Attempts to skip over the given pattern beginning with the scan pointer. If it matches, the scan pointer is advanced to the end of the match, and the length of the match is returned. Otherwise, nil is returned.
  re := mustCompile(pattern)
  if !self.IsEOS() {
    matched := re.FindString(self.String[self.offset:])
    if matched != "" {
      self.offset += len(matched)
      self.matched = matched
      return len(matched)
    }
  }
  return 0
}

func (self *StringScanner) SkipUntil(pattern string) int {
  // Advances the scan pointer until pattern is matched and consumed. Returns the number of bytes advanced, or nil if no match was found.
  //
  // Look ahead to match pattern, and advance the scan pointer to the end of the match. Return the number of characters advanced, or nil if the match was unsuccessful.
  //
  // It’s similar to scan_until, but without returning the intervening string.
  re := mustCompile(pattern)
  for !self.IsEOS() {
    matched := re.FindString(self.String[self.offset:])
    if matched != "" {
      self.offset += len(matched)
      self.matched = matched
      return len(matched)
    }
    self.skipRune()
  }
  return 0
}

func mustCompile(pattern string) *regexp.Regexp {
  // FIXME: bad naming
  return regexp.MustCompile("\\A" + pattern)
}

func (self *StringScanner) skipRune() {
  if !self.IsEOS() {
    _, size := utf8.DecodeRuneInString(self.String[self.offset:])
    self.offset += size
  }
}
