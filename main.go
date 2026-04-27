package main

import (
	"fmt"
	"magoito-compiler/internal/lexer"
	"magoito-compiler/internal/parser"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// 1. Read file
	bytes, err := os.ReadFile("example01.mag")
	if err != nil {
		panic(err)
	}

	input := antlr.NewInputStream(string(bytes))

	// 2. Create lexer
	lexer := lexer.NewMagoitoLexer(input)

	// 3. Token stream
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 4. Create parser
	parser := parser.NewMagoitoParser(stream)

	// Optional: better error messages
	parser.RemoveErrorListeners()
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// parser.SetErrorHandler(antlr.NewBailErrorStrategy())

	// 5. Parse starting rule
	tree := parser.Program()

	// 6. Print parse tree (LISP-style)
	fmt.Println(tree.ToStringTree(parser.GetRuleNames(), parser))
}
