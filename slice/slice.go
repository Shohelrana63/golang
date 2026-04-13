// SLICE
// A slice is a dynamically-sized, flexible view into the elements of an array.
// It is a reference type that provides a more powerful and convenient way to work with sequences of data compared to arrays.
// A slice does not store any data itself; it just describes a section of an underlying array.

package main

import "fmt"

func main() {
	arr := [6]string{"Go", "is", "a", "programming", "language", "!"}
	slice := arr[1:4]
	fmt.Println(slice) //[is a programming]

	slice1 := slice[1:3]
	fmt.Println(slice1) //[a programming]

	slice2 := slice1[0:1]
	fmt.Println(slice2)      //[a]
	fmt.Println(len(slice2)) //1
	fmt.Println(cap(slice2)) //3

}
