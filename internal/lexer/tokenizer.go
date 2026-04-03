package lexer

import (
	"fmt"
	"regexp"
)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

type regexHandler func(lex *lexer, regex *regexp.Regexp)

func Tokenize(source string) []Token {
	lex := createLexer(source)
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
			panic(fmt.Sprintf("lexer::Error - unrecnonized token near %s\n", lex.remainder()))
		}
	}
	lex.push(NewToken(EOF, "eof"))
	return lex.Tokens
}

func (lex *lexer) advN(n int) {
	lex.pos += n
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) atEof() bool {
	return lex.pos >= len(lex.source)
}

func defaultHandler(kind Kind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(INT, match))
	lex.advN(len(match))
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]
	lex.push(NewToken(STRING, stringLiteral))
	lex.advN(len(stringLiteral) + 2)
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	if kind, exits := reservedMap[match]; exits {
		lex.push(NewToken(kind, match))
	} else {
		lex.push(NewToken(IDENTIFIER, match))
	}
	lex.advN(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advN(match[1])
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`--.*`), skipHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`\{`), defaultHandler(LBRACE, "{")},
			{regexp.MustCompile(`}`), defaultHandler(RBRACE, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(LPAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(RPAREN, ")")},
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
