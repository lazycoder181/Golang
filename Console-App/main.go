package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		//if keyboard.Open() package does not open
		log.Fatal(err)

	}
	defer func() {
		_ = keyboard.Close()
	}() //close the keyboard as soon as the main function ends using defer

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Mocha"
	coffees[4] = "Espresso"
	coffees[5] = "Macchiato"
	coffees[6] = "Americano" 

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("Press the respective numbers for your order or press Q to exit")
	fmt.Println("1-- Cappucino")
	fmt.Println("2-- Latte")
	fmt.Println("3-- Mocha")
	fmt.Println("4-- Espresso")
	fmt.Println("5-- Macchiato")
	fmt.Println("6-- Americano")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("you chose %s", coffees[i]) //Sprintf for loom it is %q, for int it is %d, for string it is %s and for boolean it is %t.
		fmt.Println(t)
	}

	fmt.Println("Cheers! have a good one!")
	fmt.Println("Program exiting...")
}
