package main

import "fmt"

func main() {
  // unspecified array
  var x [3]int
  fmt.Println(x)

  // specified array with array literal
  var y = [3]int{10,20,30}
    fmt.Println(y)

}
