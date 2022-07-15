package utils

func Sum(numbers []int) int {
	sum := 0

	for n := range numbers {
		sum += n
	}

	return sum
}
