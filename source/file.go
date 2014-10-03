package source

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

type File struct {
  name string
}

func NewFile(name string) *File {
  return &File { name }
}

func ProgramSourceFile(name string) *File {
  return NewFile(name)
}

func (self File) GetName() string {
  return self.name
}

func (self File) ext() string {
  dot := strings.LastIndex(self.name, ".")
  if dot == -1 {
    return ""
  } else {
    return self.name[dot:len(self.name)]
  }
}

func (self *File) IsProgramSource() bool {
  return self.ext() == EXT_PROGRAM_SOURCE
}

func (self *File) IsAssemblySource() bool {
  return self.ext() == EXT_ASSEMBLY_SOURCE
}

func (self *File) IsObjectFile() bool {
  return self.ext() == EXT_OBJECT_FILE
}

func (self *File) IsStaticLibrary() bool {
  return self.ext() == EXT_STATIC_LIBRARY
}

func (self *File) IsSharedLibrary() bool {
  return self.ext() == EXT_SHARED_LIBRARY
}

func (self *File) IsExecutableFile() bool {
  return self.ext() == EXT_EXECUTABLE_FILE
}

func (self *File) Read() (string, error) {
  bytes, err := ioutil.ReadFile(self.name)
  if err != nil {
    return "", err
  }
  return string(bytes), nil
}

func (self *File) Write(s string) error {
  return ioutil.WriteFile(self.name, []byte(s), 0644) // TODO: support umask
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

func (self *File) ToProgramSource() ISource {
  return NewFile(filename(self.name, EXT_PROGRAM_SOURCE))
}

func (self *File) ToAssemblySource() ISource {
  return NewFile(filename(self.name, EXT_ASSEMBLY_SOURCE))
}

func (self *File) ToObjectFile() ISource {
  return NewFile(filename(self.name, EXT_OBJECT_FILE))
}

func (self *File) ToStaticLibrary() ISource {
  return NewFile(filename(self.name, EXT_STATIC_LIBRARY))
}

func (self *File) ToSharedLibrary() ISource {
  return NewFile(filename(self.name, EXT_SHARED_LIBRARY))
}

func (self *File) ToExecutableFile() ISource {
  return NewFile(filename(self.name, EXT_EXECUTABLE_FILE))
}
