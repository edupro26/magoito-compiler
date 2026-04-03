package main

import (
	"magoito-compiler/internal/lexer"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./example02.mag")

	tokens := lexer.Tokenize(string(bytes))
	for _, token := range tokens {
		token.Debug()
	}
}
