package main

import "fmt"

// this will error
// func main() {
// 	var x map[string]int
// 	x["key"] = 10
// 	fmt.Println(x)
// }

// maps have to be initialized before they can be used.
func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}
