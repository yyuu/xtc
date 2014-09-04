package ast

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "os"
  "os/exec"
  "testing"
  "bitbucket.org/yyuu/bs/duck"
)

func assertJsonEquals(t *testing.T, got duck.INode, expected string) {
  s := jsonString(got)
  if s != expected {
//  t.Errorf("\n// expected\n%s\n// got\n%s\n", expected, s)
    t.Errorf("\n// got\n%s\n// diff\n%s\n", s, diff(expected, s))
    t.Fail()
  }
}

func diff(x, y string) string {
  a, b := tempfile(x), tempfile(y)
  defer func() {
    os.Remove(a)
    os.Remove(b)
  }()

  cmd := exec.Command("diff", "-u", a, b)
  var out bytes.Buffer
  cmd.Stdout = &out
  cmd.Run()
  return out.String()
}

func tempfile(s string) string {
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

func loc(lineNumber int, lineOffset int) Location {
  return Location { "", lineNumber, lineOffset }
}

func jsonString(x interface{}) string {
  cs, err := json.MarshalIndent(x, "", "  ")
  if err != nil {
    panic(err)
  }
  return string(cs)
}
