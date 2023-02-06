package main

import (
	"fmt"
	"go-learning-with-test/array"
)

func main() {
	nums := []int{2, 5, 6, 3, 8}
	sum := array.SumArray(nums)
	fmt.Printf("sum of array numbers =%v", sum)
}
