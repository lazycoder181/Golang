package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press enter when ready."

func main() {

	//seed your random number generator
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	playthegame(firstNumber, secondNumber, subtraction, answer)

}

func playthegame(firstNumber, secondNumber, subtraction, answer int) {
	//create reader variable
	reader := bufio.NewReader(os.Stdin)

	//display welcome instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take them through the game

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply the second number by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')
	//give them the answer
	//logic
	//answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}
