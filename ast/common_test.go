package ast

func loc(lineNumber int, lineOffset int) Location {
  return Location { "", lineNumber, lineOffset }
}
