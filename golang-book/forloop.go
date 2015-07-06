package main

import "fmt"

func main() {
  i := 1
  // iterator declared
  for i <= 10 {
    // If the statement after the for declaration is true, the code will run. It will continue to run until that statement is false.
    fmt.Println(i)
    i++
  }
  // This for loop can be written much more like javascript for familiarity, like this:
  // for i := 1; i <= 10; i++ {
  //  fmt.Println(i)
  // }
}