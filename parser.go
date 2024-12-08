package main

import (
	"fmt"
	"reflect"
)

func parseArray(tokens []string) (json_arr []string, res_tokens []string) {
	jsonArray := []string{}
	firstToken := tokens[0]

	if firstToken == RIGHTBRACKET {
		return jsonArray, tokens[1:]
	}

	for {
		_, json_array, _, tokens := Parse(tokens)
		jsonArray = append(jsonArray, json_array...)

		firstToken := tokens[0]
		if firstToken == RIGHTBRACKET {
			return jsonArray, tokens[1:]
		} else if firstToken != COMMA {
			panic("Expected comma after object in array")
		} else {
			tokens = tokens[1:]
		}
	}
}

func parseObject(tokens []string) (json_obj map[string]string, res_tokens []string) {
	jsonObject := make(map[string]string)

	firstToken := tokens[0]
	if firstToken == RIGHTBRACE {
		return jsonObject, tokens[1:]
	}

	for {
		jsonKey := tokens[0]
		if reflect.TypeOf(jsonKey).Elem().Kind() == reflect.String {
			tokens = tokens[1:]
		} else {
			panic("Expected string key, got another else")
		}

		if tokens[0] != COLON {
			panic("Expected colon after key in object, got another else")
		}

		json_value, _, _, tokens := Parse(tokens[1:])

		jsonObject[jsonKey] = *json_value
		firstToken = tokens[0]
		fmt.Print(firstToken)
		if firstToken == RIGHTBRACE {
			return jsonObject, tokens[1:]
		} else if firstToken != COMMA {
			panic("Expected comma after pair in object")
		}
		tokens = tokens[1:]
	}
}

func Parse(tokens []string) (json_value *string, json_array []string, json_object map[string]string, tokens_return []string) {
	firstToken := tokens[0]

	if firstToken == LEFTBRACKET {
		json_arr, res_tokens := parseArray(tokens[1:])
		return nil, json_arr, nil, res_tokens
	} else if firstToken == LEFTBRACE {
		json_object, res_tokens := parseObject(tokens[1:])
		return nil, nil, json_object, res_tokens
	} else {
		return &firstToken, nil, nil, tokens[1:]
	}

}
