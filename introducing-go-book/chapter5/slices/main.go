package main

import "fmt"

//	func main() {
//		x := make([]float64, 5)
//		fmt.Println(x)
//	}

// [low : high] expression:
// func main() {
// 	arr := []float64{1, 2, 3, 4, 5, 6, 7}
// 	x := arr[0:5] // gets first 5 elements
// 	fmt.Println(x)
// }

// append
// func main() {
// 	slice1 := []int{1, 2, 3}
// 	slice2 := append(slice1, 4, 5)
// 	fmt.Println(slice1, slice2)
// }

// copy
func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
