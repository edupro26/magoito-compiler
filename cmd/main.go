package main

import (
	"fmt"
	"magoito-compiler/internal/lexer"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("./example02.mag")

	tokens, err := lexer.Tokenize(string(bytes))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range tokens {
		token.Debug()
	}
}
