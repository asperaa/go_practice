// A lice literal is like an array literal but
// without the length
package main

import "fmt"

func main() {
	
	fruits := []string{
		"Mango",
		"Apple",
		"Orange",
		"Pineapple",
	}

	fmt.Println(fruits)
	even := []bool{true, true, false, true}
	fmt.Println(even)

	struct_try := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{5, false},
	}

	fmt.Println(struct_try)



}