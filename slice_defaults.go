// To use defaults in slicing omit the high and low bounds
package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}
	array = array[1:]
	fmt.Println(array)

	array = array[:1]
	fmt.Println(array)

	array = array[1:2]
	fmt.Println(array)

}