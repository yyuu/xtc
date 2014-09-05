package entity

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Constant struct {
  name string
  typeNode duck.ITypeNode
  value duck.IExprNode
}

func NewConstant(t duck.ITypeNode, name string, value duck.IExprNode) Constant {
  return Constant { name, t, value }
}

func (self Constant) String() string {
  return fmt.Sprintf("<entity.Constant Name=%s TypeNode=%s Value=%s>", self.name, self.typeNode, self.value)
}

func (self Constant) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Name string
    TypeNode duck.ITypeNode
    Value duck.IExprNode
  }
  x.ClassName = "entity.Constant"
  x.Name = self.name
  x.TypeNode = self.typeNode
  x.Value = self.value
  return json.Marshal(x)
}

func (self Constant) IsEntity() bool {
  return true
}

func (self Constant) IsConstant() bool {
  return true
}

func (self Constant) GetName() string {
  return self.name
}

func (self Constant) GetTypeNode() duck.ITypeNode {
  return self.typeNode
}

func (self Constant) GetValue() duck.IExprNode {
  return self.value
}
