package main

import (
	"fmt"
	"go-learning-with-test/array"
	"go-learning-with-test/calc"
	"go-learning-with-test/mystring"
	"log"
)

func main() {
	nums := []int{2, 5, 6, 3, 8}
	sum := array.SumArray(nums)
	fmt.Printf("sum of array numbers =%v", sum)
	result := []int{5, 2, 10}
	multiplier := 4
	finalProduct := array.ProdArray(result, multiplier)
	fmt.Printf("\n The Products of Array := %v", finalProduct)
	greetingMsg := mystring.Greetmsg("Morning", "Priyanshi")
	fmt.Printf("\n%v", greetingMsg)
	dividedValue, err := calc.Division(20, 0)
	if err != nil {
		log.Printf("\nerror= %v", err)
	}
	dividedValue, err = calc.Division(20, 5)
	if err != nil {
		log.Printf("\nerror= %v", err)
	}
	fmt.Printf("\nvalue= %v", dividedValue)
}
