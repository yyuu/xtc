package duck

type IEntity interface {
  String() string
  MarshalJSON() ([]byte, error)
  IsEntity() bool
}

type IVariable interface {
  IEntity
  IsVariable() bool
}
