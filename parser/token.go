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
  return fmt.Sprintf("%s %d %q", self.location, self.id, self.literal)
}
