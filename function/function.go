// function.go
// This program demonstrates the declaration and usage of functions in Go.
// It defines a function that adds two integers and returns the result, then calls the function and prints the result.

// package main

// import "fmt"

// func add(num1 int, num2 int) {
// 	sum := num1 + num2
// 	fmt.Println(sum)
// }

// func main() {
// 	a := 20
// 	b := 30

// 	// sum := a + b

// 	// fmt.Println(sum)
// 	add(a, b)
// }

// RETURNING VALUES FROM A FUNCTION
// package main

// import "fmt"

// func addNum(num1 int, num2 int) int {
// 	sum := num1 + num2
// 	return sum
// }

// func main() {
// 	a := 20
// 	b := 30
// 	sum := addNum(a, b)
// 	fmt.Println(sum)
// }

package main

import "fmt"

func addNum(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func main() {
	a := 20
	b := 30
	sum, mul := addNum(a, b)

	fmt.Println(sum)
	fmt.Println(mul)
}
