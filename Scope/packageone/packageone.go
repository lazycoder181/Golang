package packageone

import "fmt"

var PackageVar = "I am a packge level variable from packageone"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
