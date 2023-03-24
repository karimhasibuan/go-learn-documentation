package main

import "fmt"

// Global variable declaration
var (
	m int
	n int
)

func main() {
	var x int = 10 // Local variable declaration with Integer Type
	var y int // Local variable declaration with Integer Type without initialization

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	var a, b, c = 5.25, 6.25, "Hello" // Local variable declaration with multiple variables
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	city := "New York" // Local variable declaration with type inference
	country := "Queens" // Local variable declaration with type inference
	fmt.Println("city = ", city)
	fmt.Println("country = ", country)
	m,n = 1,2
	fmt.Println("m = ", m)
	fmt.Println("n = ", n)
}