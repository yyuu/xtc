package ast

import (
  "bitbucket.org/yyuu/bs/entity"
)

type Declarations struct {
  Defvars []entity.DefinedVariable
  Vardecls []entity.UndefinedVariable
  Defuns []entity.DefinedFunction
  Funcdecls []entity.UndefinedFunction
  Constants []entity.Constant
  Defstructs []StructNode
  Defunions []UnionNode
  Typedefs []TypedefNode
}

func NewDeclarations() Declarations {
  return Declarations {
    Defvars: []entity.DefinedVariable { },
    Vardecls: []entity.UndefinedVariable { },
    Defuns: []entity.DefinedFunction { },
    Funcdecls: []entity.UndefinedFunction { },
    Constants: []entity.Constant { },
    Defstructs: []StructNode { },
    Defunions: []UnionNode { },
    Typedefs: []TypedefNode { },
  }
}

func copyDeclarationsWithAddDefvar(original Declarations, v entity.DefinedVariable) Declarations {
  return Declarations {
    Defvars: append(original.Defvars, v),
    Vardecls: original.Vardecls,
    Defuns: original.Defuns,
    Funcdecls: original.Funcdecls,
    Constants: original.Constants,
    Defstructs: original.Defstructs,
    Defunions: original.Defunions,
    Typedefs: original.Typedefs,
  }
}

func copyDeclarationsWithAddVardecl(original Declarations, v entity.UndefinedVariable) Declarations {
  return Declarations {
    Defvars: original.Defvars,
    Vardecls: append(original.Vardecls, v),
    Defuns: original.Defuns,
    Funcdecls: original.Funcdecls,
    Constants: original.Constants,
    Defstructs: original.Defstructs,
    Defunions: original.Defunions,
    Typedefs: original.Typedefs,
  }
}

func copyDeclarationsWithAddDefun(original Declarations, f entity.DefinedFunction) Declarations {
  return Declarations {
    Defvars: original.Defvars,
    Vardecls: original.Vardecls,
    Defuns: append(original.Defuns, f),
    Funcdecls: original.Funcdecls,
    Constants: original.Constants,
    Defstructs: original.Defstructs,
    Defunions: original.Defunions,
    Typedefs: original.Typedefs,
  }
}

func copyDeclarationsWithAddFuncdecl(original Declarations, f entity.UndefinedFunction) Declarations {
  return Declarations {
    Defvars: original.Defvars,
    Vardecls: original.Vardecls,
    Defuns: original.Defuns,
    Funcdecls: append(original.Funcdecls, f),
    Constants: original.Constants,
    Defstructs: original.Defstructs,
    Defunions: original.Defunions,
    Typedefs: original.Typedefs,
  }
}

func copyDeclarationsWithAddDefconst(original Declarations, c entity.Constant) Declarations {
  return Declarations {
    Defvars: original.Defvars,
    Vardecls: original.Vardecls,
    Defuns: original.Defuns,
    Funcdecls: original.Funcdecls,
    Constants: append(original.Constants, c),
    Defstructs: original.Defstructs,
    Defunions: original.Defunions,
    Typedefs: original.Typedefs,
  }
}

func copyDeclarationsWithAddDefstruct(original Declarations, s StructNode) Declarations {
  return Declarations {
    Defvars: original.Defvars,
    Vardecls: original.Vardecls,
    Defuns: original.Defuns,
    Funcdecls: original.Funcdecls,
    Constants: original.Constants,
    Defstructs: append(original.Defstructs, s),
    Defunions: original.Defunions,
    Typedefs: original.Typedefs,
  }
}

func copyDeclarationsWithAddDefunion(original Declarations, u UnionNode) Declarations {
  return Declarations {
    Defvars: original.Defvars,
    Vardecls: original.Vardecls,
    Defuns: original.Defuns,
    Funcdecls: original.Funcdecls,
    Constants: original.Constants,
    Defstructs: original.Defstructs,
    Defunions: append(original.Defunions, u),
    Typedefs: original.Typedefs,
  }
}

func copyDeclarationsWithAddTypedef(original Declarations, t TypedefNode) Declarations {
  return Declarations {
    Defvars: original.Defvars,
    Vardecls: original.Vardecls,
    Defuns: original.Defuns,
    Funcdecls: original.Funcdecls,
    Constants: original.Constants,
    Defstructs: original.Defstructs,
    Defunions: original.Defunions,
    Typedefs: append(original.Typedefs, t),
  }
}

func (self Declarations) String() string {
  panic("Declarations#String called")
}

func (self Declarations) MarshalJSON() ([]byte, error) {
  panic("Declarations#MarshalJSON called")
}

func (self Declarations) GetLocation() Location {
  panic("Declarations#GetLocation called")
}

func (self Declarations) AddDefvar(v entity.DefinedVariable) Declarations {
  return copyDeclarationsWithAddDefvar(self, v)
}

func (self Declarations) AddVardecl(v entity.UndefinedVariable) Declarations {
  return copyDeclarationsWithAddVardecl(self, v)
}

func (self Declarations) AddDefun(f entity.DefinedFunction) Declarations {
  return copyDeclarationsWithAddDefun(self, f)
}

func (self Declarations) AddFuncdecl(f entity.UndefinedFunction) Declarations {
  return copyDeclarationsWithAddFuncdecl(self, f)
}

func (self Declarations) AddDefconst(c entity.Constant) Declarations {
  return copyDeclarationsWithAddDefconst(self, c)
}

func (self Declarations) AddDefstruct(s StructNode) Declarations {
  return copyDeclarationsWithAddDefstruct(self, s)
}

func (self Declarations) AddDefunion(u UnionNode) Declarations {
  return copyDeclarationsWithAddDefunion(self, u)
}

func (self Declarations) AddTypedef(t TypedefNode) Declarations {
  return copyDeclarationsWithAddTypedef(self, t)
}
