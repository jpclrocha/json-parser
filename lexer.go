package main

import (
	"fmt"
	"regexp"
)

type TokenType string

const (
	LEFT_BRACE    TokenType = "LEFT_BRACE"
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"
	LEFT_BRACKET  TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"
	COLON         TokenType = "COLON"
	COMMA         TokenType = "COMMA"
	STRING        TokenType = "STRING"
	NUMBER        TokenType = "NUMBER"
	TRUE          TokenType = "TRUE"
	FALSE         TokenType = "FALSE"
	NULL          TokenType = "NULL"
)

type Token struct {
	Type  TokenType
	Value interface{}
}

type Node struct {
	Type       string
	Properties []Property
	Elements   []interface{}
	Value      interface{}
}

type Property struct {
	Type  string
	Key   Token
	Value interface{}
}

func createToken(tokenType TokenType, value ...interface{}) Token {
	if len(value) > 0 {
		return Token{Type: tokenType, Value: value[0]}
	}
	return Token{Type: tokenType}
}

func Lexer(input string) []Token {
	tokens := []Token{}
	current := 0

	for current < len(input) {
		char := string(input[current])

		switch char {
		case "{":
			tokens = append(tokens, createToken(LEFT_BRACE))
			current++
			continue
		case "}":
			tokens = append(tokens, createToken(RIGHT_BRACE))
			current++
			continue
		case "[":
			tokens = append(tokens, createToken(LEFT_BRACKET))
			current++
			continue
		case "]":
			tokens = append(tokens, createToken(RIGHT_BRACKET))
			current++
			continue
		case ":":
			tokens = append(tokens, createToken(COLON))
			current++
			continue
		case ",":
			tokens = append(tokens, createToken(COMMA))
			current++
			continue
		}

		// Checa se tem espacos em branco
		if matched, _ := regexp.MatchString(`\s`, char); matched {
			current++
			continue
		}

		// Checa se eh numero e pega o numero inteiro ex: 101, vai identificar o 1 loopar ate o final do numero, formando 101
		if matched, _ := regexp.MatchString(`[0-9]`, char); matched {
			value := ""
			for current < len(input) && matched {
				value += string(input[current])
				current++
				if current < len(input) {
					matched, _ = regexp.MatchString(`[0-9]`, string(input[current]))
				}
			}
			tokens = append(tokens, createToken(NUMBER, value))
			continue
		}

		// Strings
		if char == `"` {
			current++
			value := ""
			for current < len(input) && string(input[current]) != `"` {
				value += string(input[current])
				current++
			}
			tokens = append(tokens, createToken(STRING, value))
			current++
			continue
		}

		// chechar se eh true, false ou null
		if current+3 < len(input) {
			text := input[current : current+4]
			switch text {
			case "true":
				tokens = append(tokens, createToken(TRUE))
				current += 4
				continue
			case "fals":
				if current+4 < len(input) && input[current+4] == 'e' {
					tokens = append(tokens, createToken(FALSE))
					current += 5
					continue
				}
			case "null":
				tokens = append(tokens, createToken(NULL))
				current += 4
				continue
			}
		}

		panic(fmt.Sprintf("Nao conheco esse caractere: %s", char))
	}

	return tokens
}
