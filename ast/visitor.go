package ast

type IVisitor interface {
  Visit(interface{})
}

func Visit(v IVisitor, unknown interface{}) {
  v.Visit(unknown)
}
