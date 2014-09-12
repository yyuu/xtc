package typesys

import (
  "encoding/json"
  "bitbucket.org/yyuu/bs/core"
)

type typeNode struct {
  Location core.Location
  TypeRef core.ITypeRef
  Type core.IType
}

func newTypeNode(loc core.Location, ref core.ITypeRef) *typeNode {
  return &typeNode { loc, ref, nil }
}

func (self typeNode) String() string {
  cs, _ := json.Marshal(self)
  return string(cs)
}

func (self typeNode) GetLocation() core.Location {
  return self.Location
}

func (self typeNode) IsTypeNode() bool {
  return true
}

func (self typeNode) GetTypeRef() core.ITypeRef {
  return self.TypeRef
}

func (self typeNode) IsResolved() bool {
  return self.Type != nil
}

func (self typeNode) GetType() core.IType {
  return self.Type
}

func (self *typeNode) SetType(t core.IType) {
  self.Type = t
}
