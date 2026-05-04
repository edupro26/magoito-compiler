package semantics

import (
	"fmt"
	"magoito-compiler/internal/ast"
)

type Validator struct {
	table  *SymbolTable
	errors *ErrorList
}

func Validate(program *ast.Program) error {
	validator := &Validator{
		table:  newSymbolTable(nil),
		errors: &ErrorList{},
	}

	validator.collectTopLevel(program)
	validator.checkProgram(program)

	if validator.errors.HasErrors() {
		return validator.errors
	}
	return nil
}

func (v *Validator) collectTopLevel(p *ast.Program) {
	for _, decl := range p.Declarations {
		switch d := decl.(type) {
		case *ast.ConstDecl:
			if d.Name != "_" && !v.table.define(d.Name, d.Type, CONST) {
				v.errors.Add(&d.Pos, fmt.Sprintf(DUP_TOP_DECL+" '%s'", d.Name))
			}
		case *ast.FunDecl:
			if d.Name != "_" && !v.table.define(d.Name, d.Type, FUN) {
				v.errors.Add(&d.Pos, fmt.Sprintf(DUP_TOP_DECL+" '%s'", d.Name))
			}
		}
	}
}

func (v *Validator) checkProgram(p *ast.Program) {
	main, ok := v.table.lookup("main")
	if !ok || main.Kind != FUN {
		v.errors.Add(nil, MISSING_MAIN)
	} else {
		if !v.isValidMain(main.Type) {
			v.errors.Add(nil, "main must have type Unit -> Unit")
		}
	}
	for _, d := range p.Declarations {
		switch decl := d.(type) {
		case *ast.ConstDecl:
			v.checkConstDecl(decl)
		case *ast.FunDecl:
			v.checkFunDecl(decl)
		}
	}
}

func (v *Validator) checkConstDecl(d *ast.ConstDecl) {
	if !isValidConst(d.Value) {
		v.errors.Add(&d.Pos, BAD_CONST_INIT)
	}
	if containsFunctionType(d.Type) {
		v.errors.Add(&d.Pos, "functions are not first-class values")
	}
	v.validateType(d.Type)
	v.checkType(d.Value, d.Type)
}

func (v *Validator) checkFunDecl(d *ast.FunDecl) {
	ft, ok := d.Type.(*ast.FunctionType)
	if !ok {
		v.errors.Add(&d.Pos, "function declaration must have a function type")
		return
	}
	if len(d.Params) != len(ft.Params) {
		v.errors.Add(&d.Pos, "function parameter count does not match type")
	}
	for _, param := range ft.Params {
		v.validateType(param)
		if containsFunctionType(param) {
			v.errors.Add(&d.Pos, "function types are not allowed as parameters")
		}
	}

	v.validateType(ft.Return)
	if containsFunctionType(ft.Return) {
		v.errors.Add(&d.Pos, "function types are not allowed as return values")
	}

	local := newSymbolTable(v.table)
	seen := map[string]struct{}{}
	for i, name := range d.Params {
		if name == "_" {
			continue
		}
		if name == "print" {
			v.errors.Add(&d.Pos, "parameter may not be named 'print'")
			continue
		}
		if _, exists := seen[name]; exists {
			v.errors.Add(&d.Pos, "duplicate parameter name: "+name)
			continue
		}
		seen[name] = struct{}{}
		if i < len(ft.Params) {
			local.define(name, ft.Params[i], VAR)
		}
	}

	v.withTableType(local, func() ast.Type {
		v.checkType(d.Body, ft.Return)
		return nil
	})
}

func (v *Validator) validateType(t ast.Type) {
	switch typ := t.(type) {
	case *ast.BasicType:
		return
	case *ast.RecordType:
		seen := make(map[string]struct{}, len(typ.Fields))
		for _, f := range typ.Fields {
			if _, exists := seen[f.Label]; exists {
				v.errors.Add(&f.Pos, DUP_RECORD_FIELD+f.Label)
				continue
			}
			seen[f.Label] = struct{}{}
			v.validateType(f.Type)
		}
	case *ast.FunctionType:
		for _, p := range typ.Params {
			v.validateType(p)
		}
		v.validateType(typ.Return)
	}
}

func (v *Validator) isSameType(a, b ast.Type) bool {
	switch ta := a.(type) {
	case *ast.BasicType:
		tb, ok := b.(*ast.BasicType)
		return ok && ta.Name == tb.Name
	case *ast.RecordType:
		tb, ok := b.(*ast.RecordType)
		if !ok {
			return false
		}
		ma := recordFieldMap(ta)
		mb := recordFieldMap(tb)
		if len(ma) != len(mb) {
			return false
		}
		for label, at := range ma {
			bt, ok := mb[label]
			if !ok || !v.isSameType(at, bt) {
				return false
			}
		}
		return true
	case *ast.FunctionType:
		tb, ok := b.(*ast.FunctionType)
		if !ok || len(ta.Params) != len(tb.Params) {
			return false
		}
		for i := range ta.Params {
			if !v.isSameType(ta.Params[i], tb.Params[i]) {
				return false
			}
		}
		return v.isSameType(ta.Return, tb.Return)
	default:
		return false
	}
}

func isValidConst(e ast.Expr) bool {
	switch n := e.(type) {
	case *ast.IntLiteral, *ast.StringLiteral, *ast.BoolLiteral, *ast.UnitLiteral:
		return true
	case *ast.UnaryExpr:
		return isValidConst(n.Operand)
	case *ast.BinaryExpr:
		return isValidConst(n.Left) && isValidConst(n.Right)
	case *ast.RecordExpr:
		for _, f := range n.Fields {
			if !isValidConst(f.Value) {
				return false
			}
		}
		return true
	case *ast.ProjectionExpr:
		return isValidConst(n.Record)
	default:
		return false
	}
}

func (v *Validator) isValidMain(t ast.Type) bool {
	typ, ok := t.(*ast.FunctionType)
	if !ok || len(typ.Params) != 1 {
		return false
	}
	return v.isSameType(typ.Params[0], unitType(ast.Position{})) &&
		v.isSameType(typ.Return, unitType(ast.Position{}))
}
