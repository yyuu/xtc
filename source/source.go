package source

type ISource interface {
  GetName() string
  IsProgramSource() bool
  IsAssemblySource() bool
  IsObjectFile() bool
  IsStaticLibrary() bool
  IsSharedLibrary() bool
  IsExecutableFile() bool
  Read() (string, error)
  Write(string) error
  ToProgramSource() ISource
  ToAssemblySource() ISource
  ToObjectFile() ISource
  ToStaticLibrary() ISource
  ToSharedLibrary() ISource
  ToExecutableFile() ISource
}
