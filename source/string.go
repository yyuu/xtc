package source

type String struct {
  ext string
  readBuffer []byte
  writeBuffer []byte
}

func NewString(ext string, buf []byte) *String {
  return &String { ext, buf, []byte { } }
}

func ProgramSourceString(buf []byte) *String {
  return NewString(EXT_PROGRAM_SOURCE, buf)
}

func (self String) GetName() string {
  return ""
}

func (self *String) IsProgramSource() bool {
  return self.ext == EXT_PROGRAM_SOURCE
}

func (self *String) IsAssemblySource() bool {
  return self.ext == EXT_ASSEMBLY_SOURCE
}

func (self *String) IsObjectFile() bool {
  return self.ext == EXT_OBJECT_FILE
}

func (self *String) IsStaticLibrary() bool {
  return self.ext == EXT_STATIC_LIBRARY
}

func (self *String) IsSharedLibrary() bool {
  return self.ext == EXT_SHARED_LIBRARY
}

func (self *String) IsExecutableFile() bool {
  return self.ext == EXT_EXECUTABLE_FILE
}

func (self *String) Read() (string, error) {
  return string(self.readBuffer), nil
}

func (self *String) Write(s string) error {
  self.writeBuffer = []byte(s)
  return nil
}

func (self *String) ToProgramSource() ISource {
  return NewString(EXT_PROGRAM_SOURCE, []byte { })
}

func (self *String) ToAssemblySource() ISource {
  return NewString(EXT_ASSEMBLY_SOURCE, []byte { })
}

func (self *String) ToObjectFile() ISource {
  return NewString(EXT_OBJECT_FILE, []byte { })
}

func (self *String) ToStaticLibrary() ISource {
  return NewString(EXT_STATIC_LIBRARY, []byte { })
}

func (self *String) ToSharedLibrary() ISource {
  return NewString(EXT_SHARED_LIBRARY, []byte { })
}

func (self *String) ToExecutableFile() ISource {
  return NewString(EXT_EXECUTABLE_FILE, []byte { })
}
