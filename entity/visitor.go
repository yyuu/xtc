package entity

import (
  "bitbucket.org/yyuu/bs/core"
)

type IEntityVisitor interface {
  VisitEntity(core.IEntity)
}

func VisitEntity(v IEntityVisitor, entity core.IEntity) {
  v.VisitEntity(entity)
}
