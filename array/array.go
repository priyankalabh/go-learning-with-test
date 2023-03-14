package array

import "errors"

func SumArray(input []int) (int, error) {
	sum := 0
	if len(input) == 0 {
		return 0, errors.New("Empty array")
	}
	for i := 0; i < len(input); i++ {
		sum = sum + input[i]
	}
	return sum, nil
}

func ProdArray(input []int, multiplier int) ([]int, error) {
	products := make([]int, len(input))
	if len(input) == 0 {
		return []int{}, errors.New("input required")
	}

	for i := 0; i < len(input); i++ {
		products[i] = multiplier * input[i]
	}
	return products, nil
}
