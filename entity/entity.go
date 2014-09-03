package entity

import (
  "bitbucket.org/yyuu/bs/typesys"
)

type IEntity interface {
  IsEntity() bool
}

type IVariable interface {
  IEntity
  IsVariable() bool
}

type ILocation interface {
  GetSourceName() string
  GetLineNumber() int
  GetLineOffset() int
}

type IExprNode interface {
  IsExpr() bool
//GetLocation() ILocation
}

type IStmtNode interface {
  IsStmt() bool
}

type ITypeNode interface {
  IsType() bool
  GetTypeRef() typesys.ITypeRef
//GetLocation() ILocation
}
