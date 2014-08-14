package strscan

import (
  "testing"
)

func TestCheck(t *testing.T) {
  s := NewStringScanner("Fri Dec 12 1975 14:39")
  var matched string

  matched = s.Check("Fri")
  if matched != "Fri" {
    t.Errorf("expected %q, found %q", "Fri", matched)
  }
  if s.Pos() != 0 {
    t.Fail()
  }
  matched = s.Check("12")
  if matched != "" {
    t.Errorf("expected %q, found %q", "", matched)
  }
}

func TestCheckUntil(t *testing.T) {
  s := NewStringScanner("Fri Dec 12 1975 14:39")
  var matched string

  matched = s.CheckUntil("12")
  if matched != "Fri Dec 12" {
    t.Errorf("expected %q, found %q", "Fri Dec 12", matched)
  }
  if s.Pos() != 0 {
    t.Fail()
  }
}

func TestIsEOS(t *testing.T) {
  s := NewStringScanner("")
  if !s.IsEOS() {
    t.Errorf("should be EOS: peek(1) == %q", s.Peek(1))
  }
}

func TestIsEOS2(t *testing.T) {
  s := NewStringScanner("abcdef")
  if s.IsEOS() {
    t.Errorf("should not be EOS: peek(1) == %q", s.Peek(1))
  }

  s.Seek(6)
  if !s.IsEOS() {
    t.Errorf("should be EOS: peek(1) == %q", s.Peek(1))
  }

  s.Seek(0)
  if s.IsEOS() {
    t.Errorf("should not be EOS: peek(1) == %q", s.Peek(1))
  }
}

func TestMatch(t *testing.T) {
  s := NewStringScanner("test string")
  if s.Match("[A-Za-z]+") != 4 {
    t.Fail()
  }
  if s.Match("[A-Za-z]+") != 4 {
    t.Fail()
  }
  if s.Match(" +") != 0 {
    t.Fail()
  }
}

func TestPeek(t *testing.T) {
  s := NewStringScanner("test string")
  if s.Peek(7) != "test st" {
    t.Fail()
  }
  if s.Peek(7) != "test st" {
    t.Fail()
  }
}

func TestScan(t *testing.T) {
  s := NewStringScanner("test string")
  if s.Scan("[A-Za-z]+") != "test" {
    t.Fail()
  }
  if s.Scan("[A-Za-z]+") != "" {
    t.Fail()
  }
  if s.Scan(" +") != " " {
    t.Fail()
  }
  if s.Scan("[A-Za-z]+") != "string" {
    t.Fail()
  }
  if s.Scan(".") != "" {
    t.Fail()
  }
}

func TestScanUntil(t *testing.T) {
  s := NewStringScanner("Fri Dec 12 1975 14:39")
  if s.ScanUntil("1") != "Fri Dec 1" {
    t.Fail()
  }
  if s.ScanUntil("XYZ") != "" {
    t.Fail()
  }
}

func TestSkip(t *testing.T) {
  s := NewStringScanner("test string")
  if s.Skip("[A-Za-z]+") != 4 {
    t.Fail()
  }
  if s.Skip("[A-Za-z]+") != 0 {
    t.Fail()
  }
  if s.Skip(" +") != 1 {
    t.Fail()
  }
  if s.Skip("[A-Za-z]+") != 6 {
    t.Fail()
  }
  if s.Skip(".") != 0 {
    t.Fail()
  }
}

func TestSkipUntil(t *testing.T) {
}
