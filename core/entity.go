package core

type IEntity interface {
  String() string
  IsEntity() bool
  GetName() string
  IsDefined() bool
  IsPrivate() bool
  IsConstant() bool
  IsRefered() bool
}
