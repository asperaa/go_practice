// Internals of Slices: https://blog.golang.org/go-slices-usage-and-internals
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printValue("a", a)

	b := make([]int, 0, 5)
	printValue("b", b)

	c := b[:2]
	printValue("c", c)

	d := b[2:5]
	printValue("d", d)
}

func printValue(l string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", l, len(s), cap(s), s)
}