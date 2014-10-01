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

func VisitEntities(v IEntityVisitor, entities []core.IEntity) interface{} {
  var x interface{}
  for i := range entities {
    x = VisitEntity(v, entities[i])
  }
  return x
}
