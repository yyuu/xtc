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

func asDefinedFunction(x core.IEntity) core.IDefinedFunction {
  return x.(core.IDefinedFunction)
}

func asUndefinedFunction(x core.IEntity) core.IUndefinedFunction {
  return x.(core.IUndefinedFunction)
}

func asVariable(x core.IEntity) core.IVariable {
  return x.(core.IVariable)
}

func asVariables(xs []core.IEntity) []core.IVariable {
  ys := make([]core.IVariable, len(xs))
  for i := range xs {
    ys[i] = asVariable(xs[i])
  }
  return ys
}

func asDefinedVariable(x core.IEntity) core.IDefinedVariable {
  return x.(core.IDefinedVariable)
}

func asDefinedVariables(xs []core.IEntity) []core.IDefinedVariable {
  ys := make([]core.IDefinedVariable, len(xs))
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

func asUnionNode(x core.INode) core.IUnionNode {
  return x.(core.IUnionNode)
}

func asTypedefNode(x core.INode) core.ITypedefNode {
  return x.(core.ITypedefNode)
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

func defvars(xs...core.IDefinedVariable) []core.IDefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IDefinedVariable { }
  }
}

func vardecls(xs...*entity.UndefinedVariable) []*entity.UndefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []*entity.UndefinedVariable { }
  }
}

func defuns(xs...core.IDefinedFunction) []core.IDefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IDefinedFunction { }
  }
}

func funcdecls(xs...core.IUndefinedFunction) []core.IUndefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IUndefinedFunction { }
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

func defunions(xs...core.IUnionNode) []core.IUnionNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.IUnionNode { }
  }
}

func typedefs(xs...core.ITypedefNode) []core.ITypedefNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []core.ITypedefNode { }
  }
}
