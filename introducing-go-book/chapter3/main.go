package main

import "fmt"

//	func main() {
//		var x string = "Hello, World"
//		fmt.Println(x)
//	}
//
//	func main() {
//		var x string
//		x = "First"
//		fmt.Println(x)
//		x = "Second"
//		fmt.Println(x)
//	}
//
//	func main() {
//		x := "foo"
//		fmt.Println(x)
//	}

// func main() {
// 	x := 5
// 	x += 1
// 	fmt.Println(x)
// }

// Example from the book

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
