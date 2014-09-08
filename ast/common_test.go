package ast

import (
  "bitbucket.org/yyuu/bs/duck"
)

func loc(lineNumber int, lineOffset int) duck.ILocation {
  return NewLocation("", lineNumber, lineOffset)
}
