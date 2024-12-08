package utils

func LexFloat(input string) (json_num *string, input_num string) {
	jsonNumber := ""

	rest := input[len(jsonNumber):]

	if len(jsonNumber) == 0 {
		return nil, input
	}

	for i := 0; i < len(jsonNumber); i++ {
		if string(jsonNumber[i]) == "." {
			return &jsonNumber, rest
		}
	}

	return nil, input
}
