package packageone

import "fmt"

/* var privatevar = "I am private"
var Publicvar = "I am public or exported"

func notExported() {

}

func Exported() {

}
*/
var PackageVar = "I am a package level variable in the packageone"

func PrintMe(s1, s2, PackageVar) {
	fmt.Println(s1, s2, PackageVar)
}
