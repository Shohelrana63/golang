package main

import "fmt"

func main() {
	arr := [6]string{"Go", "is", "a", "programming", "language", "!"}
	slice := arr[1:4]
	fmt.Println(slice)
}
