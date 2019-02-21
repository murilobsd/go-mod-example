package main

import "fmt"

var x int
var y float64

func main() {
	x = 42
	y = 42.123
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x type = %T\n", x)
	fmt.Printf("y type = %T\n", y)
}
