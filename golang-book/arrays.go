package main

import "fmt"

func main() {
  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83
  // Can also be declared as x := [5]float64{ 98, 93, 77, 82, 83 }
  // This type of declaration can be split onto multiple lines after the comma if the array gets too long. If it is split on multiple lines the last element needs a trailing comma.

  slice := x[0:3]
  // Although this is not used in this application, a slice is a section of another array that can be manipulated .
  fmt.Println(slice)

  var total float64 = 0
  for _, value := range x {
    // This is a type of for loop that will run for the length of the array. "range x" means loop over the x array. value is a reference to the array element for each loop.
    // the _ at the start of the declaration tells Go that we do not need access to the iterator.
    total += value
  }

  fmt.Println(total / float64(len(x)))
  // In order to divide the total by the length of the array, we have to convert it to the same data type as the array data.
}