package utils

func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func Add(a, b int) int {
	return a + b
}
