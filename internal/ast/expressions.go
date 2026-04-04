package ast

// Expr is the marker interface for expressions
type Expr interface {
	exprNode()
}

// IntLiteral is an integer literal, e.g. 0, 42
type IntLiteral struct {
	Value int
}

// StringLiteral is a string literal, e.g. "hello"
type StringLiteral struct {
	Value string
}

// BoolLiteral is one of: true, false
type BoolLiteral struct {
	Value bool
}

// UnitLiteral is the unit value: unit
type UnitLiteral struct{}

// Identifier is a variable reference
type Identifier struct {
	Name string
}

// BinaryExpr represents a binary operation: left op right
type BinaryExpr struct {
	Left  Expr
	Op    string
	Right Expr
}

// UnaryExpr represents a unary operation: op operand
type UnaryExpr struct {
	Op      string
	Operand Expr
}

// SequenceExpr represents sequential composition: left ; right
type SequenceExpr struct {
	Left  Expr
	Right Expr
}

// IfExpr represents: if cond then expr else expr
type IfExpr struct {
	Condition Expr
	Then      Expr
	Else      Expr
}

// WhileExpr represents: while cond do body
type WhileExpr struct {
	Condition Expr
	Body      Expr
}

// VarDecl represents: var id : type = value
type VarDecl struct {
	Name  string
	Type  Type
	Value Expr
}

// Assignment represents: id = value
type Assignment struct {
	Name  string
	Value Expr
}

// CallExpr represents a function call: id(args)
type CallExpr struct {
	Callee string
	Args   []Expr
}

// RecordExpr represents a record construction: {id1 = expr1, ..., idn = exprn}
type RecordExpr struct {
	Fields []RecordField
}

// RecordField is one field in a record construction
type RecordField struct {
	Label string
	Value Expr
}

// ProjectionExpr represents record projection: record.field
type ProjectionExpr struct {
	Record Expr
	Field  string
}

func (*IntLiteral) exprNode()     {}
func (*StringLiteral) exprNode()  {}
func (*BoolLiteral) exprNode()    {}
func (*UnitLiteral) exprNode()    {}
func (*Identifier) exprNode()     {}
func (*BinaryExpr) exprNode()     {}
func (*UnaryExpr) exprNode()      {}
func (*SequenceExpr) exprNode()   {}
func (*IfExpr) exprNode()         {}
func (*WhileExpr) exprNode()      {}
func (*VarDecl) exprNode()        {}
func (*Assignment) exprNode()     {}
func (*CallExpr) exprNode()       {}
func (*RecordExpr) exprNode()     {}
func (*ProjectionExpr) exprNode() {}
