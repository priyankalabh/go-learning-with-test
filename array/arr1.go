package main

import "fmt"

func main() {
	arr := [5]int{2, 5, 10, 12, 15}

	for i := 0; i < 5; i++ {
		fmt.Printf("%v \t", arr[i])
	}
}
