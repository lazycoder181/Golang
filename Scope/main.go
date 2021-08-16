package main

import (
	"myapp/packageone"
)

var myVar = "Package level variable in main"

func main() {
	var blockVar = "Block level variable in main function."

	packageone.PrintMe(myVar, blockVar)

}
