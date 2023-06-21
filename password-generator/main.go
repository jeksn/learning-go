package main

import (
  "fmt"
  "math/rand"
  "time"
)

const passwordLength = 10

func main(){
  rand.Seed(time.Now().UnixNano())

  var password string 
  for i := 0; i < passwordLength; i++ {
    password += string(randomChar())
  }

  fmt.Println(password)

}

func randomChar() rune {
  chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  return rune(chars[rand.Intn(len(chars))])
}
