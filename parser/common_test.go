package parser

import (
  "bitbucket.org/yyuu/bs/duck"
)

func loc(lineNumber int, lineOffset int) duck.Location {
  return duck.NewLocation("", lineNumber, lineOffset)
}
