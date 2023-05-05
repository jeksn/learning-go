package main

import "fmt"

// this will error
// func main() {
// 	var x map[string]int
// 	x["key"] = 10
// 	fmt.Println(x)
// }

// maps have to be initialized before they can be used.
// func main() {
// 	x := make(map[string]int)
// 	x["key"] = 10
// 	fmt.Println(x["key"])
// }

// example of map
//
//	func main() {
//		elements := make(map[string]string)
//		elements["H"] = "Hydrogen"
//		elements["He"] = "Helium"
//		elements["Li"] = "Lithium"
//		elements["Be"] = "Beryllium"
//		elements["B"] = "Boron"
//		elements["C"] = "Carbon"
//		elements["N"] = "Nitrogen"
//		elements["O"] = "Oxygen"
//		elements["F"] = "Fluorine"
//		elements["Ne"] = "Neon"
//		//fmt.Println(elements["Li"])
//		// check for zero value
//		// name, ok := elements["Li"]
//		// fmt.Println(name, ok)
//		// if ok
//		if name, ok := elements["Li"]; ok {
//			fmt.Println(name, ok)
//		}
//	}
//
// shorter way to define map
func main() {
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name": "Lithium", "state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {"name": "Boron",
			"state": "solid",
		},
		"C": {"name": "Carbon", "state": "solid"},
	}
	if name, ok := elements["Li"]; ok {
		fmt.Println(name, ok)
	}
}
