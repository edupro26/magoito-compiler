package semantics

import "magoito-compiler/internal/ast"

func (v *Validator) inferType(expr ast.Expr) ast.Type {
	switch e := expr.(type) {
	case *ast.IntLiteral:
		return intType(e.Pos)
	case *ast.StringLiteral:
		return stringType(e.Pos)
	case *ast.BoolLiteral:
		return boolType(e.Pos)
	case *ast.UnitLiteral:
		return unitType(e.Pos)
	case *ast.Identifier:
		sym, ok := v.table.lookup(e.Name)
		if !ok {
			v.errors.Add(&e.Pos, "undefined identifier: "+e.Name)
			return nil
		}
		if sym.Kind == FUN {
			v.errors.Add(&e.Pos, "function used as value: "+e.Name)
			return nil
		}
		return sym.Type
	case *ast.UnaryExpr:
		return v.inferUnary(e)
	case *ast.BinaryExpr:
		return v.inferBinary(e)
	case *ast.SequenceExpr:
		v.inferType(e.Left)
		nextTable := v.table
		if decl, ok := e.Left.(*ast.VarDecl); ok {
			nextTable = v.extendVarScope(decl)
		}
		return v.withTableType(nextTable, func() ast.Type {
			return v.inferType(e.Right)
		})
	case *ast.IfExpr:
		return v.inferIf(e)
	case *ast.WhileExpr:
		return v.inferWhile(e)
	case *ast.VarDecl:
		return v.inferVarDecl(e)
	case *ast.Assignment:
		return v.inferAssignment(e)
	case *ast.CallExpr:
		return v.inferCall(e)
	case *ast.RecordExpr:
		return v.inferRecord(e)
	case *ast.ProjectionExpr:
		return v.inferProjection(e)

	}
	return nil
}

func (v *Validator) checkType(expr ast.Expr, expect ast.Type) bool {
	if expect == nil {
		v.inferType(expr)
		return true
	}

	switch e := expr.(type) {
	case *ast.IfExpr:
		if !v.checkType(e.Condition, boolType(e.Pos)) {
			v.errors.Add(&e.Pos, "if condition must be Bool")
		}
		if e.Else == nil {
			v.checkType(e.Then, unitType(e.Pos))
			return v.expectType(&e.Pos, unitType(e.Pos), expect)
		}
		v.checkType(e.Then, expect)
		v.checkType(e.Else, expect)
		return true
	case *ast.WhileExpr:
		if !v.checkType(e.Condition, boolType(e.Pos)) {
			v.errors.Add(&e.Pos, "while condition must be Bool")
		}
		v.inferType(e.Body)
		return v.expectType(&e.Pos, unitType(e.Pos), expect)
	case *ast.SequenceExpr:
		v.inferType(e.Left)
		nextTable := v.table
		if decl, ok := e.Left.(*ast.VarDecl); ok {
			nextTable = v.extendVarScope(decl)
		}
		return v.withTableType(nextTable, func() ast.Type {
			if v.checkType(e.Right, expect) {
				return expect
			}
			return nil
		}) != nil
	case *ast.VarDecl:
		v.inferVarDecl(e)
		return v.expectType(&e.Pos, unitType(e.Pos), expect)
	case *ast.Assignment:
		v.inferAssignment(e)
		return v.expectType(&e.Pos, unitType(e.Pos), expect)
	default:
		inferred := v.inferType(expr)
		if inferred == nil {
			return false
		}
		if !v.isSubtype(inferred, expect) {
			v.reportTypeMismatch(exprPos(expr), expect, inferred)
			return false
		}
		return true
	}
}

func (v *Validator) isSubtype(sub, sup ast.Type) bool {
	if sub == nil || sup == nil {
		return false
	}
	if v.isSameType(sub, sup) {
		return true
	}
	subRec, okSub := sub.(*ast.RecordType)
	supRec, okSup := sup.(*ast.RecordType)
	if okSub && okSup {
		subFields := recordFieldMap(subRec)
		for _, sf := range supRec.Fields {
			st, ok := subFields[sf.Label]
			if !ok || !v.isSubtype(st, sf.Type) {
				return false
			}
		}
		return true
	}
	return false
}

