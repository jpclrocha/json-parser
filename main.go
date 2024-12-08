package main

import (
	"fmt"
	"os"
)

func main() {
	dat, _ := os.ReadFile("./tests/step2/valid.json")
	tokens := Lexer(string(dat))

	if len(tokens) < 1 {
		fmt.Print("JSON invalido! O programa recebeu um JSON vazio!")
	}

	for i := 0; i < len(tokens); i++ {
		fmt.Printf("%s", tokens[i])
	}
	// Parse(tokens)
}
