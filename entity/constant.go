package entity

type Constant struct {
  Name string
  TypeNode ITypeNode
  Value IExprNode
}

func NewConstant(name string, t ITypeNode, value IExprNode) Constant {
  return Constant {
    Name: name,
    TypeNode: t,
    Value: value,
  }
}

func (self Constant) IsEntity() bool {
  return true
}
