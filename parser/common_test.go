package parser

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/entity"
)

func loc(lineNumber int, lineOffset int) duck.ILocation {
  return ast.NewLocation("", lineNumber, lineOffset)
}

func defvars(xs...entity.DefinedVariable) []entity.DefinedVariable {
  return xs
}

func vardecls(xs...entity.UndefinedVariable) []entity.UndefinedVariable {
  return xs
}

func defuns(xs...entity.DefinedFunction) []entity.DefinedFunction {
  return xs
}

func funcdecls(xs...entity.UndefinedFunction) []entity.UndefinedFunction {
  return xs
}

func defconsts(xs...entity.Constant) []entity.Constant {
  return xs
}

func defstructs(xs...ast.StructNode) []ast.StructNode {
  return xs
}

func defunions(xs...ast.UnionNode) []ast.UnionNode {
  return xs
}

func typedefs(xs...ast.TypedefNode) []ast.TypedefNode {
  return xs
}
