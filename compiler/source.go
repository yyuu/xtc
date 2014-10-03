package compiler

import (
  "io/ioutil"
  "strings"
)

const (
  EXT_PROGRAM_SOURCE  = ".xtc"
  EXT_ASSEMBLY_SOURCE = ".s"
  EXT_OBJECT_FILE     = ".o"
  EXT_STATIC_LIBRARY  = ".a"
  EXT_SHARED_LIBRARY  = ".so"
  EXT_EXECUTABLE_FILE = ""
)

type SourceFile struct {
  name string
}

func NewSourceFile(name string) *SourceFile {
  return &SourceFile { name }
}

func (self SourceFile) GetName() string {
  return self.name
}

func (self SourceFile) ext() string {
  dot := strings.LastIndex(self.name, ".")
  if dot == -1 {
    return ""
  } else {
    return self.name[dot:len(self.name)]
  }
}

func (self *SourceFile) IsProgramSource() bool {
  return self.ext() == EXT_PROGRAM_SOURCE
}

func (self *SourceFile) IsAssemblySource() bool {
  return self.ext() == EXT_ASSEMBLY_SOURCE
}

func (self *SourceFile) IsObjectFile() bool {
  return self.ext() == EXT_OBJECT_FILE
}

func (self *SourceFile) IsStaticLibrary() bool {
  return self.ext() == EXT_STATIC_LIBRARY
}

func (self *SourceFile) IsSharedLibrary() bool {
  return self.ext() == EXT_SHARED_LIBRARY
}

func (self *SourceFile) IsExecutableFile() bool {
  return self.ext() == EXT_EXECUTABLE_FILE
}

func (self *SourceFile) Read() ([]byte, error) {
  bytes, err := ioutil.ReadFile(self.name)
  if err != nil {
    return []byte { }, err
  }
  return bytes, nil
}

func (self *SourceFile) Write(bytes []byte) error {
  return ioutil.WriteFile(self.name, bytes, 0644) // TODO: support umask
}

func filename(name, ext string) string {
  dot := strings.LastIndex(name, ".")
  var base string
  if dot == -1 {
    base = name
  } else {
    base = name[dot:len(name)]
  }
  return base + ext
}

func (self *SourceFile) ToProgramSource() *SourceFile {
  return NewSourceFile(filename(self.name, EXT_PROGRAM_SOURCE))
}

func (self *SourceFile) ToAssemblySource() *SourceFile {
  return NewSourceFile(filename(self.name, EXT_ASSEMBLY_SOURCE))
}

func (self *SourceFile) ToObjectFile() *SourceFile {
  return NewSourceFile(filename(self.name, EXT_OBJECT_FILE))
}

func (self *SourceFile) ToStaticLibrary() *SourceFile {
  return NewSourceFile(filename(self.name, EXT_STATIC_LIBRARY))
}

func (self *SourceFile) ToSharedLibrary() *SourceFile {
  return NewSourceFile(filename(self.name, EXT_SHARED_LIBRARY))
}

func (self *SourceFile) ToExecutableFile() *SourceFile {
  return NewSourceFile(filename(self.name, EXT_EXECUTABLE_FILE))
}
