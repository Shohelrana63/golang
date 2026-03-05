// This program demonstrates the use of conditional statements in Go.
// It checks the value of variable 'a' and prints a message based on its value.

package main

func main() {
	a := 20
	if a > 10 {
		println("a is greater than 10")
	} else if a == 10 {
		println("a is equal to 10")
	} else {
		println("a is less than 10")
	}
}
