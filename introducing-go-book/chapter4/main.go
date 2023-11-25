package main

import "fmt"

/* one way of doing it
func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}
*/
/* another way of doing it
func main() { for i := 1; i<=10; i++{
	fmt.Println(i)
}
}
*/
/* even and odd numbers
 */
func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
