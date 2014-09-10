package parser

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/core"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

func asExprNode(x core.INode) core.IExprNode {
  return x.(core.IExprNode)
}

func asExprNodes(xs []core.INode) []core.IExprNode {
  ys := make([]core.IExprNode, len(xs))
  for i := range xs {
    ys[i] = asExprNode(xs[i])
  }
  return ys
}

func asStmtNode(x core.INode) core.IStmtNode {
  return x.(core.IStmtNode)
}

func asStmtNodes(xs []core.INode) []core.IStmtNode {
  ys := make([]core.IStmtNode, len(xs))
  for i := range xs {
    ys[i] = asStmtNode(xs[i])
  }
  return ys
}

func asTypeNode(x core.INode) core.ITypeNode {
  return x.(core.ITypeNode)
}

func asTypeNodes(xs []core.INode) []core.ITypeNode {
  ys := make([]core.ITypeNode, len(xs))
  for i := range xs {
    ys[i] = asTypeNode(xs[i])
  }
  return ys
}

func asSlot(x core.INode) core.ISlot {
  return x.(core.ISlot)
}

func asSlots(xs []core.INode) []core.ISlot {
  ys := make([]core.ISlot, len(xs))
  for i := range xs {
    ys[i] = asSlot(xs[i])
  }
  return ys
}

func asTypeDefinition(x core.INode) core.ITypeDefinition {
  return x.(core.ITypeDefinition)
}

func asTypeDefinitions(xs []core.INode) []core.ITypeDefinition {
  ys := make([]core.ITypeDefinition, len(xs))
  for i := range xs {
    ys[i] = asTypeDefinition(xs[i])
  }
  return ys
}

func asDeclarations(x core.INode) *ast.Declarations {
  return x.(*ast.Declarations)
}

func asDefinedFunction(x core.IEntity) *entity.DefinedFunction {
  return x.(*entity.DefinedFunction)
}

func asUndefinedFunction(x core.IEntity) *entity.UndefinedFunction {
  return x.(*entity.UndefinedFunction)
}

func asDefinedVariable(x core.IEntity) *entity.DefinedVariable {
  return x.(*entity.DefinedVariable)
}

func asDefinedVariables(xs []core.IEntity) []*entity.DefinedVariable {
  ys := make([]*entity.DefinedVariable, len(xs))
  for i := range xs {
    ys[i] = asDefinedVariable(xs[i])
  }
  return ys
}

func asUndefinedVariable(x core.IEntity) *entity.UndefinedVariable {
  return x.(*entity.UndefinedVariable)
}

func asConstant(x core.IEntity) *entity.Constant {
  return x.(*entity.Constant)
}

func asStructNode(x core.INode) core.IStructNode {
  return x.(core.IStructNode)
}

func asUnionNode(x core.INode) *ast.UnionNode {
  return x.(*ast.UnionNode)
}

func asTypedefNode(x core.INode) *ast.TypedefNode {
  return x.(*ast.TypedefNode)
}

func asParams(x core.IEntity) *entity.Params {
  return x.(*entity.Params)
}

func asParameter(x core.IEntity) *entity.Parameter {
  return x.(*entity.Parameter)
}

func parametersTypeRef(params *entity.Params) *typesys.ParamTypeRefs {
  paramDescs := params.GetParamDescs()
  ps := make([]core.ITypeRef, len(paramDescs))
  for i := range paramDescs {
    ps[i] = paramDescs[i].GetTypeNode().GetTypeRef()
  }
  return typesys.NewParamTypeRefs(params.GetLocation(), ps, false)
}

func defvars(xs...*entity.DefinedVariable) []*entity.DefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []*entity.DefinedVariable { }
  }
}

func vardecls(xs...*entity.UndefinedVariable) []*entity.UndefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []*entity.UndefinedVariable { }
  }
}

func defuns(xs...*entity.DefinedFunction) []*entity.DefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []*entity.DefinedFunction { }
  }
}

func funcdecls(xs...*entity.UndefinedFunction) []*entity.UndefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []*entity.UndefinedFunction { }
  }
}

func defconsts(xs...*entity.Constant) []*entity.Constant {
  if 0 < len(xs) {
    return xs
  } else {
    return []*entity.Constant { }
  }
}

func defstructs(xs...core.IStructNode) []core.IStructNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IStructNode { }
  }
}

func defunions(xs...*ast.UnionNode) []*ast.UnionNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []*ast.UnionNode { }
  }
}

func typedefs(xs...*ast.TypedefNode) []*ast.TypedefNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []*ast.TypedefNode { }
  }
}
