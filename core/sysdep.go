package core

const (
  PLATFORM_X86_LINUX = iota
  PLATFORM_AMD64_LINUX
)

type IAssemblyCode interface {
  ToSource() string
}
