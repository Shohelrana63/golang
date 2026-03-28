// STRUCT
// A struct is a collection of fields (variables) that can have different types.

package main

import "fmt"

type User struct {
	Name string //MEMBER VARIABLE OR PROPERTY
	Age  int
}

func main() {
	var user1 User
	user1 = User{ //INSTANCE
		Name: "Shohel",
		Age:  90,
	}
	fmt.Println(user1.Name)
	fmt.Println(user1.Age)

	user2 := User{
		Name: "Baig",
		Age:  80,
	}
	fmt.Println(user2.Name)
	fmt.Println(user2.Age)
}
