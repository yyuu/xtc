package ast

import (
  "bitbucket.org/yyuu/bs/duck"
)

type IVisitor interface {
  Visit(duck.INode)
}

func Visit(v IVisitor, ast AST) {
  decl := ast.GetDeclarations()

  visitDefvars(v, decl.GetDefvars())
  visitVardecls(v, decl.GetVardecls())
  visitDefuns(v, decl.GetDefuns())
  visitFuncdecls(v, decl.GetFuncdecls())
  visitConstants(v, decl.GetConstants())
  visitDefstructs(v, decl.GetDefstructs())
  visitDefunions(v, decl.GetDefunions())
  visitTypedefs(v, decl.GetTypedefs())
}

func visitDefvars(v IVisitor, xs []duck.IDefinedVariable) {
//for i := range xs {
//  xs[i]
//}
}

func visitVardecls(v IVisitor, xs []duck.IUndefinedVariable) {
//for i := range xs {
//  xs[i]
//}
}

func visitDefuns(v IVisitor, xs []duck.IDefinedFunction) {
  for i := range xs {
    v.Visit(xs[i].GetBody())
  }
}

func visitFuncdecls(v IVisitor, xs []duck.IUndefinedFunction) {
//for i := range xs {
//  xs[i]
//}
}

func visitConstants(v IVisitor, xs []duck.IConstant) {
//for i := range xs {
//  xs[i]
//}
}

func visitDefstructs(v IVisitor, xs []StructNode) {
//for i := range xs {
//  xs[i]
//}
}

func visitDefunions(v IVisitor, xs []UnionNode) {
//for i := range xs {
//  xs[i]
//}
}

func visitTypedefs(v IVisitor, xs []TypedefNode) {
//for i := range xs {
//  xs[i]
//}
}
