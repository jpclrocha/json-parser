package main

type Property struct {
	Type  string
	Key   any
	Value any
}

type Element struct {
	Type  string
	Value *any
}
type Expression struct {
	Type       string
	Properties *[]Property
	Elements   *[]Element
	Value      any
}

type Ast struct {
	Type string
	Body []Expression
}

func walk(currIndex *int, tokens []Token) Expression {
	token := tokens[*currIndex]
	(*currIndex)++

	if token.Type == LEFT_BRACE {
		(*currIndex)++
		token = tokens[*currIndex]

		properties := []Property{}
		node := Expression{
			Type:       "Object",
			Properties: &properties,
		}

		for token.Type != RIGHT_BRACE {
			property := Property{
				Type:  "Property",
				Key:   token,
				Value: nil,
			}

			(*currIndex)++
			token = tokens[*currIndex]

			property.Value = walk(currIndex, tokens)
			*node.Properties = append(*node.Properties, property)

			token = tokens[*currIndex]
			if token.Type == COMMA {
				(*currIndex)++
				token = tokens[*currIndex]
			}
		}
		(*currIndex)++
		return node
	}
	if token.Type == RIGHT_BRACE {
		(*currIndex)++
		properties := []Property{}
		return Expression{
			Type:       "Object",
			Properties: &properties,
		}
	}
	if token.Type == LEFT_BRACKET {
		(*currIndex)++
		token = tokens[*currIndex]

		elements := []Element{}

		node := Expression{
			Type:     "Array",
			Elements: &elements,
		}

		for token.Type != RIGHT_BRACKET {
			*node.Elements = append(*node.Elements, *walk(currIndex, tokens).Elements...)
			token = tokens[*currIndex]

			if token.Type == COMMA {
				(*currIndex)++
				token = tokens[*currIndex]
			}
		}
		(*currIndex)++
		return node
	}
	if token.Type == STRING {
		(*currIndex)++
		return Expression{
			Type:  "String",
			Value: token.Value,
		}
	}
	if token.Type == NUMBER {
		(*currIndex)++
		return Expression{
			Type:  "Number",
			Value: token.Value,
		}
	}

	if token.Type == TRUE {
		(*currIndex)++
		return Expression{
			Type:  "Boolean",
			Value: true,
		}
	}

	if token.Type == FALSE {
		(*currIndex)++
		return Expression{
			Type:  "Boolean",
			Value: false,
		}
	}

	if token.Type == NULL {
		(*currIndex)++
		return Expression{
			Type:  "Null",
			Value: nil,
		}
	}

	panic(token.Type)
}

func Parse(tokens []Token) Ast {
	if len(tokens) < 1 {
		panic("JSON invalido! Tokens = 0")
	}

	currIndex := 0
	body := []Expression{}
	ast := Ast{
		Type: "Program",
		Body: body,
	}

	for currIndex < len(tokens) {
		ast.Body = append(ast.Body, walk(&currIndex, tokens))
		currIndex++
	}

	return ast
}
