// Pointer
// pointer is a variable that holds the memory address of another variable
// pointer is used to pass the reference of a variable to a function
// pointer is used to modify the value of a variable in a function

package main

import "fmt"

func main() {
	a := 90
	p := &a // & is the address operator, it gives the memory address of variable a
	fmt.Println("Value of a: ", a)
	fmt.Println("Address of a: ", p)
	fmt.Println("Value at address p: ", *p) // * is the dereference operator, it gives the value at the memory address
	*p = 100                                // modifying the value at address p, which modifies the value of a
	fmt.Println("Value of a after modification: ", a)
}
