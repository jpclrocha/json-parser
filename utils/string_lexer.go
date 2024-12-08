package utils

func LexString(input string) (json_str *string, input_text string) {
	jsonString := ""

	if string(input[0]) == `"` {
		input = input[1:]
	} else {
		return nil, input
	}

	for i := 0; i < len(input); i++ {
		if string(input[i]) == `"` {
			return &jsonString, string(input[len(jsonString)+1:])
		} else {
			jsonString += string(input[i])
		}
	}

	panic("Uma aspa era esperada com fechamento de texto!")
}
