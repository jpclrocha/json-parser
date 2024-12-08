package utils

const (
	NULL_LEN = len("null")
)

func LexNull(input string) (json_null *string, input_text string) {
	inputLen := len(input)
	if inputLen >= NULL_LEN && input[:NULL_LEN] == "null" {
		val := "true"
		return &val, input[NULL_LEN:]
	}

	return nil, input
}
