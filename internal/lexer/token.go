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
	Kind  Kind
	Value string
}

func (t Token) isOneOfMany(tokens ...Kind) bool {
	for _, token := range tokens {
		if token == t.Kind {
			return true
		}
	}
	return false
}

func (t Token) Debug() {
	if t.isOneOfMany(IDENTIFIER, INT, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(t.Kind), t.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(t.Kind))
	}
}

func NewToken(kind Kind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

func TokenKindString(kind Kind) string {
	switch kind {
	case EOF:
		return "eof"
	case IDENTIFIER:
		return "ident"
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
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case COLON:
		return ":"
	case COMMA:
		return ","
	case DOT:
		return "."
	case ARROW:
		return "->"
	case ASSIGN:
		return "="
	case SEMICOLON:
		return ";"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case STAR:
		return "*"
	case SLASH:
		return "/"
	case PERCENT:
		return "%"
	case POWER:
		return "^"
	case NOT:
		return "!"
	case EQ:
		return "=="
	case NEQ:
		return "!="
	case LT:
		return "<"
	case LTE:
		return "<="
	case GT:
		return ">"
	case GTE:
		return ">="
	case AND:
		return "&&"
	case OR:
		return "||"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
