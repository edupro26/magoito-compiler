package codegen

import (
	"fmt"
	"magoito-compiler/internal/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type varScope struct {
	parent *varScope
	ptr    map[string]value.Value
	typ    map[string]ast.Type
}

func newScope(p *varScope) *varScope {
	return &varScope{
		parent: p,
		ptr:    map[string]value.Value{},
		typ:    map[string]ast.Type{},
	}
}

func (s *varScope) define(name string, ptr value.Value, typ ast.Type) {
	s.ptr[name] = ptr
	s.typ[name] = typ
}

func (s *varScope) lookup(name string) (value.Value, ast.Type, bool) {
	for scope := s; scope != nil; scope = scope.parent {
		if ptr, ok := scope.ptr[name]; ok {
			return ptr, scope.typ[name], true
		}
	}
	return nil, nil, false
}

type Context struct {
	module *ir.Module

	printf *ir.Func
	strcmp *ir.Func
	pow    *ir.Func

	currentFunc  *ir.Func
	currentBlock *ir.Block
	scope        *varScope

	constTypes map[string]ast.Type
	consts     map[string]*ir.Global

	strings map[string]*ir.Global

	funcTypes map[string]*ast.FunctionType
	functions map[string]*ir.Func

	stringCounter int
	labelCounter  int
}

func Generate(program *ast.Program) (*ir.Module, error) {
	ctx := &Context{
		module:     ir.NewModule(),
		constTypes: map[string]ast.Type{},
		consts:     map[string]*ir.Global{},
		strings:    map[string]*ir.Global{},
		funcTypes:  map[string]*ast.FunctionType{},
		functions:  map[string]*ir.Func{},
	}
	ctx.compileRuntimeHelpers()
	if err := ctx.compileConstDecls(program); err != nil {
		return nil, err
	}
	if err := ctx.compileFunctions(program); err != nil {
		return nil, err
	}
	return ctx.module, nil
}

func (ctx *Context) compileConstDecls(program *ast.Program) error {
	for _, decl := range program.Declarations {
		if c, ok := decl.(*ast.ConstDecl); ok {
			value, err := ctx.evaluateConst(c.Value)
			if err != nil {
				return err
			}
			name := fmt.Sprintf("const_%s", c.Name)
			literal := ctx.compileConst(value, c.Type)
			global := ctx.module.NewGlobalDef(name, literal)
			global.Immutable = true
			ctx.consts[c.Name] = global
			ctx.constTypes[c.Name] = c.Type
		}
	}
	return nil
}

func (ctx *Context) compileFunctions(program *ast.Program) error {
	for _, decl := range program.Declarations {
		if err := ctx.declareFunction(decl); err != nil {
			return err
		}
	}
	for _, decl := range program.Declarations {
		if fn, ok := decl.(*ast.FunDecl); ok {
			if err := ctx.compileFunction(fn); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ctx *Context) declareFunction(d ast.Declaration) error {
	switch decl := d.(type) {
	case *ast.FunDecl:
		ft, _ := decl.Type.(*ast.FunctionType)
		retTyp, err := ctx.compileType(ft.Return)
		if err != nil {
			return err
		}
		params := make([]*ir.Param, len(ft.Params))
		for i, p := range ft.Params {
			pt, err := ctx.compileType(p)
			if err != nil {
				return err
			}
			params[i] = ir.NewParam(fmt.Sprintf("p%d", i), pt)
		}
		fn := ctx.module.NewFunc(decl.Name, retTyp, params...)
		ctx.functions[decl.Name] = fn
		ctx.funcTypes[decl.Name] = ft
	}
	return nil
}
