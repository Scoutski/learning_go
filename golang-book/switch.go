package main

import "fmt"

func main() {
  i := 5

  switch i {
  case 1: fmt.Println("one")
  case 2: fmt.Println("two")
  case 3: fmt.Println("three")
  case 4: fmt.Println("four")
  case 5: fmt.Println("five")
  case 6: fmt.Println("six")
  default: fmt.Println("Unknown Number")
  // Will use default case if no others are met
  }
}