// An empty slice has a value zero
package main

import "fmt"

func main() {
	var empty []int
	fmt.Println(empty, len(empty), cap(empty))
	if empty == nil{
		fmt.Println("nil")
	}
}