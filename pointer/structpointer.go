package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	user := User{
		Name:   "Shohel",
		Age:    28,
		Salary: 0,
	}

	p := &user      // p is a pointer to the user struct
	fmt.Println(*p) // dereferencing the pointer to get the value of the user struct

	fmt.Println("User Name: ", p.Name) // accessing the Name field using the pointer

}
