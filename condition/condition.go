// This program demonstrates the use of conditional statements in Go.
// It checks the value of variable 'a' and prints a message based on its value.

package main

import "fmt"

func main() {
	// a := 20
	// if a > 10 {
	// 	println("a is greater than 10")
	// } else if a == 10 {
	// 	println("a is equal to 10")
	// } else {
	// 	println("a is less than 10")
	// }

	// isResult := true
	// if isResult {
	// 	fmt.Println("Result is true")
	// } else {
	// 	fmt.Println("Result is false")
	// }

	a := 15
	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2, 3:
		fmt.Println("a is 2 or 3")
	default:
		fmt.Println("a is something else")
	}
}
