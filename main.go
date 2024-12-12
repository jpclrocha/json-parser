package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Print("Digite o caminho do arquivo: ")
	fmt.Scanf("%s", &path)
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	tokens := Lexer(string(input))

	jsonast := Parser(tokens)
	fmt.Printf(" %+v\n", jsonast)
	fmt.Println("JSON valido!")
	fmt.Println("exit status 0")
	os.Exit(0)
}
