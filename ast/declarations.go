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

func NewDeclarations(defvars []entity.DefinedVariable, vardecls []entity.UndefinedVariable, defuns []entity.DefinedFunction, funcdecls []entity.UndefinedFunction, constants []entity.Constant, defstructs []StructNode, defunions []UnionNode, typedefs []TypedefNode) Declarations {
  return Declarations { defvars, vardecls, defuns, funcdecls, constants, defstructs, defunions, typedefs }
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
  return NewDeclarations(append(self.defvars, v), self.vardecls, self.defuns, self.funcdecls, self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddVardecl(v entity.UndefinedVariable) Declarations {
  return NewDeclarations(self.defvars, append(self.vardecls, v), self.defuns, self.funcdecls, self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddDefun(f entity.DefinedFunction) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, append(self.defuns, f), self.funcdecls, self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddFuncdecl(f entity.UndefinedFunction) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, self.defuns, append(self.funcdecls, f), self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddDefconst(c entity.Constant) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, self.defuns, self.funcdecls, append(self.constants, c), self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddDefstruct(s StructNode) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, self.defuns, self.funcdecls, self.constants, append(self.defstructs, s), self.defunions, self.typedefs)
}

func (self Declarations) AddDefunion(u UnionNode) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, self.defuns, self.funcdecls, self.constants, self.defstructs, append(self.defunions, u), self.typedefs)
}

func (self Declarations) AddTypedef(t TypedefNode) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, self.defuns, self.funcdecls, self.constants, self.defstructs, self.defunions, append(self.typedefs, t))
}
