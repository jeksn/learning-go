package main

import (
  "fmt"
  "math/rand"
)

const passwordLength = 15

func randomChar() rune {
  chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789%-_+#"
  return rune(chars[rand.Intn(len(chars))])
}

func main(){
  var password string 
  for i := 0; i < passwordLength; i++ {
    password += string(randomChar())
  }
  fmt.Println(password)
}
