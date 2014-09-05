package parser

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
)

func loc(lineNumber int, lineOffset int) duck.ILocation {
  return ast.NewLocation("", lineNumber, lineOffset)
}

func defvars(xs...duck.IDefinedVariable) []duck.IDefinedVariable {
  if len(xs) == 0 {
    return []duck.IDefinedVariable { }
  } else {
    return xs
  }
}

func vardecls(xs...duck.IUndefinedVariable) []duck.IUndefinedVariable {
  if len(xs) == 0 {
    return []duck.IUndefinedVariable { }
  } else {
    return xs
  }
}

func defuns(xs...duck.IDefinedFunction) []duck.IDefinedFunction {
  if len(xs) == 0 {
    return []duck.IDefinedFunction { }
  } else {
    return xs
  }
}

func funcdecls(xs...duck.IUndefinedFunction) []duck.IUndefinedFunction {
  if len(xs) == 0 {
    return []duck.IUndefinedFunction { }
  } else {
    return xs
  }
}

func defconsts(xs...duck.IConstant) []duck.IConstant {
  if len(xs) == 0 {
    return []duck.IConstant { }
  } else {
    return xs
  }
}

func defstructs(xs...ast.StructNode) []ast.StructNode {
  if len(xs) == 0 {
    return []ast.StructNode { }
  } else {
    return xs
  }
}

func defunions(xs...ast.UnionNode) []ast.UnionNode {
  if len(xs) == 0 {
    return []ast.UnionNode { }
  } else {
    return xs
  }
}

func typedefs(xs...ast.TypedefNode) []ast.TypedefNode {
  if len(xs) == 0 {
    return []ast.TypedefNode { }
  } else {
    return xs
  }
}
