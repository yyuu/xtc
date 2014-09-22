package ast

import (
  "fmt"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
)

type Declaration struct {
  ClassName string
  Defvars []*entity.DefinedVariable
  Vardecls []*entity.UndefinedVariable
  Defuns []*entity.DefinedFunction
  Funcdecls []*entity.UndefinedFunction
  Constants []*entity.Constant
  Defstructs []*StructNode
  Defunions []*UnionNode
  Typedefs []*TypedefNode
}

func NewDeclaration(defvars []*entity.DefinedVariable, vardecls []*entity.UndefinedVariable, defuns []*entity.DefinedFunction, funcdecls []*entity.UndefinedFunction, constants []*entity.Constant, defstructs []*StructNode, defunions []*UnionNode, typedefs []*TypedefNode) *Declaration {
  return &Declaration { "ast.Declaration", defvars, vardecls, defuns, funcdecls, constants, defstructs, defunions, typedefs }
}

func (self Declaration) String() string {
  return fmt.Sprintf("(begin (define defvars %s)\n" +
                     "       (define vardecls %s)\n" +
                     "       (define defuns %s)\n" +
                     "       (define funcdecls %s)\n" +
                     "       (define constants %s)\n" +
                     "       (define defstructs %s)\n" +
                     "       (define defunions %s)\n" +
                     "       (define typedefs %s))", self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declaration) GetLocation() core.Location {
  panic("Declaration#GetLocation called")
}

func (self Declaration) AddDefvar(v *entity.DefinedVariable) *Declaration {
  return NewDeclaration(append(self.Defvars, v), self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declaration) AddVardecl(v *entity.UndefinedVariable) *Declaration {
  return NewDeclaration(self.Defvars, append(self.Vardecls, v), self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declaration) AddDefun(f *entity.DefinedFunction) *Declaration {
  return NewDeclaration(self.Defvars, self.Vardecls, append(self.Defuns, f), self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declaration) AddFuncdecl(f *entity.UndefinedFunction) *Declaration {
  return NewDeclaration(self.Defvars, self.Vardecls, self.Defuns, append(self.Funcdecls, f), self.Constants, self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declaration) AddDefconst(c *entity.Constant) *Declaration {
  return NewDeclaration(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, append(self.Constants, c), self.Defstructs, self.Defunions, self.Typedefs)
}

func (self Declaration) AddDefstruct(s *StructNode) *Declaration {
  return NewDeclaration(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, append(self.Defstructs, s), self.Defunions, self.Typedefs)
}

func (self Declaration) AddDefunion(u *UnionNode) *Declaration {
  return NewDeclaration(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, append(self.Defunions, u), self.Typedefs)
}

func (self Declaration) AddTypedef(t *TypedefNode) *Declaration {
  return NewDeclaration(self.Defvars, self.Vardecls, self.Defuns, self.Funcdecls, self.Constants, self.Defstructs, self.Defunions, append(self.Typedefs, t))
}

func (self Declaration) GetDefvars() []*entity.DefinedVariable {
  return self.Defvars
}

func (self Declaration) GetVardecls() []*entity.UndefinedVariable {
  return self.Vardecls
}

func (self Declaration) GetDefuns() []*entity.DefinedFunction {
  return self.Defuns
}

func (self Declaration) GetFuncdecls() []*entity.UndefinedFunction {
  return self.Funcdecls
}

func (self Declaration) GetConstants() []*entity.Constant {
  return self.Constants
}

func (self Declaration) GetDefstructs() []*StructNode {
  return self.Defstructs
}

func (self Declaration) GetDefunions() []*UnionNode {
  return self.Defunions
}

func (self Declaration) GetTypedefs() []*TypedefNode {
  return self.Typedefs
}
