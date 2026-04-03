package lexer

import (
	"fmt"
)

const (
	EOF Kind = iota

	IDENTIFIER
	INT
	STRING
	TRUE
	FALSE
	UNIT

	CONST
	FUN
	VAR
	IF
	THEN
	ELSE
	WHILE
	DO
	PRINT

	WILDCARD // _

	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	COLON     // :
	COMMA     // ,
	DOT       // .
	ARROW     // ->
	ASSIGN    // =
	SEMICOLON // ;

	PLUS    // +
	MINUS   // -
	STAR    // *
	SLASH   // /
	PERCENT // %
	POWER   // ^
	NOT     // !

	EQ  // ==
	NEQ // !=
	LT  // <
	LTE // <=
	GT  // >
	GTE // >=

	AND // &&
	OR  // ||
)

var reservedMap = map[string]Kind{
	"var":   VAR,
	"const": CONST,
	"fun":   FUN,
	"print": PRINT,
	"unit":  UNIT,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"then":  THEN,
	"else":  ELSE,
	"while": WHILE,
	"do":    DO,
}

type Kind int

type Token struct {
	kind  Kind
	value string
}

func (t Token) Debug() {
	// TODO Better string representation
	if t.kind == IDENTIFIER || t.kind == INT || t.kind == STRING {
		fmt.Printf("%s (%s)\n", tokenToString(t.kind), t.value)
	} else {
		fmt.Printf("%s ()\n", tokenToString(t.kind))
	}
}

func newToken(kind Kind, value string) Token {
	return Token{
		kind:  kind,
		value: value,
	}
}

func tokenToString(kind Kind) string {
	switch kind {
	case EOF:
		return "eof"
	case IDENTIFIER:
		return "identifier"
	case INT:
		return "int"
	case STRING:
		return "string"
	case CONST:
		return "const"
	case FUN:
		return "fun"
	case VAR:
		return "var"
	case IF:
		return "if"
	case THEN:
		return "then"
	case ELSE:
		return "else"
	case WHILE:
		return "while"
	case DO:
		return "do"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case UNIT:
		return "unit"
	case PRINT:
		return "print"
	case WILDCARD:
		return "wildcard"
	case LPAREN:
		return "left_paren"
	case RPAREN:
		return "right_paren"
	case LBRACE:
		return "left_brace"
	case RBRACE:
		return "right_brace"
	case COLON:
		return "colon"
	case COMMA:
		return "comma"
	case DOT:
		return "dot"
	case ARROW:
		return "arrow"
	case ASSIGN:
		return "assign"
	case SEMICOLON:
		return "semi_colon"
	case PLUS:
		return "plus"
	case MINUS:
		return "minus"
	case STAR:
		return "star"
	case SLASH:
		return "slash"
	case PERCENT:
		return "percent"
	case POWER:
		return "power"
	case NOT:
		return "not"
	case EQ:
		return "equals"
	case NEQ:
		return "not_equals"
	case LT:
		return "less_than"
	case LTE:
		return "less_than_equals"
	case GT:
		return "greater_than"
	case GTE:
		return "greater_than_equals"
	case AND:
		return "and"
	case OR:
		return "or"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
