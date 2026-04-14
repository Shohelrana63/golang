// SLICE
// A slice is a dynamically-sized, flexible view into the elements of an array.
// It is a reference type that provides a more powerful and convenient way to work with sequences of data compared to arrays.
// A slice does not store any data itself; it just describes a section of an underlying array.

package main

import "fmt"

func main() {
	// arr := [6]string{"Go", "is", "a", "programming", "language", "!"}
	// slice := arr[1:4]
	// fmt.Println(slice) //[is a programming]

	// slice1 := slice[1:3]
	// fmt.Println(slice1) //[a programming]

	// slice2 := slice1[0:1]
	// fmt.Println(slice2)      //[a]
	// fmt.Println(len(slice2)) //1
	// fmt.Println(cap(slice2)) //3

	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println("slice", s,
	// 	"len", len(s),
	// 	"cap", cap(s)) // [1 2 3 4 5]

	// s := make([]int, 3)

	// fmt.Println("slice", s)
	// fmt.Println("len", len(s))
	// fmt.Println("cap", cap(s))

	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	fmt.Println(x)      //[1 2 3]
	fmt.Println(len(x)) //3
	fmt.Println(cap(x)) //4
	y := x
	fmt.Println(y) //[1 2 3]
	x = append(x, 4)
	y = append(y, 5)

	fmt.Println(x) //[1 2 3 5]
	fmt.Println(y) //[1 2 3 5]
}
