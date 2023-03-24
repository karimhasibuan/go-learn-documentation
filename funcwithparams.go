package main

import "fmt"

// Function with parameters
func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	fmt.Println("Sum of 10 and 20 is", add(10, 20)) // Calling add function
}