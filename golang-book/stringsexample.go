package main

import "fmt"

func main() {
  fmt.Println(len("Hello World"))
  // len is the way to return the length of a function.
  fmt.Println("Hello World"[1])
  // returns the character at index 1 as a byte (so 101).
  fmt.Println("Hello " + "World")
}