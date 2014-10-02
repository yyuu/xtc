package parser

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
)

type token struct {
  id int
  literal string
  location core.Location
}

func (self token) String() string {
  return fmt.Sprintf("#<token:%d %s %q>", self.id, self.location, self.literal)
}
