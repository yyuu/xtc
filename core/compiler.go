package core

import (
  "io/ioutil"
  "os"
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
  path string
  temporary bool
  anonymous bool
}

func NewSourceFile(path string) *SourceFile {
  return &SourceFile { path, false, false }
}

func NewTemporarySourceFile(ext string, bytes []byte) (*SourceFile, error) {
  // FIXME: create tempfile safely
  file, err := ioutil.TempFile("/tmp", "tmp")
  if err != nil {
    return nil, err
  }
  err = file.Close()
  if err != nil {
    return nil, err
  }
  basename := file.Name()
  err = os.Remove(basename)
  if err != nil {
    return nil, err
  }
  path := basename + ext
  err = ioutil.WriteFile(path, bytes, 0644) // TODO: support umask
  if err != nil {
    return nil, err
  }
  return &SourceFile { path, true, true }, nil
}

func (self SourceFile) GetName() string {
  if self.anonymous {
    return ""
  } else {
    return self.path
  }
}

func (self SourceFile) ext() string {
  dot := strings.LastIndex(self.path, ".")
  if dot == -1 {
    return ""
  } else {
    return self.path[dot:len(self.path)]
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

func (self *SourceFile) ReadAll() ([]byte, error) {
  bytes, err := ioutil.ReadFile(self.path)
  if err != nil {
    return []byte { }, err
  }
  return bytes, nil
}

func (self *SourceFile) WriteAll(bytes []byte) error {
  return ioutil.WriteFile(self.path, bytes, 0644) // TODO: support umask
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
  return &SourceFile { filename(self.path, EXT_PROGRAM_SOURCE), true, self.anonymous }
}

func (self *SourceFile) ToAssemblySource() *SourceFile {
  return &SourceFile { filename(self.path, EXT_ASSEMBLY_SOURCE), true, self.anonymous }
}

func (self *SourceFile) ToObjectFile() *SourceFile {
  return &SourceFile { filename(self.path, EXT_OBJECT_FILE), true, self.anonymous }
}

func (self *SourceFile) ToStaticLibrary() *SourceFile {
  return &SourceFile { filename(self.path, EXT_STATIC_LIBRARY), false, self.anonymous }
}

func (self *SourceFile) ToSharedLibrary() *SourceFile {
  return &SourceFile { filename(self.path, EXT_SHARED_LIBRARY), false, self.anonymous }
}

func (self *SourceFile) ToExecutableFile() *SourceFile {
  return &SourceFile { filename(self.path, EXT_EXECUTABLE_FILE), false, self.anonymous }
}

func (self *SourceFile) IsTemporary() bool {
  return self.temporary
}

func (self *SourceFile) Remove() error {
  return os.Remove(self.path)
}
