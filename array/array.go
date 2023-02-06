package array

func SumArray(input []int) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		sum = sum + input[i]
	}
	return sum
}
