package main

import "fmt"

func print(numbers *[5]int) { // receiving a pointer to an array of 5 integers
	fmt.Println(numbers)
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	print(&arr) // passing the address of the array to the function
}
