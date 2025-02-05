package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func substract(x, y int) int {
	return x - y
}

func multiplie(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func show_pi() float32 {
	return math.Pi
}

func main() {
	fmt.Println("Hello User!")
	fmt.Println(add(1, 5))
	fmt.Println(show_pi())
	fmt.Println("Goodbye User!")
}
