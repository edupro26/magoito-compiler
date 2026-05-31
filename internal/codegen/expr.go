package codegen

import (
	"fmt"
	"magoito-compiler/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (ctx *Context) compileExpr(e ast.Expr) (value.Value, ast.Type, error) {
	switch expr := e.(type) {
	case *ast.IntLiteral:
		return ctx.compileIntLiteral(expr)
	case *ast.BoolLiteral:
		return ctx.compileBoolLiteral(expr)
	case *ast.UnitLiteral:
		return ctx.compileUnitLiteral()
	case *ast.StringLiteral:
		return ctx.compileStringLiteral(expr)
	case *ast.Identifier:
		return ctx.compileIdentifier(expr.Name)
	case *ast.UnaryExpr:
		return ctx.compileUnaryExpr(expr)
	case *ast.BinaryExpr:
		return ctx.compileBinaryExpr(expr)
	case *ast.SequenceExpr:
		return ctx.compileSequenceExpr(expr)
	case *ast.IfExpr:
		return ctx.compileIfExpr(expr)
	case *ast.WhileExpr:
		return ctx.compileWhileExpr(expr)
	case *ast.VarDecl:
		return ctx.compileVarDecl(expr, false)
	case *ast.Assignment:
		return ctx.compileAssignment(expr)
	case *ast.CallExpr:
		return ctx.compileCallExpr(expr)
	case *ast.RecordExpr, *ast.ProjectionExpr:
		return nil, nil, fmt.Errorf("records are not supported in codegen")
	}
	return nil, nil, nil
}

func (*Context) compileIntLiteral(e *ast.IntLiteral) (value.Value, ast.Type, error) {
	return constant.NewInt(types.I32, int64(e.Value)), &ast.BasicType{Name: "Int"}, nil
}

func (*Context) compileBoolLiteral(e *ast.BoolLiteral) (value.Value, ast.Type, error) {
	if e.Value {
		return constant.NewInt(types.I1, 1), &ast.BasicType{Name: "Bool"}, nil
	}
	return constant.NewInt(types.I1, 0), &ast.BasicType{Name: "Bool"}, nil
}

func (*Context) compileUnitLiteral() (value.Value, ast.Type, error) {
	return constant.NewInt(types.I1, 0), &ast.BasicType{Name: "Unit"}, nil
}

func (ctx *Context) compileStringLiteral(e *ast.StringLiteral) (value.Value, ast.Type, error) {
	gl, err := ctx.getStringGlobal(e.Value)
	if err != nil {
		return nil, nil, err
	}
	ind := constant.NewInt(types.I32, 0)
	arrTyp := types.NewArray(uint64(len(e.Value)+1), types.I8)
	return ctx.currentBlock.NewGetElementPtr(arrTyp, gl, ind, ind),
		&ast.BasicType{Name: "String"}, nil
}

func (ctx *Context) compileIdentifier(name string) (value.Value, ast.Type, error) {
	if v, typ, ok := ctx.scope.lookup(name); ok {
		t, err := ctx.compileType(typ)
		if err != nil {
			return nil, nil, err
		}
		val := ctx.currentBlock.NewLoad(t, v)
		return val, typ, nil
	}
	if gl, ok := ctx.consts[name]; ok {
		typ := ctx.constTypes[name]
		t, err := ctx.compileType(typ)
		if err != nil {
			return nil, nil, err
		}
		val := ctx.currentBlock.NewLoad(t, gl)
		return val, typ, nil
	}
	return nil, nil, nil
}

func (ctx *Context) compileUnaryExpr(e *ast.UnaryExpr) (value.Value, ast.Type, error) {
	op, typ, err := ctx.compileExpr(e.Operand)
	if err != nil {
		return nil, nil, err
	}
	switch e.Op {
	case "-":
		return ctx.currentBlock.NewSub(constant.NewInt(types.I32, 0), op), typ, nil
	case "!":
		return ctx.currentBlock.NewXor(op, constant.NewInt(types.I1, 1)), typ, nil
	}
	return nil, nil, nil
}

func (ctx *Context) compileSequenceExpr(e *ast.SequenceExpr) (value.Value, ast.Type, error) {
	if decl, ok := e.Left.(*ast.VarDecl); ok {
		prevScope := ctx.scope
		ctx.scope = newScope(prevScope)
		defer func() { ctx.scope = prevScope }()
		if _, _, err := ctx.compileVarDecl(decl, true); err != nil {
			return nil, nil, err
		}
		return ctx.compileExpr(e.Right)
	}
	if _, _, err := ctx.compileExpr(e.Left); err != nil {
		return nil, nil, err
	}
	return ctx.compileExpr(e.Right)
}

