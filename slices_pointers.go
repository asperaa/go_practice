// A slice does not store any data.
// It just describes the section of an array.
// Changing the elements of a slice modifies the 
// the corresponding elements of its underlying array

package main

import "fmt"

func main(){
	var color = [4] string{
		"Blue",
		"Green",
		"Yellow",
		"Orange",
	}

	fmt.Println(color)
	a := color[0:2]
	b := color[2:4]

	b[1] = "Changed, Haha"
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(color)
}

