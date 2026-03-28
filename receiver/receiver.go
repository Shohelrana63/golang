// STRUCT
// A struct is a collection of fields (variables) that can have different types.

package main

import "fmt"

type User struct {
	Name string //MEMBER VARIABLE OR PROPERTY
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Ager: ", usr.Age)
}

//RECEIVER FUNCTION

func (usr User) printDetails() {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Ager: ", usr.Age)
}

func (usr User) printCallFun(year int) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Year: ", year)
}
func main() {
	var user1 User
	user1 = User{ //INSTANCE
		Name: "Shohel",
		Age:  90,
	}
	// printUserDetails(user1)
	user1.printDetails()

	user2 := User{
		Name: "Baig",
		Age:  80,
	}
	printUserDetails(user2)

	user1.printCallFun(2000)
}
