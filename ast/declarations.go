package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/entity"
)

type Declarations struct {
  defvars []entity.DefinedVariable
  vardecls []entity.UndefinedVariable
  defuns []entity.DefinedFunction
  funcdecls []entity.UndefinedFunction
  constants []entity.Constant
  defstructs []StructNode
  defunions []UnionNode
  typedefs []TypedefNode
}

func NewDeclarations() Declarations {
  return Declarations {
    defvars: []entity.DefinedVariable { },
    vardecls: []entity.UndefinedVariable { },
    defuns: []entity.DefinedFunction { },
    funcdecls: []entity.UndefinedFunction { },
    constants: []entity.Constant { },
    defstructs: []StructNode { },
    defunions: []UnionNode { },
    typedefs: []TypedefNode { },
  }
}

func copyDeclarationsWithAddDefvar(original Declarations, v entity.DefinedVariable) Declarations {
  return Declarations {
    defvars: append(original.defvars, v),
    vardecls: original.vardecls,
    defuns: original.defuns,
    funcdecls: original.funcdecls,
    constants: original.constants,
    defstructs: original.defstructs,
    defunions: original.defunions,
    typedefs: original.typedefs,
  }
}

func copyDeclarationsWithAddVardecl(original Declarations, v entity.UndefinedVariable) Declarations {
  return Declarations {
    defvars: original.defvars,
    vardecls: append(original.vardecls, v),
    defuns: original.defuns,
    funcdecls: original.funcdecls,
    constants: original.constants,
    defstructs: original.defstructs,
    defunions: original.defunions,
    typedefs: original.typedefs,
  }
}

func copyDeclarationsWithAddDefun(original Declarations, f entity.DefinedFunction) Declarations {
  return Declarations {
    defvars: original.defvars,
    vardecls: original.vardecls,
    defuns: append(original.defuns, f),
    funcdecls: original.funcdecls,
    constants: original.constants,
    defstructs: original.defstructs,
    defunions: original.defunions,
    typedefs: original.typedefs,
  }
}

func copyDeclarationsWithAddFuncdecl(original Declarations, f entity.UndefinedFunction) Declarations {
  return Declarations {
    defvars: original.defvars,
    vardecls: original.vardecls,
    defuns: original.defuns,
    funcdecls: append(original.funcdecls, f),
    constants: original.constants,
    defstructs: original.defstructs,
    defunions: original.defunions,
    typedefs: original.typedefs,
  }
}

func copyDeclarationsWithAddDefconst(original Declarations, c entity.Constant) Declarations {
  return Declarations {
    defvars: original.defvars,
    vardecls: original.vardecls,
    defuns: original.defuns,
    funcdecls: original.funcdecls,
    constants: append(original.constants, c),
    defstructs: original.defstructs,
    defunions: original.defunions,
    typedefs: original.typedefs,
  }
}

func copyDeclarationsWithAddDefstruct(original Declarations, s StructNode) Declarations {
  return Declarations {
    defvars: original.defvars,
    vardecls: original.vardecls,
    defuns: original.defuns,
    funcdecls: original.funcdecls,
    constants: original.constants,
    defstructs: append(original.defstructs, s),
    defunions: original.defunions,
    typedefs: original.typedefs,
  }
}

func copyDeclarationsWithAddDefunion(original Declarations, u UnionNode) Declarations {
  return Declarations {
    defvars: original.defvars,
    vardecls: original.vardecls,
    defuns: original.defuns,
    funcdecls: original.funcdecls,
    constants: original.constants,
    defstructs: original.defstructs,
    defunions: append(original.defunions, u),
    typedefs: original.typedefs,
  }
}

func copyDeclarationsWithAddTypedef(original Declarations, t TypedefNode) Declarations {
  return Declarations {
    defvars: original.defvars,
    vardecls: original.vardecls,
    defuns: original.defuns,
    funcdecls: original.funcdecls,
    constants: original.constants,
    defstructs: original.defstructs,
    defunions: original.defunions,
    typedefs: append(original.typedefs, t),
  }
}

func (self Declarations) String() string {
  return fmt.Sprintf("(begin (define defvars %s)\n" +
                     "       (define vardecls %s)\n" +
                     "       (define defuns %s)\n" +
                     "       (define funcdecls %s)\n" +
                     "       (define constants %s)\n" +
                     "       (define defstructs %s)\n" +
                     "       (define defunions %s)\n" +
                     "       (define typedefs %s))", self.defvars, self.vardecls, self.defuns, self.funcdecls, self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) MarshalJSON() ([]byte, error) {
  var x struct {
    ClassName string
    Defvars []entity.DefinedVariable
    Vardecls []entity.UndefinedVariable
    Defuns []entity.DefinedFunction
    Funcdecls []entity.UndefinedFunction
    Constants []entity.Constant
    Defstructs []StructNode
    Defunions []UnionNode
    Typedefs []TypedefNode
  }
  x.ClassName = "ast.Declarations"
  x.Defvars = self.defvars
  x.Vardecls = self.vardecls
  x.Defuns = self.defuns
  x.Funcdecls = self.funcdecls
  x.Constants = self.constants
  x.Defstructs = self.defstructs
  x.Defunions = self.defunions
  x.Typedefs = self.typedefs
  return json.Marshal(x)
}

func (self Declarations) GetLocation() duck.ILocation {
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
