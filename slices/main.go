package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "mouse")
	animals = append(animals, "lion")
	animals = append(animals, "tiger")

	for i, x := range animals {
		fmt.Println(i, x)
	}

	//Print first two elements
	fmt.Println(animals[0:2])

	//to find if a sice is sorted
	fmt.Println("Are the elements sorted?", sort.StringsAreSorted(animals))

	//sort them
	sort.Strings(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)

}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a

}
