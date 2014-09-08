package parser

import (
  "bitbucket.org/yyuu/bs/ast"
  "bitbucket.org/yyuu/bs/duck"
  "bitbucket.org/yyuu/bs/entity"
  "bitbucket.org/yyuu/bs/typesys"
)

func asExprNode(x duck.INode) duck.IExprNode {
  return x.(duck.IExprNode)
}

func asExprNodes(xs []duck.INode) []duck.IExprNode {
  ys := make([]duck.IExprNode, len(xs))
  for i := range xs {
    ys[i] = asExprNode(xs[i])
  }
  return ys
}

func asStmtNode(x duck.INode) duck.IStmtNode {
  return x.(duck.IStmtNode)
}

func asStmtNodes(xs []duck.INode) []duck.IStmtNode {
  ys := make([]duck.IStmtNode, len(xs))
  for i := range xs {
    ys[i] = asStmtNode(xs[i])
  }
  return ys
}

func asTypeNode(x duck.INode) duck.ITypeNode {
  return x.(duck.ITypeNode)
}

func asTypeNodes(xs []duck.INode) []duck.ITypeNode {
  ys := make([]duck.ITypeNode, len(xs))
  for i := range xs {
    ys[i] = asTypeNode(xs[i])
  }
  return ys
}

func asSlot(x duck.INode) ast.Slot {
  return x.(ast.Slot)
}

func asSlots(xs []duck.INode) []ast.Slot {
  ys := make([]ast.Slot, len(xs))
  for i := range xs {
    ys[i] = asSlot(xs[i])
  }
  return ys
}

func asTypeDefinition(x duck.INode) duck.ITypeDefinition {
  return x.(duck.ITypeDefinition)
}

func asTypeDefinitions(xs []duck.INode) []duck.ITypeDefinition {
  ys := make([]duck.ITypeDefinition, len(xs))
  for i := range xs {
    ys[i] = asTypeDefinition(xs[i])
  }
  return ys
}

func asDeclarations(x duck.INode) ast.Declarations {
  return x.(ast.Declarations)
}

func asDefinedFunction(x duck.IEntity) duck.IDefinedFunction {
  return x.(duck.IDefinedFunction)
}

func asUndefinedFunction(x duck.IEntity) duck.IUndefinedFunction {
  return x.(duck.IUndefinedFunction)
}

func asVariable(x duck.IEntity) duck.IVariable {
  return x.(duck.IVariable)
}

func asVariables(xs []duck.IEntity) []duck.IVariable {
  ys := make([]duck.IVariable, len(xs))
  for i := range xs {
    ys[i] = asVariable(xs[i])
  }
  return ys
}

func asDefinedVariable(x duck.IEntity) duck.IDefinedVariable {
  return x.(duck.IDefinedVariable)
}

func asDefinedVariables(xs []duck.IEntity) []duck.IDefinedVariable {
  ys := make([]duck.IDefinedVariable, len(xs))
  for i := range xs {
    ys[i] = asDefinedVariable(xs[i])
  }
  return ys
}

func asUndefinedVariable(x duck.IEntity) duck.IUndefinedVariable {
  return x.(duck.IUndefinedVariable)
}

func asConstant(x duck.IEntity) duck.IConstant {
  return x.(duck.IConstant)
}

func asStructNode(x duck.INode) ast.StructNode {
  return x.(ast.StructNode)
}

func asUnionNode(x duck.INode) ast.UnionNode {
  return x.(ast.UnionNode)
}

func asTypedefNode(x duck.INode) ast.TypedefNode {
  return x.(ast.TypedefNode)
}

func asParams(x duck.IEntity) entity.Params {
  return x.(entity.Params)
}

func asParameter(x duck.IEntity) entity.Parameter {
  return x.(entity.Parameter)
}

func parametersTypeRef(params entity.Params) typesys.ParamTypeRefs {
  paramDescs := params.GetParamDescs()
  ps := make([]duck.ITypeRef, len(paramDescs))
  for i := range paramDescs {
    ps[i] = paramDescs[i].GetTypeNode().GetTypeRef()
  }
  return typesys.NewParamTypeRefs(params.GetLocation(), ps, false)
}

func defvars(xs...duck.IDefinedVariable) []duck.IDefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []duck.IDefinedVariable { }
  }
}

func vardecls(xs...duck.IUndefinedVariable) []duck.IUndefinedVariable {
  if 0 < len(xs) {
    return xs
  } else {
    return []duck.IUndefinedVariable { }
  }
}

func defuns(xs...duck.IDefinedFunction) []duck.IDefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []duck.IDefinedFunction { }
  }
}

func funcdecls(xs...duck.IUndefinedFunction) []duck.IUndefinedFunction {
  if 0 < len(xs) {
    return xs
  } else {
    return []duck.IUndefinedFunction { }
  }
}

func defconsts(xs...duck.IConstant) []duck.IConstant {
  if 0 < len(xs) {
    return xs
  } else {
    return []duck.IConstant { }
  }
}

func defstructs(xs...ast.StructNode) []ast.StructNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []ast.StructNode { }
  }
}

func defunions(xs...ast.UnionNode) []ast.UnionNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []ast.UnionNode { }
  }
}

func typedefs(xs...ast.TypedefNode) []ast.TypedefNode {
  if 0 < len(xs) {
    return xs
  } else {
    return []ast.TypedefNode { }
  }
}
