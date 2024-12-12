package main

import (
	"fmt"
	"os"
)

func Parser(tokens []Token) Node {
	if len(tokens) < 1 {
		// panicStr := fmt.Sprintf("Quantidade de tokens invalido! Tokens totais: %d", len(tokens))
		// panic(panicStr)
		fmt.Printf("Quantidade de tokens invalido! Tokens totais: %d\n", len(tokens))
		os.Exit(1)
	}
	current := 0

	var walk func() interface{}
	walk = func() interface{} {
		token := tokens[current]

		switch token.Type {
		case LEFT_BRACE:
			current++
			node := Node{
				Type:       "Objeto",
				Properties: []Property{},
			}

			for tokens[current].Type != RIGHT_BRACE {
				if tokens[current].Type != STRING {
					// panic("As chaves do JSON precisam ser textos!")
					fmt.Println("As chaves do JSON precisam ser textos!")
					os.Exit(1)
				}

				property := Property{
					Type: "Propriedade",
					Key:  tokens[current],
				}
				current++

				// Expect colon
				if tokens[current].Type != COLON {
					// panic("Dois pontos (:) eh necessario depois de uma chave dentro de um objeto!")
					fmt.Println("Dois pontos (:) eh necessario depois de uma chave dentro de um objeto!")
					os.Exit(1)
				}
				current++

				property.Value = walk()
				node.Properties = append(node.Properties, property)

				// Checa se tem mais alguma propriedade objeto separado por virgula
				if tokens[current].Type == COMMA {
					current++
					// Nao pode ser uma virgula antes do final do objeto ex: {"chave":"valor",}
					if tokens[current].Type == RIGHT_BRACE {
						// panic("Virgula nao eh permitida no final de um objeto!")
						fmt.Println("Virgula nao eh permitida no final de um objeto!")
						os.Exit(1)
					}
				}
			}
			current++
			return node

		case LEFT_BRACKET:
			current++
			node := Node{
				Type:     "Array",
				Elements: []interface{}{},
			}

			for tokens[current].Type != RIGHT_BRACKET {
				node.Elements = append(node.Elements, walk())

				// Checa se tem mais algum valor na array
				if tokens[current].Type == COMMA {
					current++
					// Nao pode ser uma virgula antes do final da array ex: [{"chave":"valor"},]
					if tokens[current].Type == RIGHT_BRACKET {
						// panic("Virgula nao eh permitida no final de uma array!")
						fmt.Println("Virgula nao eh permitida no final de uma array!")
						os.Exit(1)
					}
				}
			}
			current++
			return node

		case STRING:
			current++
			return Node{
				Type:  "String",
				Value: token.Value,
			}

		case NUMBER:
			current++
			return Node{
				Type:  "Numero",
				Value: token.Value,
			}

		case TRUE:
			current++
			return Node{
				Type:  "Boolean",
				Value: true,
			}

		case FALSE:
			current++
			return Node{
				Type:  "Boolean",
				Value: false,
			}

		case NULL:
			current++
			return Node{
				Type:  "Null",
				Value: nil,
			}

		default:
			// panic(fmt.Sprintf("Token nao esperado: %v", token.Type))
			fmt.Printf("Token nao esperado: %v\n", token.Type)
			os.Exit(1)
			return nil
		}
	}

	ast := Node{
		Type:     "Programa",
		Elements: []interface{}{},
	}

	for current < len(tokens) {
		ast.Elements = append(ast.Elements, walk())
	}

	return ast
}
