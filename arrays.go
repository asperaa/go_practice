package main

import "fmt"

func main() {
	
	// declare an array of strings
	var animal [3]string
	animal[0] = "cat"
	animal[1] = "dog"
	animal[2] = "elephant"

	fmt.Println(animal[1])

	// declare an array of even numbers and initialise it
	even := [5]int{2, 4, 6, 8, 10}
	fmt.Println(even)



}