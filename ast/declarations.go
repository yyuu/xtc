package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type Declarations struct {
  ClassName string
  Defvars []core.IDefinedVariable
  Vardecls []core.IUndefinedVariable
  Defuns []core.IDefinedFunction
  Funcdecls []core.IUndefinedFunction
  Constants []*entity.Constant
  Defstructs []core.IStructNode
  Defunions []core.IUnionNode
  Typedefs []core.ITypedefNode
}

func NewDeclarations(defvars []core.IDefinedVariable, vardecls []core.IUndefinedVariable, defuns []core.IDefinedFunction, funcdecls []core.IUndefinedFunction, constants []*entity.Constant, defstructs []core.IStructNode, defunions []core.IUnionNode, typedefs []core.ITypedefNode) *Declarations {
  return &Declarations { "ast.Declarations", defvars, vardecls, defuns, funcdecls, constants, defstructs, defunions, typedefs }
}

func (self Declarations) String() string {
  return fmt.Sprintf("(begin (define defvars %s)\n" +
                     "       (define vardecls %s)\n" +
                     "       (define defuns %s)\n" +
                     "       (define funcdecls %s)\n" +
                     "       (define constants %s)\n" +
                     "       (define defstructs %s)\n" +
                     "       (define defunions %s)\n" +
                     "       (define typedefs %s))", self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declarations) GetLocation() core.Location {
  panic("Declarations#GetLocation called")
}

func (self Declarations) AddDefvar(v core.IDefinedVariable) *Declarations {
  return NewDeclarations(append(self.Defvars, v), self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declarations) AddVardecl(v core.IUndefinedVariable) *Declarations {
  return NewDeclarations(self.Defvars, append(self.Vardecls, v), self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declarations) AddDefun(f core.IDefinedFunction) *Declarations {
  return NewDeclarations(self.Defvars, self.Vardecls, append(self.Defuns, f), self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declarations) AddFuncdecl(f core.IUndefinedFunction) *Declarations {
  return NewDeclarations(self.Defvars, self.Vardecls, self.Defuns, append(self.Funcdecls, f), self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declarations) AddDefconst(c *entity.Constant) *Declarations {
  return NewDeclarations(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, append(self.Constants, c), self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declarations) AddDefstruct(s core.IStructNode) *Declarations {
  return NewDeclarations(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, append(self.Defstructs, s), self.Defunions, self.Typedefs)
}

func (self Declarations) AddDefunion(u core.IUnionNode) *Declarations {
  return NewDeclarations(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, append(self.Defunions, u), self.Typedefs)
}

func (self Declarations) AddTypedef(t core.ITypedefNode) *Declarations {
  return NewDeclarations(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, append(self.Typedefs, t))
}

func (self Declarations) GetDefvars() []core.IDefinedVariable {
  return self.Defvars
}

func (self Declarations) GetVardecls() []core.IUndefinedVariable {
  return self.Vardecls
}

func (self Declarations) GetDefuns() []core.IDefinedFunction {
  return self.Defuns
}

func (self Declarations) GetFuncdecls() []core.IUndefinedFunction {
  return self.Funcdecls
}

func (self Declarations) GetConstants() []*entity.Constant {
  return self.Constants
}

func (self Declarations) GetDefstructs() []core.IStructNode {
  return self.Defstructs
}

func (self Declarations) GetDefunions() []core.IUnionNode {
  return self.Defunions
}

func (self Declarations) GetTypedefs() []core.ITypedefNode {
  return self.Typedefs
}
