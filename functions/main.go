package main

import "fmt"

func main() {

	myTotal := sumMany(1, 2, 3, 4, 5, 67, 88)
	fmt.Println(myTotal)
}

func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}

	return total
}
