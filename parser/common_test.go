package parser

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "os"
  "os/exec"
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
)

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

func loc(lineNumber int, lineOffset int) duck.ILocation {
  return ast.NewLocation("", lineNumber, lineOffset)
}

func jsonString(x interface{}) string {
  cs, err := json.MarshalIndent(x, "", "  ")
  if err != nil {
    panic(err)
  }
  return string(cs)
}
