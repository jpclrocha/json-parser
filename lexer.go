package main

import (
	"fmt"
	"regexp"
)

const (
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"
	LEFT_BRACKET  = "LEFT_BRACKET"
	RIGHT_BRACKET = "RIGHT_BRACKET"
	COLON         = "COLON"
	COMMA         = "COMMA"
	STRING        = "STRING"
	NUMBER        = "NUMBER"
	TRUE          = "TRUE"
	FALSE         = "FALSE"
	NULL          = "NULL"
)

type Token struct {
	Type  string
	Value *string
}

func Lexer(input string) []Token {
	currIndex := 0
	tokens := []Token{}

	for currIndex < len(input) {
		char := string(input[currIndex])

		switch char {
		case "{":
			tokens = append(tokens, createToken(LEFT_BRACE, nil))
			currIndex++
			continue
		case "}":
			tokens = append(tokens, createToken(RIGHT_BRACE, nil))
			currIndex++
			continue
		case "[":
			tokens = append(tokens, createToken(LEFT_BRACKET, nil))
			currIndex++
			continue
		case "]":
			tokens = append(tokens, createToken(RIGHT_BRACKET, nil))
			currIndex++
			continue
		case ":":
			tokens = append(tokens, createToken(COLON, nil))
			currIndex++
			continue
		case ",":
			tokens = append(tokens, createToken(COMMA, nil))
			currIndex++
			continue

		// TESTANDO TODAS AS POSSIBILIDADES DE SER VAZIO -> " ", \t, \n \r \f \v
		case " ", "\t", "\n", "\r", "\f", "\v":
			currIndex++
			continue
		case `"`:
			value := ""
			currIndex++
			char := string(input[currIndex])
			for char != `"` {
				value += char
				currIndex++
				char = string(input[currIndex])
			}
			currIndex++
			tokens = append(tokens, createToken(STRING, &value))
			continue

		case "t":
			if string(input[currIndex+1]) == "r" && string(input[currIndex+2]) == "u" && string(input[currIndex+3]) == "e" {
				tokens = append(tokens, createToken(TRUE, nil))
				currIndex += 4
				continue
			}

		case "f":
			if string(input[currIndex+1]) == "a" && string(input[currIndex+2]) == "l" && string(input[currIndex+3]) == "s" && string(input[currIndex+4]) == "e" {
				tokens = append(tokens, createToken(FALSE, nil))
				currIndex += 5
				continue
			}

		case "n":
			if string(input[currIndex+1]) == "u" && string(input[currIndex+2]) == "l" && string(input[currIndex+3]) == "l" {
				tokens = append(tokens, createToken(NULL, nil))
				currIndex += 4
				continue
			}
		}

		if regexp.MustCompile("[0-9]").MatchString(char) {
			value := ""

			for {
				fmt.Print(value)
				fmt.Print(value)
				if regexp.MustCompile("[0-9]").MatchString(char) {
					value += char
					currIndex++
				}
				break
			}
			tokens = append(tokens, createToken(NUMBER, &value))
			continue
		}

		panicString := fmt.Sprintf("Caracter invalido: %s", char)
		panic(panicString)
	}

	return tokens
}

func createToken(tokenType string, value *string) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}
