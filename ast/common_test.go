package ast

import (
  "bitbucket.org/yyuu/bs/core"
)

func loc(lineNumber int, lineOffset int) core.Location {
  return core.NewLocation("", lineNumber, lineOffset)
}