func (v *Validator) inferUnary(e *ast.UnaryExpr) ast.Type {
	switch e.Op {
	case "-":
		v.checkType(e.Operand, intType(e.Pos))
		return intType(e.Pos)
	case "!":
		v.checkType(e.Operand, boolType(e.Pos))
		return boolType(e.Pos)
	default:
		return nil
	}
}

func (v *Validator) inferBinary(e *ast.BinaryExpr) ast.Type {
	switch e.Op {
	case "+", "-", "*", "/", "%", "^":
		v.checkType(e.Left, intType(e.Pos))
		v.checkType(e.Right, intType(e.Pos))
		return intType(e.Pos)
	case "<", "<=", ">", ">=":
		v.checkType(e.Left, intType(e.Pos))
		v.checkType(e.Right, intType(e.Pos))
		return boolType(e.Pos)
	case "&&", "||":
		v.checkType(e.Left, boolType(e.Pos))
		v.checkType(e.Right, boolType(e.Pos))
		return boolType(e.Pos)
	case "==", "!=":
		left := v.inferType(e.Left)
		right := v.inferType(e.Right)
		if left != nil && right != nil && !v.isSameType(left, right) {
			v.errors.Add(&e.Pos, "equality operands must have the same type")
		}
		if containsFunctionType(left) || containsFunctionType(right) {
			v.errors.Add(&e.Pos, "functions cannot be compared for equality")
		}
		return boolType(e.Pos)
	default:
		return nil
	}
}

func (v *Validator) inferIf(e *ast.IfExpr) ast.Type {
	v.checkType(e.Condition, boolType(e.Pos))
	thenType := v.inferType(e.Then)
	if e.Else == nil {
		if thenType != nil && !v.isSameType(thenType, unitType(e.Pos)) {
			v.errors.Add(&e.Pos, "if-then expression must have type Unit")
		}
		return unitType(e.Pos)
	}
	elseType := v.inferType(e.Else)
	if thenType != nil && elseType != nil && !v.isSameType(thenType, elseType) {
		v.errors.Add(&e.Pos, "if branches must have the same type")
	}
	return thenType
}

func (v *Validator) inferWhile(e *ast.WhileExpr) ast.Type {
	v.checkType(e.Condition, boolType(e.Pos))
	v.inferType(e.Body)
	return unitType(e.Pos)
}

func (v *Validator) inferVarDecl(e *ast.VarDecl) ast.Type {
	if e.Name == "print" {
		v.errors.Add(&e.Pos, "variable may not be named 'print'")
	}
	if containsFunctionType(e.Type) {
		v.errors.Add(&e.Pos, "functions are not first-class values")
	}
	v.validateType(e.Type)
	v.checkType(e.Value, e.Type)
	return unitType(e.Pos)
}

func (v *Validator) inferAssignment(e *ast.Assignment) ast.Type {
	sym, ok := v.table.lookup(e.Name)
	if !ok {
		v.errors.Add(&e.Pos, "undefined identifier: "+e.Name)
		return unitType(e.Pos)
	}
	if sym.Kind != VAR {
		v.errors.Add(&e.Pos, "cannot assign to non-variable: "+e.Name)
		return unitType(e.Pos)
	}
	v.checkType(e.Value, sym.Type)
	return unitType(e.Pos)
}

func (v *Validator) inferCall(e *ast.CallExpr) ast.Type {
	if e.Callee == "print" {
		for _, arg := range e.Args {
			v.inferType(arg)
		}
		if len(e.Args) != 1 {
			v.errors.Add(&e.Pos, "print expects exactly one argument")
		}
		return unitType(e.Pos)
	}

	sym, ok := v.table.lookup(e.Callee)
	if !ok {
		v.errors.Add(&e.Pos, "undefined function: "+e.Callee)
		for _, arg := range e.Args {
			v.inferType(arg)
		}
		return nil
	}
	if sym.Kind != FUN {
		v.errors.Add(&e.Pos, "identifier is not a function: "+e.Callee)
		for _, arg := range e.Args {
			v.inferType(arg)
		}
		return nil
	}
	ft, ok := sym.Type.(*ast.FunctionType)
	if !ok {
		v.errors.Add(&e.Pos, "identifier does not have a function type: "+e.Callee)
		for _, arg := range e.Args {
			v.inferType(arg)
		}
		return nil
	}
	if len(e.Args) != len(ft.Params) {
		v.errors.Add(&e.Pos, "wrong number of arguments for "+e.Callee)
	}
	for i, arg := range e.Args {
		if i < len(ft.Params) {
			v.checkType(arg, ft.Params[i])
		} else {
			v.inferType(arg)
		}
	}
	return ft.Return
}

