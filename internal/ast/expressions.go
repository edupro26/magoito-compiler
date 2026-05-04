package ast

// Expr is the marker interface for expressions
type Expr interface {
	exprNode()
}

// IntLiteral is an integer literal, e.g. 0, 42
type IntLiteral struct {
	Value int
	Pos   Position
}

// StringLiteral is a string literal, e.g. "hello"
type StringLiteral struct {
	Value string
	Pos   Position
}

// BoolLiteral is one of: true, false
type BoolLiteral struct {
	Value bool
	Pos   Position
}

// UnitLiteral is the unit value: unit (no payload, but has a position).
type UnitLiteral struct {
	Pos Position
}

// Identifier is a variable reference
type Identifier struct {
	Name string
	Pos  Position
}

// BinaryExpr represents a binary operation: left op right
type BinaryExpr struct {
	Left  Expr
	Op    string
	Right Expr
	Pos   Position
}

// UnaryExpr represents a unary operation: op operand
type UnaryExpr struct {
	Op      string
	Operand Expr
	Pos     Position
}

// SequenceExpr represents sequential composition: left ; right
type SequenceExpr struct {
	Left  Expr
	Right Expr
	Pos   Position
}

// IfExpr represents: if cond then expr else expr
type IfExpr struct {
	Condition Expr
	Then      Expr
	Else      Expr
	Pos       Position
}

// WhileExpr represents: while cond do body
type WhileExpr struct {
	Condition Expr
	Body      Expr
	Pos       Position
}

// VarDecl represents: var id : type = value
type VarDecl struct {
	Name  string
	Type  Type
	Value Expr
	Pos   Position
}

// Assignment represents: id = value
type Assignment struct {
	Name  string
	Value Expr
	Pos   Position
}

// CallExpr represents a function call: id(args)
type CallExpr struct {
	Callee string
	Args   []Expr
	Pos    Position
}

// RecordExpr represents a record construction: {id1 = expr1, ..., idn = exprn}
type RecordExpr struct {
	Fields []RecordField
	Pos    Position
}

// RecordField is one field in a record construction
type RecordField struct {
	Label string
	Value Expr
	Pos   Position
}

// ProjectionExpr represents record projection: record.field
type ProjectionExpr struct {
	Record Expr
	Field  string
	Pos    Position
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
