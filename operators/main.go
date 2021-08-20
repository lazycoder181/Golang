package main

import (
	"fmt"
	"math"
)

func main() {
	answer := 7*3 + 10
	fmt.Println("Answer is: ", answer)

	area_circle()
	display_month()

}

func area_circle() {
	radius := 12.0
	area := math.Pi * radius * radius
	fmt.Println("area of the circle is: ", area)
}

func display_month() {
	for month := 1; month <= 12; month++ {
		fmt.Println("The after", month, "is", month%12+1)
	}

}
