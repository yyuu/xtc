package entity

type DefinedVariable struct {
  Private bool
  Name string
  TypeNode ITypeNode
  NumRefered int
  Initializer IExprNode
}

func NewDefinedVariable(isPrivate bool, t ITypeNode, name string, init IExprNode) DefinedVariable {
  return DefinedVariable {
    Private: isPrivate,
    Name: name,
    TypeNode: t,
    NumRefered: 0,
    Initializer: init,
  }
}

func (self DefinedVariable) IsEntity() bool {
  return true
}

func (self DefinedVariable) IsDefined() bool {
  return true
}

func (self DefinedVariable) HasInitializer() bool {
  return self.Initializer != nil
}

type UndefinedVariable struct {
  Private bool
  Name string
  TypeNode ITypeNode
}

func NewUndefinedVariable(t ITypeNode, name string) UndefinedVariable {
  return UndefinedVariable {
    Private: false,
    Name: name,
    TypeNode: t,
  }
}

func (self UndefinedVariable) IsEntity() bool {
  return true
}
