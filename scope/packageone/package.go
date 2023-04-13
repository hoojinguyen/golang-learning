package packageone

import "fmt"

var PackageVar = "This is Package Var in packageone"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, "\n", blockVar, "\n", PackageVar)
}
