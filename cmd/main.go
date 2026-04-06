package main

import (
	"fmt"
	"magoito-compiler/cmd/printer"
	"magoito-compiler/internal/lexer"
	"magoito-compiler/internal/parser"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./example01.mag")

	tokens, err := lexer.Tokenize(string(bytes))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	program, err := parser.Parse(tokens)
	if err != nil {
		fmt.Println("syntax error:", err)
		os.Exit(1)
	}
	fmt.Printf("parsed %d top-level declaration(s)\n\n", len(program.Declarations))
	printer.Print(program)
}
