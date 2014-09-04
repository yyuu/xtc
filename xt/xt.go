package xt

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "os"
  "os/exec"
  "reflect"
  "testing"
)

func AssertEquals(t *testing.T, s string, got interface{}, expected interface{}) {
  if got != expected {
    t.Errorf("AssertEquals: %s: expected %v, got %v", s, expected, got)
    t.Fail()
  }
}

func AssertDeepEquals(t *testing.T, s string, got interface{}, expected interface{}) {
  if ! reflect.DeepEqual(got, expected) {
    t.Errorf("AssertDeepEquals: %s: expected %v, got %v", s, expected, got)
    t.Fail()
  }
}

func AssertNotEquals(t *testing.T, s string, got interface{}, expected interface{}) {
  if got == expected {
    t.Errorf("AssertNotEquals: %s: expected %v, got %v", s, expected, got)
    t.Fail()
  }
}

func AssertTrue(t *testing.T, s string, got bool) {
  if ! got {
    t.Errorf("AssertTrue: %s: expected true, got %v", s, got)
    t.Fail()
  }
}

func AssertFalse(t *testing.T, s string, got bool) {
  if got {
    t.Errorf("AssertFalse: %s: expected false, got %v", s, got)
    t.Fail()
  }
}

func AssertNil(t *testing.T, s string, got interface{}) {
  if got != nil {
    t.Errorf("AssertNil: %s: expected nil, got %v", s, got)
    t.Fail()
  }
}

func AssertNotNil(t *testing.T, s string, got interface{}) {
  if got == nil {
    t.Errorf("AssertNotNil: %s: expected not nil", s)
    t.Fail()
  }
}

func AssertStringEquals(t *testing.T, s string, got string, expected string) {
  if got != expected {
    t.Errorf("AssertStringEquals: %s: expected %q, got %q", s, expected, got)
    t.Fail()
  }
}

func AssertStringNotEquals(t *testing.T, s string, got string, expected string) {
  if got != expected {
    t.Errorf("AssertStringNotEquals: %s: expected %q, got %q", s, expected, got)
    t.Fail()
  }
}

func AssertStringEqualsDiff(t *testing.T, s string, got string, expected string) {
  if got != expected {
    t.Errorf("AssertStringEqualsDiff: %s:\n// got %q\n// diff\n%s", s, got, diffString(expected, got))
    t.Fail()
  }
}

func diffString(x string, y string) string {
  xFile, yFile := writeFile(x), writeFile(y)
  defer func() {
    os.Remove(xFile)
    os.Remove(yFile)
  }()
  cmd := exec.Command("diff", "-u", xFile, yFile)
  var out bytes.Buffer
  cmd.Stdout = &out
  cmd.Run()
  return out.String()
}

func writeFile(s string) string {
  tmpdir := os.Getenv("TMP")
  if tmpdir == "" {
    tmpdir = "/tmp"
  }
  fp, err := ioutil.TempFile(tmpdir, "tmp")
  if err != nil {
    panic(err)
  }
  fp.WriteString(s)
  fp.Close()
  return fp.Name()
}

func JSON(x interface{}) string {
  cs, err := json.MarshalIndent(x, "", "  ")
  if err != nil {
    panic(err)
  }
  return string(cs)
}
