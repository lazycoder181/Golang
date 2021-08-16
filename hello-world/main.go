package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	//For reading the text
	reader := bufio.NewReader(os.Stdin)

	whattosay := doctor.Intro()
	fmt.Println(whattosay)

	for {
		fmt.Print("-> ")
		userinput, _ := reader.ReadString('\n')

		userinput = strings.Replace(userinput, "\r\n", "", -1)

		if userinput == "quit" {
			break
		} else {
			response := doctor.Response(userinput)
			fmt.Println(response)
		}

	}

}
