package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("../tests/step1/invalid.json")
	if err != nil {
		panic(err)
	}
	tokens := Lexer(string(input))

	jsonast := Parser(tokens)
	fmt.Printf(" %+v\n", jsonast)
	fmt.Println("JSON valido!")
	os.Exit(0)
}
