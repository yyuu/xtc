package ast

import (
  "encoding/json"
  "fmt"
  "bitbucket.org/yyuu/bs/duck"
)

type Declarations struct {
  defvars []duck.IDefinedVariable
  vardecls []duck.IUndefinedVariable
  defuns []duck.IDefinedFunction
  funcdecls []duck.IUndefinedFunction
  constants []duck.IConstant
  defstructs []StructNode
  defunions []UnionNode
  typedefs []TypedefNode
}

func NewDeclarations(defvars []duck.IDefinedVariable, vardecls []duck.IUndefinedVariable, defuns []duck.IDefinedFunction, funcdecls []duck.IUndefinedFunction, constants []duck.IConstant, defstructs []StructNode, defunions []UnionNode, typedefs []TypedefNode) Declarations {
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
    Defvars []duck.IDefinedVariable
    Vardecls []duck.IUndefinedVariable
    Defuns []duck.IDefinedFunction
    Funcdecls []duck.IUndefinedFunction
    Constants []duck.IConstant
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

func (self Declarations) AddDefvar(v duck.IDefinedVariable) Declarations {
  return NewDeclarations(append(self.defvars, v), self.vardecls, self.defuns, self.funcdecls, self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddVardecl(v duck.IUndefinedVariable) Declarations {
  return NewDeclarations(self.defvars, append(self.vardecls, v), self.defuns, self.funcdecls, self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddDefun(f duck.IDefinedFunction) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, append(self.defuns, f), self.funcdecls, self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddFuncdecl(f duck.IUndefinedFunction) Declarations {
  return NewDeclarations(self.defvars, self.vardecls, self.defuns, append(self.funcdecls, f), self.constants, self.defstructs, self.defunions, self.typedefs)
}

func (self Declarations) AddDefconst(c duck.IConstant) Declarations {
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

func (self Declarations) GetDefvars() []duck.IDefinedVariable {
  return self.defvars
}

func (self Declarations) GetVardecls() []duck.IUndefinedVariable {
  return self.vardecls
}

func (self Declarations) GetDefuns() []duck.IDefinedFunction {
  return self.defuns
}

func (self Declarations) GetFuncdecls() []duck.IUndefinedFunction {
  return self.funcdecls
}

func (self Declarations) GetConstants() []duck.IConstant {
  return self.constants
}

func (self Declarations) GetDefstructs() []StructNode {
  return self.defstructs
}

func (self Declarations) GetDefunions() []UnionNode {
  return self.defunions
}

func (self Declarations) GetTypedefs() []TypedefNode {
  return self.typedefs
}
