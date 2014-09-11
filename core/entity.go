package core

type IEntity interface {
  String() string
  GetName() string
  IsDefined() bool
  IsPrivate() bool
  IsConstant() bool
  IsRefered() bool
}
