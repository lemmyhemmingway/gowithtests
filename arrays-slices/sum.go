package arraysslices

func SumAll(array ...[]int) []int {
	var sums []int

	for _, numbers := range array {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func Sum(array []int) int {
	var total int
	for _, v := range array {
		total += v
	}
	return total
}

func SumAllTails(array ...[]int) []int {
	var sums []int

	for _, numbers := range array {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
