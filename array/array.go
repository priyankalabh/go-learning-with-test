package array

func SumArray(input []int) int {
	sum := 0

	for i := 0; i < len(input); i++ {
		sum = sum + input[i]
	}
	return sum
}

func ProdArray(input []int, multiplier int) []int {
	products := make([]int, len(input))

	for i := 0; i < len(input); i++ {

		products[i] = multiplier * input[i]
	}
	return products
}
