package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	doubleNum := transformNumber(&numbers, getTransformerFunction(&numbers))
	fmt.Println(doubleNum)
	tripleNum := transformNumber(&numbers, triple)
	fmt.Println(tripleNum)
}

func transformNumber(numbers *[]int, transformFunc transformFunc) []int {
	var transNumbers []int
	for _, val := range *numbers {
		transNumbers = append(transNumbers, transformFunc(val))
	}
	return transNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// return the function instead of value
func getTransformerFunction(number *[]int) transformFunc {
	if (*number)[0] == 1 {
		return double
	}
	return triple
}

// anonymouse function
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