func (ctx *Context) compileIfExpr(e *ast.IfExpr) (value.Value, ast.Type, error) {
	ctx.labelCounter++
	then := ctx.currentFunc.NewBlock(fmt.Sprintf("then%d", ctx.labelCounter))
	ctx.labelCounter++
	merge := ctx.currentFunc.NewBlock(fmt.Sprintf("merge%d", ctx.labelCounter))

	condVal, _, err := ctx.compileExpr(e.Condition)
	if err != nil {
		return nil, nil, err
	}
	if e.Else == nil {
		ctx.currentBlock.NewCondBr(condVal, then, merge)
		ctx.currentBlock = then
		if _, _, err := ctx.compileExpr(e.Then); err != nil {
			return nil, nil, err
		}
		if ctx.currentBlock.Term == nil {
			ctx.currentBlock.NewBr(merge)
		}
		ctx.currentBlock = merge
		return constant.NewInt(types.I1, 0), &ast.BasicType{Name: "Unit"}, nil
	}

	ctx.labelCounter++
	els := ctx.currentFunc.NewBlock(fmt.Sprintf("else%d", ctx.labelCounter))
	ctx.currentBlock.NewCondBr(condVal, then, els)

	ctx.currentBlock = then
	thenVal, thenType, err := ctx.compileExpr(e.Then)
	if err != nil {
		return nil, nil, err
	}
	thenEnd := ctx.currentBlock
	if thenEnd.Term == nil {
		thenEnd.NewBr(merge)
	}

	ctx.currentBlock = els
	elseVal, elseType, err := ctx.compileExpr(e.Else)
	if err != nil {
		return nil, nil, err
	}
	elseEnd := ctx.currentBlock
	if elseEnd.Term == nil {
		elseEnd.NewBr(merge)
	}

	thenBasic, okThen := thenType.(*ast.BasicType)
	elseBasic, okElse := elseType.(*ast.BasicType)
	if okThen && okElse && thenBasic.Name != elseBasic.Name {
		return nil, nil, fmt.Errorf("if branches must have same type")
	}

	ctx.currentBlock = merge
	phi := merge.NewPhi(
		ir.NewIncoming(thenVal, thenEnd),
		ir.NewIncoming(elseVal, elseEnd),
	)
	phiType, err := ctx.compileType(thenType)
	if err != nil {
		return nil, nil, err
	}
	phi.Typ = phiType
	return phi, thenType, nil
}

func (ctx *Context) compileWhileExpr(e *ast.WhileExpr) (value.Value, ast.Type, error) {
	ctx.labelCounter++
	cond := ctx.currentFunc.NewBlock(fmt.Sprintf("cond%d", ctx.labelCounter))
	ctx.labelCounter++
	body := ctx.currentFunc.NewBlock(fmt.Sprintf("body%d", ctx.labelCounter))
	ctx.labelCounter++
	end := ctx.currentFunc.NewBlock(fmt.Sprintf("end%d", ctx.labelCounter))

	ctx.currentBlock.NewBr(cond)
	ctx.currentBlock = cond
	condVal, _, err := ctx.compileExpr(e.Condition)
	if err != nil {
		return nil, nil, err
	}
	ctx.currentBlock.NewCondBr(condVal, body, end)

	ctx.currentBlock = body
	if _, _, err := ctx.compileExpr(e.Body); err != nil {
		return nil, nil, err
	}
	if ctx.currentBlock.Term == nil {
		ctx.currentBlock.NewBr(cond)
	}

	ctx.currentBlock = end
	return constant.NewInt(types.I1, 0), &ast.BasicType{Name: "Unit"}, nil
}

func (ctx *Context) compileVarDecl(e *ast.VarDecl, bind bool) (value.Value, ast.Type, error) {
	val, _, err := ctx.compileExpr(e.Value)
	if err != nil {
		return nil, nil, err
	}
	if bind && e.Name != "_" {
		t, err := ctx.compileType(e.Type)
		if err != nil {
			return nil, nil, err
		}
		if !types.Equal(val.Type(), t) {
			return nil, nil, fmt.Errorf("cannot store %s in %s", val.Type(), t)
		}
		ptr := ctx.currentBlock.NewAlloca(t)
		ctx.currentBlock.NewStore(val, ptr)
		ctx.scope.define(e.Name, ptr, e.Type)
	}
	return constant.NewInt(types.I1, 0), &ast.BasicType{Name: "Unit"}, nil
}

func (ctx *Context) compileAssignment(e *ast.Assignment) (value.Value, ast.Type, error) {
	val, _, err := ctx.compileExpr(e.Value)
	if err != nil {
		return nil, nil, err
	}
	ptr, declType, ok := ctx.scope.lookup(e.Name)
	if !ok {
		return nil, nil, fmt.Errorf("assignment to undefined variable: %s", e.Name)
	}
	llvmType, err := ctx.compileType(declType)
	if err != nil {
		return nil, nil, err
	}
	if !types.Equal(val.Type(), llvmType) {
		return nil, nil, fmt.Errorf("cannot store %s in %s", val.Type(), llvmType)
	}
	ctx.currentBlock.NewStore(val, ptr)
	return constant.NewInt(types.I1, 0), &ast.BasicType{Name: "Unit"}, nil
}

func (ctx *Context) compileCallExpr(e *ast.CallExpr) (value.Value, ast.Type, error) {
	if e.Callee == "print" {
		val, typ, err := ctx.compileExpr(e.Args[0])
		if err != nil {
			return nil, nil, err
		}
		if err := ctx.compilePrintln(val, typ); err != nil {
			return nil, nil, err
		}
		return constant.NewInt(types.I1, 0), &ast.BasicType{Name: "Unit"}, nil
	}
	args := make([]value.Value, len(e.Args))
	for i, arg := range e.Args {
		val, _, err := ctx.compileExpr(arg)
		if err != nil {
			return nil, nil, err
		}
		args[i] = val
	}
	call := ctx.currentBlock.NewCall(ctx.functions[e.Callee], args...)
	return call, ctx.funcTypes[e.Callee].Return, nil
}
