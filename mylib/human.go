package mylib

func Average(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	total := 0
	for _, num := range numbers {
		total += num
	}

	return total / len(numbers)
}
