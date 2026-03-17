package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
	x = 70
	y = 30
)

func main() {
	fmt.Println("custom package")
	mathlib.Math(x, y)
}
