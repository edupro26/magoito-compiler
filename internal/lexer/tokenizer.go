package lexer

import (
	"errors"
	"fmt"
	"regexp"
)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	tokens   []Token
	source   string
	pos      int
	line     int
	col      int
}

type LexError struct {
	Message string
	Line    int
	Col     int
}

func (e *LexError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error: %s", e.Line, e.Col, e.Message)
}

type regexHandler func(lex *lexer, regex *regexp.Regexp)

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		line:   1,
		col:    1,
		source: source,
		tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`--.*`), skipHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`"[^"\n]*`), unclosedStringHandler},
			{regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_']*`), symbolHandler},
			{regexp.MustCompile(`0[0-9]+`), invalidIntegerHandler},
			{regexp.MustCompile(`[0-9]+\.[0-9]+`), invalidFloatHandler},
			{regexp.MustCompile(`0|[1-9][0-9]*`), integerHandler},
			{regexp.MustCompile(`_`), defaultHandler(WILDCARD, "_")},
			{regexp.MustCompile(`\{`), defaultHandler(LBRACE, "{")},
			{regexp.MustCompile(`}`), defaultHandler(RBRACE, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(LPAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(RPAREN, ")")},
			{regexp.MustCompile(`->`), defaultHandler(ARROW, "->")},
			{regexp.MustCompile(`==`), defaultHandler(EQ, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NEQ, "!=")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGN, "=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LTE, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LT, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GTE, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GT, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMICOLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(MINUS, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`\^`), defaultHandler(POWER, "^")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
		},
	}
}

func Tokenize(source string) (tokens []Token, err error) {
	lex := createLexer(source)
	defer func() {
		if r := recover(); r != nil {
			if errVal, ok := r.(error); ok {
				var le *LexError
				if errors.As(errVal, &le) {
					err = le
					return
				}
			}
			panic(r)
		}
	}()

	for !lex.atEof() {
		match := false
		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				match = true
				break
			}
		}
		if !match {
			return nil, &LexError{
				Message: fmt.Sprintf("illegal token '%v'", string(lex.source[lex.pos])),
				Line:    lex.line,
				Col:     lex.col,
			}
		}
	}
	lex.push(newToken(EOF, "eof", lex.line, lex.col))
	return lex.tokens, nil
}

func (lex *lexer) advN(n int) {
	for i := 0; i < n; i++ {
		if lex.source[lex.pos+i] == '\n' {
			lex.line++
			lex.col = 1
		} else {
			lex.col++
		}
	}
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.tokens = append(lex.tokens, token)
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) atEof() bool {
	return lex.pos >= len(lex.source)
}

func defaultHandler(kind Kind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		line, col := lex.line, lex.col
		lex.push(newToken(kind, value, line, col))
		lex.advN(len(value))
	}
}

func integerHandler(lex *lexer, regex *regexp.Regexp) {
	line, col := lex.line, lex.col
	match := regex.FindString(lex.remainder())
	lex.push(newToken(INT, match, line, col))
	lex.advN(len(match))
}

func invalidIntegerHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	panic(&LexError{
		Message: fmt.Sprintf("invalid integer literal '%v'", match),
		Line:    lex.line,
		Col:     lex.col,
	})
}

func invalidFloatHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	panic(&LexError{
		Message: fmt.Sprintf("invalid integer literal '%v'", match),
		Line:    lex.line,
		Col:     lex.col,
	})
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	line, col := lex.line, lex.col
	match := regex.FindStringIndex(lex.remainder())
	literal := lex.remainder()[match[0]+1 : match[1]-1]
	lex.push(newToken(STRING, literal, line, col))
	lex.advN(len(literal) + 2)
}

func unclosedStringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	panic(&LexError{
		Message: fmt.Sprintf("unclosed string literal '%v'", match),
		Line:    lex.line,
		Col:     lex.col,
	})
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	line, col := lex.line, lex.col
	match := regex.FindString(lex.remainder())

	if kind, exits := reservedMap[match]; exits {
		lex.push(newToken(kind, match, line, col))
	} else {
		lex.push(newToken(IDENTIFIER, match, line, col))
	}

	lex.advN(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advN(match[1])
}
