package main
import "fmt"

func add(num1, num2 float32) float32 {
	return num1 + num2
}

func multiply(a, b string) (string, string) {
	return a, b
}
func main() {

	var firstName string = "Aditya"
	var lastName string = "Ankur"
	
	var num3 int = 34
	var num5 float32 = float64(num3)
	var num4 float32 = 78
	fmt.Println(add(num4, num5))
	fmt.Println(multiply(firstName, lastName))


}
