package ast

type Program struct {
	Declarations []Declaration
}

// Position represents a source location in 1-based line and 0-based column.
type Position struct {
	Line int
	Col  int
}

// ─── Declarations ────────────────────────────────────────────────────────────

// Declaration is the marker interface for top-level declarations
type Declaration interface {
	declNode()
}

// ConstDecl represents: const id : type = expr
type ConstDecl struct {
	Name  string
	Type  Type
	Value Expr
	Pos   Position
}

// FunDecl represents: fun id (params) : type = expr
type FunDecl struct {
	Name   string
	Params []string
	Type   Type // must be a FunctionType
	Body   Expr
	Pos    Position
}

func (*ConstDecl) declNode() {}
func (*FunDecl) declNode()   {}

// ─── Types ───────────────────────────────────────────────────────────────────

// Type is the marker interface for type expressions
type Type interface {
	typeNode()
}

// BasicType represents types: Int, Bool, Unit, String
type BasicType struct {
	Name string
	Pos  Position
}

// RecordType represents: {id1: type1, ..., idn: typen}
type RecordType struct {
	Fields []RecordTypeField
	Pos    Position
}

// RecordTypeField is one field in a record type
type RecordTypeField struct {
	Label string
	Type  Type
	Pos   Position
}

// FunctionType represents: type -> type  or  (type1,...,typen) -> type
type FunctionType struct {
	Params []Type
	Return Type
	Pos    Position
}

func (*BasicType) typeNode()    {}
func (*RecordType) typeNode()   {}
func (*FunctionType) typeNode() {}

// Util
func AsType(v any) Type {
	if v == nil {
		return nil
	}
	if t, ok := v.(Type); ok {
		return t
	}
	return nil
}

func AsExpr(v any) Expr {
	if v == nil {
		return nil
	}
	if e, ok := v.(Expr); ok {
		return e
	}
	return nil
}
