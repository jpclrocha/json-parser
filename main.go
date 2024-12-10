package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("./tests/step4/valid.json")
	if err != nil {
		panic(err)
	}
	tokens := Lexer(string(input))

	Parse(tokens)
	fmt.Println("JSON valido!")
	os.Exit(0)
}
