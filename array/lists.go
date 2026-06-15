package main

import "fmt"

type product struct {
	title string
	id    string
	price float64
}

type Book struct {
	Title  string
	Author string
}

func main() {

	var products [4]product

	products[0] = product{
		title: "Simple note",
		id:    "1",
		price: 1.0,
	}
	prices := [4]float64{
		10.99, 7.88, 76.67, 858,
	}

	//slicing array
	slicedArray := products[0:3]
	//0 ->3 products[:3]  / [3:] convenient way for the slicing

	fmt.Printf(" product %v", products[0].id)
	fmt.Println(prices)
	fmt.Println(len(slicedArray), cap(slicedArray)) //length and capacity

	//dynamic length array
	ages := [5]float64{43, 56, 78, 56, 57}
	slicedAges := ages[0:3]
	fmt.Println(ages)
	updatedAges := append(slicedAges, 67)
	fmt.Println(updatedAges)

	// ... -> packing unpacking
	listofAges := [5]int{1, 2, 3, 4, 5}
	updatedListOfAges := append(listofAges[:3], listofAges[3:5]...)
	fmt.Println(updatedListOfAges)

	//the help of the make() function we can create the arrays/ maps .. data structures [dataType , length , max capacity]
	//using this we can prevent continuously reallocating memory
	numbers := make([]int, 3, 5)
	numbers[0] = 1
	numbers[1] = 2
	fmt.Println(numbers)

}
