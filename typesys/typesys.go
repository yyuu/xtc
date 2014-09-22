package typesys

import (
  "bitbucket.org/yyuu/bs/core"
)

func NewTypes(xs...core.IType) []core.IType {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IType { }
  }
}

func NewTypeRefs(xs...core.ITypeRef) []core.ITypeRef {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.ITypeRef { }
  }
}
