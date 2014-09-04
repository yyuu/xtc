package parser

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
)

func loc(lineNumber int, lineOffset int) duck.ILocation {
  return ast.NewLocation("", lineNumber, lineOffset)
}
