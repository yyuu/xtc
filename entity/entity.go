package entity

type IEntity interface {
  IsEntity() bool
}

type IVariable interface {
  IEntity
  IsVariable() bool
}
