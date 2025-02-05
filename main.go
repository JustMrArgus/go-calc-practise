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

func square(x int) float64 {
	if x < 0 {
		return -1
	}
	return math.Sqrt(float64(x))
}

func main() {
	fmt.Println("Welcome to GoCalc!")
	for {
		fmt.Println("Choose your option:")
		fmt.Println("1 - Add")
		fmt.Println("2 - Substract")
		fmt.Println("3 - Multiplie")
		fmt.Println("4 - Divide")
		fmt.Println("5 - Square root")
		fmt.Println("")

		var option int
		fmt.Println("Enter number of option: ")
		fmt.Scan(&option)

		var num1 int
		var num2 int
		fmt.Println("Enter first number: ")
		fmt.Scan(&num1)
		fmt.Println("Enter second number: ")
		fmt.Scan(&num2)

		switch option {
		case 1:
			fmt.Println("")
			fmt.Println(add(num1, num2))
			fmt.Println("")
		case 2:
			fmt.Println("")
			fmt.Println(substract(num1, num2))
			fmt.Println("")
		case 3:
			fmt.Println("")
			fmt.Println(multiplie(num1, num2))
			fmt.Println("")
		case 4:
			fmt.Println("")
			fmt.Println(divide(num1, num2))
			fmt.Println("")
		case 5:
			fmt.Println("")
			fmt.Println(square(num1))
			fmt.Println("")
		default:
			fmt.Println("")
			fmt.Println("Invalid number of option, choose another one")
			fmt.Println("")
		}
	}
}
