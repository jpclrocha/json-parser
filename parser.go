package main

import "fmt"

func Parser(tokens []Token) Node {
	if len(tokens) < 1 {
		panicStr := fmt.Sprintf("Quantidade de tokens invalido! Tokens totais: %d", len(tokens))
		panic(panicStr)
	}
	current := 0

	var walk func() interface{}
	walk = func() interface{} {
		token := tokens[current]

		switch token.Type {
		case LEFT_BRACE:
			current++
			node := Node{
				Type:       "ObjectExpression",
				Properties: []Property{},
			}

			for tokens[current].Type != RIGHT_BRACE {
				// Verify key is a string
				if tokens[current].Type != STRING {
					panic("Object keys must be strings")
				}

				property := Property{
					Type: "Property",
					Key:  tokens[current],
				}
				current++

				// Expect colon
				if tokens[current].Type != COLON {
					panic("Expected ':' after object key")
				}
				current++

				property.Value = walk()
				node.Properties = append(node.Properties, property)

				// Check for valid comma placement
				if tokens[current].Type == COMMA {
					current++
					// Ensure not a trailing comma
					if tokens[current].Type == RIGHT_BRACE {
						panic("Trailing comma is not allowed in JSON")
					}
				}
			}
			current++
			return node

		case LEFT_BRACKET:
			current++
			node := Node{
				Type:     "ArrayExpression",
				Elements: []interface{}{},
			}

			for tokens[current].Type != RIGHT_BRACKET {
				node.Elements = append(node.Elements, walk())

				// Check for valid comma placement
				if tokens[current].Type == COMMA {
					current++
					// Ensure not a trailing comma
					if tokens[current].Type == RIGHT_BRACKET {
						panic("Trailing comma is not allowed in JSON")
					}
				}
			}
			current++
			return node

		case STRING:
			current++
			return Node{
				Type:  "StringLiteral",
				Value: token.Value,
			}

		case NUMBER:
			current++
			return Node{
				Type:  "NumberLiteral",
				Value: token.Value,
			}

		case TRUE:
			current++
			return Node{
				Type:  "BooleanLiteral",
				Value: true,
			}

		case FALSE:
			current++
			return Node{
				Type:  "BooleanLiteral",
				Value: false,
			}

		case NULL:
			current++
			return Node{
				Type:  "NullLiteral",
				Value: nil,
			}

		default:
			panic(fmt.Sprintf("Unexpected token type: %v", token.Type))
		}
	}

	ast := Node{
		Type:     "Program",
		Elements: []interface{}{},
	}

	for current < len(tokens) {
		ast.Elements = append(ast.Elements, walk())
	}

	return ast
}
