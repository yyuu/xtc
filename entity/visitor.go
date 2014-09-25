package entity

import (
  "bitbucket.org/yyuu/bs/core"
)

type IEntityVisitor interface {
  VisitEntity(core.IEntity) interface{}
}

func VisitEntity(v IEntityVisitor, entity core.IEntity) interface{} {
  return v.VisitEntity(entity)
}
