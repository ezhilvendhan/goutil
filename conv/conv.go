package conv

import (
	"strconv"
)

func ToIntArray(input []string) []int {
    length := len(input)
	output := make([]int, length)
	for idx := 0; idx < length; idx++  {
		intval, err := strconv.Atoi(input[idx])
		if err == nil {
			output[idx] = intval
		}
	}
	return output
}
