package main

import "fmt"

func main() {

	factor := factorial(2)
	fmt.Println(factor)

}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// variadic function we can pass the multiple variables as input of func
func variadic(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number

	}
	return sum
}
