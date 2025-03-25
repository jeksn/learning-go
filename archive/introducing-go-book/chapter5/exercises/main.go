package main

import "fmt"

func main() {

	//How do you access the fourth element of an array or slice?
	arr := []float64{1, 2, 3, 4, 5, 6, 7}
	z := arr[3] // gets first third element
	fmt.Println(z)

	// What is the length of a slice created using make([]int, 3, 9)?
	x := make([]int, 3, 9)
	fmt.Println(len(x))

	// Given the following array, what would y[2:5] give you?
	y := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(y[2:5])
}
