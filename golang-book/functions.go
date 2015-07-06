package main

import "fmt"

func main() {
  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))
}

func average(xs []float64) float64 {
  // xs is the paramater name and []float64 is the type (a slice of float64 elements).
  // The float64 after the paranthesis is the return type from the function.
  total := 0.0
  for _, number := range xs {
    total += number
  }
  return total / float64(len(xs))
}

// It is possible to return multiple values from a function if necessary. Here is an example:
// func f() (int, int) {
//   return 5, 6
// }

// func main() {
//   x, y := f()
// }

// To declare that a function takes 0 or more parametes, ... should be written before the type name in the declaration. This is called a Variadic function.
// e.g. func thisThing(x ...int) {   }
// This method is often used if an error can be returned, which makes it optional.
