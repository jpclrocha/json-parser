package main

import (
	utils "json-parser/utils"
)

const (
	COMMA        = ","
	COLON        = ":"
	LEFTBRACKET  = "["
	RIGHTBRACKET = "]"
	LEFTBRACE    = "{"
	RIGHTBRACE   = "}"
	QUOTE        = `"`
)

func Lexer(input string) (tokens_vals []string) {
	whitespaceChars := []string{" ", "\t", "\b", "\n", "\r"}
	syntaxChars := []string{COMMA, COLON, LEFTBRACKET, RIGHTBRACKET, LEFTBRACE, RIGHTBRACE, QUOTE}
	tokens := []string{}

	for i := 0; i < len(input); i++ {
		json_string, _ := utils.LexString(input)
		if json_string != nil {
			tokens = append(tokens, *json_string)
			continue
		}

		json_number, _ := utils.LexNumber(input)
		if json_number != nil {
			tokens = append(tokens, *json_number)
			continue
		}

		json_float, _ := utils.LexFloat(input)
		if json_float != nil {
			tokens = append(tokens, *json_float)
			continue
		}

		json_bool, _ := utils.LexBoolean(input)
		if json_bool != nil {
			tokens = append(tokens, *json_bool)
			continue
		}

		json_null, _ := utils.LexNull(input)
		if json_null != nil {
			tokens = append(tokens, *json_null)
			continue
		}

		var errFlag *bool = nil

		for j := 0; j < len(whitespaceChars); j++ {
			if string(input[0]) == whitespaceChars[j] {
				input = input[1:]
				err := false
				errFlag = &err
			}
		}

		for j := 0; j < len(syntaxChars); j++ {
			if string(input[0]) == syntaxChars[j] {
				tokens = append(tokens, string(input[0]))
				input = input[1:]
				err := false
				errFlag = &err
			}
		}

		if errFlag == nil {
			panic("Caracter nao esperado!")
		}
	}

	return tokens
}
