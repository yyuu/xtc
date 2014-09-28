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

func AsDeclaration(x core.INode) *Declaration {
  return x.(*Declaration)
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

func (self *Declaration) AddDefvar(v *entity.DefinedVariable) {
  self.Defvars = append(self.Defvars, v)
}

func (self *Declaration) AddVardecl(v *entity.UndefinedVariable) {
  self.Vardecls = append(self.Vardecls, v)
}

func (self *Declaration) AddDefun(f *entity.DefinedFunction) {
  self.Defuns = append(self.Defuns, f)
}

func (self *Declaration) AddFuncdecl(f *entity.UndefinedFunction) {
  self.Funcdecls = append(self.Funcdecls, f)
}

func (self *Declaration) AddConstant(c *entity.Constant) {
  self.Constants = append(self.Constants, c)
}

func (self *Declaration) AddDefstruct(s *StructNode) {
  self.Defstructs = append(self.Defstructs, s)
}

func (self *Declaration) AddDefunion(u *UnionNode) {
  self.Defunions = append(self.Defunions, u)
}

func (self *Declaration) AddTypedef(t *TypedefNode) {
  self.Typedefs = append(self.Typedefs, t)
}

func (self *Declaration) AddDefvars(vs []*entity.DefinedVariable) {
  for i := range vs {
    self.AddDefvar(vs[i])
  }
}

func (self *Declaration) AddVardecls(vs []*entity.UndefinedVariable) {
  for i := range vs {
    self.AddVardecl(vs[i])
  }
}

func (self *Declaration) AddDefuns(fs []*entity.DefinedFunction) {
  for i := range fs {
    self.AddDefun(fs[i])
  }
}

func (self *Declaration) AddFuncdecls(fs []*entity.UndefinedFunction) {
  for i := range fs {
    self.AddFuncdecl(fs[i])
  }
}

func (self *Declaration) AddConstants(cs []*entity.Constant) {
  for i := range cs {
    self.AddConstant(cs[i])
  }
}

func (self *Declaration) AddDefstructs(ss []*StructNode) {
  for i := range ss {
    self.AddDefstruct(ss[i])
  }
}

func (self *Declaration) AddDefunions(us []*UnionNode) {
  for i := range us {
    self.AddDefunion(us[i])
  }
}

func (self *Declaration) AddTypedefs(ts []*TypedefNode) {
  for i := range ts {
    self.AddTypedef(ts[i])
  }
}

func (self *Declaration) GetDefvars() []*entity.DefinedVariable {
  return self.Defvars
}

func (self *Declaration) GetVardecls() []*entity.UndefinedVariable {
  return self.Vardecls
}

func (self *Declaration) GetDefuns() []*entity.DefinedFunction {
  return self.Defuns
}

func (self *Declaration) GetFuncdecls() []*entity.UndefinedFunction {
  return self.Funcdecls
}

func (self *Declaration) GetConstants() []*entity.Constant {
  return self.Constants
}

func (self *Declaration) GetDefstructs() []*StructNode {
  return self.Defstructs
}

func (self *Declaration) GetDefunions() []*UnionNode {
  return self.Defunions
}

func (self *Declaration) GetTypedefs() []*TypedefNode {
  return self.Typedefs
}

func (self *Declaration) AddDeclaration(other *Declaration) {
  self.AddDefvars(other.GetDefvars())
  self.AddVardecls(other.GetVardecls())
  self.AddDefuns(other.GetDefuns())
  self.AddFuncdecls(other.GetFuncdecls())
  self.AddConstants(other.GetConstants())
  self.AddDefstructs(other.GetDefstructs())
  self.AddDefunions(other.GetDefunions())
  self.AddTypedefs(other.GetTypedefs())
}
