package main


import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello"))	// using reverse from stringutil
	fmt.Println(stringutil.ToUpper("Hello"))  // using ToUpper from stringutil
}