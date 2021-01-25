package arrays

import "strconv"

func NumberArray(length int) []int {
	var result []int

	for i := 0; i < length; i++ {
		result = append(result, i*10)
	}

	return result
}

func StringArray(length int) (result []string) {

	for i := 0; i < length; i++ {
		result = append(result, "String "+strconv.Itoa(i))
	}

	return result
}
