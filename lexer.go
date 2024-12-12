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

		// Whitespace
		if matched, _ := regexp.MatchString(`\s`, char); matched {
			current++
			continue
		}

		// Numbers
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
		if char == "\"" {
			current++
			value := ""
			for current < len(input) && string(input[current]) != "\"" {
				value += string(input[current])
				current++
			}
			tokens = append(tokens, createToken(STRING, value))
			current++
			continue
		}

		// Keywords: true, false, null
		if current+3 < len(input) {
			substr := input[current : current+4]
			switch substr {
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

		panic(fmt.Sprintf("Unknown character: %s", char))
	}

	return tokens
}
