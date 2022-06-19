package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	for _, char := range input {
		output = string(char) + output
	}
	return output
}

