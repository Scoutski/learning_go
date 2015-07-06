package main

import "fmt"

var (
  a = 5
  b = 10
  c = 15
)
// This syntax allows us to define multiple new variables which would all have a global scope in this file.

func main() {
  var x string = "first"
  // x := "string" could also be written and go would interpret the type from the literal string passed in. The shorter form is definitely acceptable and should be used when possible.
  fmt.Println(x)
  x = "second"
  fmt.Println(x)
  // This will print the same variable twice but the variable will have different contents at each point.
}