func (v *Validator) inferRecord(e *ast.RecordExpr) ast.Type {
	fields := make([]ast.RecordTypeField, 0, len(e.Fields))
	seen := map[string]ast.Position{}
	for _, f := range e.Fields {
		if _, exists := seen[f.Label]; exists {
			v.errors.Add(&f.Pos, DUP_RECORD_FIELD+f.Label)
		} else {
			seen[f.Label] = f.Pos
		}
		fieldType := v.inferType(f.Value)
		fields = append(fields, ast.RecordTypeField{Label: f.Label, Type: fieldType, Pos: f.Pos})
	}
	return &ast.RecordType{Fields: fields, Pos: e.Pos}
}

func (v *Validator) inferProjection(e *ast.ProjectionExpr) ast.Type {
	base := v.inferType(e.Record)
	rec, ok := base.(*ast.RecordType)
	if !ok {
		if base != nil {
			v.errors.Add(&e.Pos, "projection on non-record value")
		}
		return nil
	}
	for _, f := range rec.Fields {
		if f.Label == e.Field {
			return f.Type
		}
	}
	v.errors.Add(&e.Pos, "record has no field: "+e.Field)
	return nil
}

func (v *Validator) extendVarScope(decl *ast.VarDecl) *SymbolTable {
	if decl.Name == "_" {
		return v.table
	}
	return v.table.extend(decl.Name, decl.Type, VAR)
}

func (v *Validator) withTableType(table *SymbolTable, fn func() ast.Type) ast.Type {
	prev := v.table
	v.table = table
	res := fn()
	v.table = prev
	return res
}

func (v *Validator) reportTypeMismatch(pos *ast.Position, expected, actual ast.Type) {
	if pos == nil {
		return
	}
	v.errors.Add(pos, TYPE_MISMATCH+": expected "+formatType(expected)+", got "+formatType(actual))
}

func (v *Validator) expectType(pos *ast.Position, actual, expected ast.Type) bool {
	if v.isSubtype(actual, expected) {
		return true
	}
	v.reportTypeMismatch(pos, expected, actual)
	return false
}

func exprPos(expr ast.Expr) *ast.Position {
	switch e := expr.(type) {
	case *ast.IntLiteral:
		return &e.Pos
	case *ast.StringLiteral:
		return &e.Pos
	case *ast.BoolLiteral:
		return &e.Pos
	case *ast.UnitLiteral:
		return &e.Pos
	case *ast.Identifier:
		return &e.Pos
	case *ast.UnaryExpr:
		return &e.Pos
	case *ast.BinaryExpr:
		return &e.Pos
	case *ast.SequenceExpr:
		return &e.Pos
	case *ast.IfExpr:
		return &e.Pos
	case *ast.WhileExpr:
		return &e.Pos
	case *ast.VarDecl:
		return &e.Pos
	case *ast.Assignment:
		return &e.Pos
	case *ast.CallExpr:
		return &e.Pos
	case *ast.RecordExpr:
		return &e.Pos
	case *ast.ProjectionExpr:
		return &e.Pos
	default:
		return nil
	}
}

func intType(pos ast.Position) ast.Type    { return &ast.BasicType{Name: "Int", Pos: pos} }
func stringType(pos ast.Position) ast.Type { return &ast.BasicType{Name: "String", Pos: pos} }
func boolType(pos ast.Position) ast.Type   { return &ast.BasicType{Name: "Bool", Pos: pos} }
func unitType(pos ast.Position) ast.Type   { return &ast.BasicType{Name: "Unit", Pos: pos} }
