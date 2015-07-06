package main

import "fmt"

func main() {
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  // Scanf is used to receive user input at the command line. &input tells the function where to store the user input.

  output := input * 2
  fmt.Println(output)
}