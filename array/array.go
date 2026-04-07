//ARRAY
// ARRAYS are fixed-size collections of elements of the same type.
// In Go, arrays are defined with a specific size and type.
// Once an array is created, its size cannot be changed.
// Arrays can be indexed to access individual elements.

package main

import "fmt"

func main() {
	var arr [2]int
	arr[0] = 20
	arr[1] = 30
	fmt.Println(arr[0])
	fmt.Println(arr[1])

	arr1 := [3]int{10, 20, 30}
	arr2 := [3]string{"Go", "Python", "Java"}
	fmt.Println(arr1)
	fmt.Println(arr2)

}
