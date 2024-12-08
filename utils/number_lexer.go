package utils

func LexNumber(input string) (json_num *string, input_num string) {
	jsonNumber := ""

	numberChars := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", "E", "e"}

	for i := 0; i < len(input); i++ {
		for char := 0; char < len(numberChars); char++ {
			if char == i {
				jsonNumber += string(input[i])
			} else {
				break
			}
		}
	}

	rest := input[len(jsonNumber):]

	if len(jsonNumber) == 0 {
		return nil, input
	}

	return &jsonNumber, rest

}
