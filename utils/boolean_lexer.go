package utils

const (
	TRUE_LEN  = len("true")
	FALSE_LEN = len("false")
)

func LexBoolean(input string) (json_bool *string, input_text string) {
	inputLen := len(input)

	if inputLen >= TRUE_LEN && input[:TRUE_LEN] == "true" {
		val := "true"
		return &val, input[TRUE_LEN:]
	} else if inputLen >= FALSE_LEN && input[:FALSE_LEN] == "false" {
		val := "false"
		return &val, input[FALSE_LEN:]
	}

	return nil, input
}
