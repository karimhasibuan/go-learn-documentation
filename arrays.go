package main

import (
	"fmt"

)

func main() {
	var theArray [3]string
	theArray[0] = "Canada" // initialize the first element
	theArray[1] = "USA"   // initialize the second element
	theArray[2] = "Mexico" // initialize the third element

	fmt.Println(theArray[0]) // print the first element of theArray array
	fmt.Println(theArray[1]) // print the second element of theArray array
	fmt.Println(theArray[2]) // print the third element of theArray array
}