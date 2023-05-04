package main

import "fmt"

//	func main() {
//		var x [5]float64
//		x[0] = 98
//		x[1] = 93
//		x[2] = 77
//		x[3] = 82
//		x[4] = 83
//		var total float64 = 0
//		for i := 0; i < 5; i++ {
//			total += x[i]
//		}
//		fmt.Println(total / 5)
//	}
//

// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83
// 	var total float64 = 0
// 	for i := 0; i < len(x); i++ {
// 		total += x[i]
// 	}
// 	// fmt.Println(total / len(x)) // this will error
// 	fmt.Println(total / float64(len(x))) // this fixes the types
// }

func main() {
	x := [5]float64{
		98,
		93,
		77,
		82,
		83, // needs a closing ,
	}
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	// fmt.Println(total / len(x)) // this will error
	fmt.Println(total / float64(len(x))) // this fixes the types
}
