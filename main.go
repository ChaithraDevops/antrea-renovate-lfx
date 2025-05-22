package main

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
)

func main() {
  token := jwt.New(jwt.SigningMethodHS256)
  fmt.Println("Generated Token:", token)
}
