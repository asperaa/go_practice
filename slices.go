// A slice is dynamically-sized flexible view into the elements of the array.
package main

import "fmt"

func main() {
	even := [6]int{2, 4, 6, 8, 10, 12}
	var slice []int = even[1: 4]
	fmt.Println(slice)
}