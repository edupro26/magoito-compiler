package parser

import (
	"fmt"
	"magoito-compiler/internal/lexer"

	"github.com/antlr4-go/antlr/v4"
)

type syntaxError struct {
	line, col int
	msg       string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error: %s", e.line, e.col, e.msg)
}

type magoitoErrorListener struct {
	*antlr.DefaultErrorListener
	err error
}

func newLynxErrorListener() *magoitoErrorListener {
	return &magoitoErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (l *magoitoErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol any,
	line, col int,
	msg string,
	ex antlr.RecognitionException,
) {
	panic(&syntaxError{line, col, msg})
}

func ParseProgram(input string) (tree IProgramContext, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rErr, ok := r.(*syntaxError); ok {
				tree, err = nil, rErr
			} else {
				panic(r)
			}
		}
	}()

	stream := antlr.NewInputStream(input)
	lex := lexer.NewMagoitoLexer(stream)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(newLynxErrorListener())

	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := NewMagoitoParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(newLynxErrorListener())

	return p.Program(), nil
}
