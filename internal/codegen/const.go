package codegen

import (
	"fmt"
	"magoito-compiler/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

type constVal struct {
	intVal  int
	boolVal bool
	strVal  string
	unitVal bool
}

func (ctx *Context) evaluateConst(e ast.Expr) (constVal, error) {
	switch exp := e.(type) {
	case *ast.IntLiteral:
		return constVal{intVal: exp.Value}, nil
	case *ast.BoolLiteral:
		return constVal{boolVal: exp.Value}, nil
	case *ast.StringLiteral:
		return constVal{strVal: exp.Value}, nil
	case *ast.UnitLiteral:
		return constVal{unitVal: true}, nil
	case *ast.UnaryExpr:
		return ctx.evaluateUnaryExpr(exp), nil
	case *ast.BinaryExpr:
		return ctx.evaluateBinaryExpr(exp), nil
	case *ast.RecordExpr, *ast.ProjectionExpr:
		return constVal{}, fmt.Errorf("records are not supported in codegen")
	}
	return constVal{}, nil
}

func (ctx *Context) evaluateUnaryExpr(e *ast.UnaryExpr) constVal {
	v, _ := ctx.evaluateConst(e.Operand)
	switch e.Op {
	case "-":
		return constVal{intVal: -v.intVal}
	case "!":
		return constVal{boolVal: !v.boolVal}
	}
	return constVal{}
}

func (ctx *Context) evaluateBinaryExpr(e *ast.BinaryExpr) constVal {
	left, _ := ctx.evaluateConst(e.Left)
	right, _ := ctx.evaluateConst(e.Right)
	switch e.Op {
	case "+":
		return constVal{intVal: left.intVal + right.intVal}
	case "-":
		return constVal{intVal: left.intVal - right.intVal}
	case "*":
		return constVal{intVal: left.intVal * right.intVal}
	case "/":
		return constVal{intVal: left.intVal / right.intVal}
	case "%":
		return constVal{intVal: left.intVal % right.intVal}
	case "^":
		return constVal{intVal: powInt(left.intVal, right.intVal)}
	case "<":
		return constVal{boolVal: left.intVal < right.intVal}
	case "<=":
		return constVal{boolVal: left.intVal <= right.intVal}
	case ">":
		return constVal{boolVal: left.intVal > right.intVal}
	case ">=":
		return constVal{boolVal: left.intVal >= right.intVal}
	case "&&":
		return constVal{boolVal: left.boolVal && right.boolVal}
	case "||":
		return constVal{boolVal: left.boolVal || right.boolVal}
	case "==":
		return constVal{boolVal: constEqual(left, right)}
	case "!=":
		return constVal{boolVal: !constEqual(left, right)}
	}
	return constVal{}
}

func (ctx *Context) compileConst(v constVal, t ast.Type) constant.Constant {
	switch typ := t.(type) {
	case *ast.BasicType:
		switch typ.Name {
		case "Int":
			return constant.NewInt(types.I32, int64(v.intVal))
		case "Bool":
			if v.boolVal {
				return constant.NewInt(types.I1, 1)
			}
			return constant.NewInt(types.I1, 0)
		case "Unit":
			return constant.NewInt(types.I1, 0)
		case "String":
			global, _ := ctx.getStringGlobal(v.strVal)
			arrTyp := types.NewArray(uint64(len(v.strVal)+1), types.I8)
			ind := constant.NewInt(types.I32, 0)
			return constant.NewGetElementPtr(arrTyp, global, ind, ind)
		}
	}
	return nil
}

func (ctx *Context) getStringGlobal(value string) (*ir.Global, error) {
	if gl, ok := ctx.strings[value]; ok {
		return gl, nil
	}
	name := fmt.Sprintf(".str%d", ctx.stringCounter)
	ctx.stringCounter++
	constArr := constant.NewCharArrayFromString(value + "\x00")
	global := ctx.module.NewGlobalDef(name, constArr)
	global.Immutable = true
	ctx.strings[value] = global
	return global, nil
}

func powInt(base, exp int) int {
	if exp <= 0 {
		return 1
	}
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
	}
	return res
}

func constEqual(a, b constVal) bool {
	if a.strVal != "" || b.strVal != "" {
		return a.strVal == b.strVal
	}
	if a.unitVal || b.unitVal {
		return a.unitVal == b.unitVal
	}
	if a.boolVal != b.boolVal {
		return false
	}
	return a.intVal == b.intVal
}
