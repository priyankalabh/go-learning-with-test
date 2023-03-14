package main

import (
	"fmt"
	"go-learning-with-test/array"
	"go-learning-with-test/calc"
	"go-learning-with-test/mystring"
	"go-learning-with-test/utils"
	"log"
)

func main() {
	nums := []int{2, 5, 6, 3, 8}
	sum, err := array.SumArray(nums)
	if err != nil {
		log.Printf("SumArray error=%v", err)
	} else {
		fmt.Printf("sum of array numbers =%v", sum)

	}
	result := []int{5, 2, 10}
	multiplier := 4
	finalProduct, err := array.ProdArray(result, multiplier)
	if err != nil {
		log.Printf("\n Product array error=%v", err)
	} else {
		fmt.Printf("\n The Products of Array := %v", finalProduct)

	}
	greetingMsg, err := mystring.Greetmsg("Morning", "")
	if err != nil {
		log.Printf("\n greetings  error=%v", err)
	} else {
		fmt.Printf("\n%v", greetingMsg)

	}

	dividedValue, err := calc.Division(20, 0)
	if err != nil {
		log.Printf("\nerror= %v", err)
	}
	dividedValue, err = calc.Division(20, 5)
	if err != nil {
		log.Printf("\nerror= %v", err)
	}
	fmt.Printf("\nvalue= %v", dividedValue)

	minNum, err := utils.Min(4, 8)
	if err != nil {
		log.Printf("min error=%v", err)
	}
	fmt.Printf("min number=%v", minNum)
}
