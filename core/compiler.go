package core

import (
  "io/ioutil"
  "os"
  "path"
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
  dir string
  base string
  ext string
  generated bool
}

func NewSourceFile(name, file, ext string) *SourceFile {
  if ext == "" {
    ext = detectExt(name)
  }
  dir := path.Dir(file)
  base := path.Base(file)
  n := strings.LastIndex(base, ext)
  return &SourceFile { name, dir, base[0:n], ext, false }
}

func NewTemporarySourceFile(name, ext string, bytes []byte) (*SourceFile, error) {
  if ext == "" {
    ext = detectExt(name)
  }
  // FIXME: create tempfile safely
  dir, err := ioutil.TempDir("/tmp", "xtc.")
  if err != nil {
    return nil, err
  }
  base := "tmp"
  file := path.Join(dir, base+ext)
  err = ioutil.WriteFile(file, bytes, 0644) // TODO: support umask
  if err != nil {
    return nil, err
  }
  return &SourceFile { name, dir, base, ext, true }, nil
}

func detectExt(s string) string {
  n := strings.LastIndex(s, ".")
  if n == -1 {
    return ""
  } else {
    return s[n:len(s)]
  }
}

func (self SourceFile) String() string {
  return self.GetPath()
}

func (self SourceFile) GetName() string {
  return self.name
}

func (self SourceFile) GetPath() string {
  return path.Join(self.dir, self.base+self.ext)
}

func (self SourceFile) ReadAll() ([]byte, error) {
  bytes, err := ioutil.ReadFile(self.GetPath())
  if err != nil {
    return []byte { }, err
  }
  return bytes, nil
}

func (self SourceFile) WriteAll(bytes []byte) error {
  return ioutil.WriteFile(self.GetPath(), bytes, 0644) // TODO: support umask
}

func (self SourceFile) Remove() error {
  return os.Remove(self.GetPath())
}

func (self SourceFile) IsGenerated() bool {
  return self.generated
}

func (self SourceFile) IsProgramSource() bool {
  return self.ext == EXT_PROGRAM_SOURCE
}

func (self SourceFile) IsAssemblySource() bool {
  return self.ext == EXT_ASSEMBLY_SOURCE
}

func (self SourceFile) IsObjectFile() bool {
  return self.ext == EXT_OBJECT_FILE
}

func (self SourceFile) IsStaticLibrary() bool {
  return self.ext == EXT_STATIC_LIBRARY
}

func (self SourceFile) IsSharedLibrary() bool {
  return self.ext == EXT_SHARED_LIBRARY
}

func (self SourceFile) IsExecutableFile() bool {
  return self.ext == EXT_EXECUTABLE_FILE
}

func (self SourceFile) newName(ext string) string {
  if self.name == "" {
    return ""
  } else {
    n := strings.LastIndex(self.name, self.ext)
    return self.name[0:n] + ext
  }
}

func (self SourceFile) ToProgramSource() *SourceFile {
  return &SourceFile { self.newName(EXT_PROGRAM_SOURCE), self.dir, self.base, EXT_PROGRAM_SOURCE, true }
}

func (self SourceFile) ToAssemblySource() *SourceFile {
  return &SourceFile { self.newName(EXT_ASSEMBLY_SOURCE), self.dir, self.base, EXT_ASSEMBLY_SOURCE, true }
}

func (self SourceFile) ToObjectFile() *SourceFile {
  return &SourceFile { self.newName(EXT_OBJECT_FILE), self.dir, self.base, EXT_OBJECT_FILE, true }
}

func (self SourceFile) ToStaticLibrary() *SourceFile {
  return &SourceFile { self.newName(EXT_STATIC_LIBRARY), self.dir, self.base, EXT_STATIC_LIBRARY, true }
}

func (self SourceFile) ToSharedLibrary() *SourceFile {
  return &SourceFile { self.newName(EXT_SHARED_LIBRARY), self.dir, self.base, EXT_SHARED_LIBRARY, true }
}

func (self SourceFile) ToExecutableFile() *SourceFile {
  return &SourceFile { self.newName(EXT_EXECUTABLE_FILE), self.dir, self.base, EXT_EXECUTABLE_FILE, true }
}
